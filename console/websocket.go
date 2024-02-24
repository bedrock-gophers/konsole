package console

import (
	"fmt"
	"github.com/bedrock-gophers/console/console/sets"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebSocketServer struct {
	sub chat.Subscriber

	connectionMu sync.Mutex
	connections  sets.Set[*websocket.Conn]

	password  string
	formatter Formatter
	router    *mux.Router
}

func NewWebSocketServer(sub chat.Subscriber, password string, f Formatter) *WebSocketServer {
	if f == nil {
		f = NopFormatter{}
	}

	ws := &WebSocketServer{
		sub:         sub,
		connections: sets.New[*websocket.Conn](),
		password:    password,
		formatter:   f,
		router:      mux.NewRouter(),
	}
	return ws
}

func (ws *WebSocketServer) Message(a ...any) {
	ws.sub.Message(a...)

	s := make([]string, len(a))
	for i, b := range a {
		s[i] = fmt.Sprint(b)
	}

	for _, c := range ws.connections.Values() {
		_ = c.WriteMessage(websocket.TextMessage, []byte(strings.Join(s, " ")))
	}
}

func (ws *WebSocketServer) route() {
	ws.router.HandleFunc("/", ws.processRequest)
	ws.router.Use(ws.middleware)
}

func (ws *WebSocketServer) ListenAndServe(addr string) error {
	ws.route()
	return http.ListenAndServe(addr, ws.router)
}

func (ws *WebSocketServer) ListenAndServeTLS(addr string, certFile, keyFile string) error {
	ws.route()
	return http.ListenAndServeTLS(addr, certFile, keyFile, ws.router)
}

func (ws *WebSocketServer) processRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	go ws.handleConn(conn)
}

func (ws *WebSocketServer) handleConn(conn *websocket.Conn) {
	_ = conn.SetReadDeadline(time.Now().Add(time.Minute))

	defer func() {
		ws.connections.Delete(conn)
	}()

	if !ws.connections.Contains(conn) {
		if conn.WriteMessage(websocket.TextMessage, []byte(MessageLoginRequest)) != nil {
			return
		}
	}
	for {
		n, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if n == websocket.TextMessage {
			ws.handleMessage(conn, strings.TrimSpace(string(msg)))
		}
	}
}

func (ws *WebSocketServer) handleMessage(conn *websocket.Conn, msg string) {
	if !ws.connections.Contains(conn) {
		if msg != ws.password {
			ip := resolveIP(conn.RemoteAddr().String())
			rateLimited[ip] = time.Now().Add(time.Second * 3)

			if conn.WriteMessage(websocket.TextMessage, []byte(MessageWrongPassword)) != nil {
				return
			}
			_ = conn.Close()
			return
		}
		if conn.WriteMessage(websocket.TextMessage, []byte(MessageLoginSuccess)) != nil {
			return
		}
		ws.connections.Add(conn)
		_ = conn.SetReadDeadline(time.Time{})
		return
	}
	m := ws.formatter.FormatMessage(msg)

	if strings.HasPrefix(msg, "!") {
		m = ws.formatter.FormatAlert(strings.TrimPrefix(msg, "!"))
	} else if strings.HasPrefix(msg, "/") {
		msg = strings.TrimPrefix(msg, "/")
		c, ok := cmd.ByAlias(strings.Split(msg, " ")[0])
		if !ok {
			_ = conn.WriteMessage(websocket.TextMessage, []byte(text.Colourf(MessageUnknownCommand, msg)))
			return
		}
		c.Execute(strings.TrimPrefix(strings.TrimPrefix(msg, strings.Split(msg, " ")[0]), " "), source{conn: conn})
		return
	}
	_, _ = chat.Global.WriteString(m)
}
