package console

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/gorilla/websocket"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"strings"
)

type source struct {
	conn *websocket.Conn
}

func (s source) Position() mgl64.Vec3 {
	return mgl64.Vec3{}
}

func (s source) SendCommandOutput(o *cmd.Output) {
	var messages []string
	for _, m := range o.Messages() {
		messages = append(messages, m)
	}
	for _, e := range o.Errors() {
		messages = append(messages, text.Colourf("<red>%s</red>", e.Error()))
	}
	_ = s.conn.WriteMessage(websocket.TextMessage, []byte(strings.Join(messages, "\n")))
}

func (s source) World() *world.World {
	return nil
}
