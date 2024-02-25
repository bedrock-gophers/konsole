package konsole

import (
	"net/http"
	"strings"
	"time"
)

var rateLimited = make(map[string]time.Time)

func (ws *WebSocketServer) middleware(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := resolveIP(r.RemoteAddr)
		if t, ok := rateLimited[ip]; ok && t.After(time.Now()) {
			return
		}

		f.ServeHTTP(w, r)
	})
}

func resolveIP(addr string) string {
	ip := addr[:strings.LastIndex(addr, ":")]
	ip = strings.ReplaceAll(ip, "[", "")
	ip = strings.ReplaceAll(ip, "]", "")
	return ip
}
