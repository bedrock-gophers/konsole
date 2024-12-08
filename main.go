package main

import (
	"fmt"
	"github.com/bedrock-gophers/konsole/konsole"
	"github.com/bedrock-gophers/konsole/konsole/app"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"log/slog"
)

/*
 */
func main() {
	a := app.New("/n87asdh867adb68g63gd326g7d23g67dg23asiuhdhuiiusdhhuasdhuiauisauidgiugh2861286")
	go a.ListenAndServe(":6969")

	ws := konsole.NewWebSocketServer(chat.StdoutSubscriber{}, "test", testFormatter{})
	go ws.ListenAndServe(":8080")

	chat.Global.Subscribe(ws)
	conf, _ := server.DefaultConfig().Config(slog.Default())

	srv := conf.New()
	srv.CloseOnProgramEnd()
	cmd.Register(cmd.New("test", "", nil, testCommand{}, testCommand2{}))

	srv.Listen()
	for p := range srv.Accept() {
		p.Handle(testHandler{})
	}
}

type testHandler struct {
	player.NopHandler
}

func (h testHandler) HandleChat(ctx *player.Context, message *string) {
	ctx.Cancel()
	_, _ = chat.Global.WriteString(text.Colourf("<grey>%s: %s</grey>", ctx.Val().Name(), *message))
}

type testFormatter struct {
	konsole.NopFormatter
}

func (testFormatter) FormatMessage(s string) string {
	return text.Colourf("<purple>[CONSOLE]: %s</purple>", s)
}

func (testFormatter) FormatAlert(s string) string {
	return text.Colourf("<b><red>[</red><yellow>ALERT<red>]:</red> <yellow>%s</yellow></b>", s)
}

type testCommand struct {
	Sub cmd.SubCommand `cmd:"test"`
}

func (testCommand) Run(_ cmd.Source, _ *cmd.Output, _ *world.Tx) {
	fmt.Println("hey")
}

type testCommand2 struct {
	Sub cmd.SubCommand `cmd:"test2"`
}

func (testCommand2) Run(_ cmd.Source, _ *cmd.Output, _ *world.Tx) {
	fmt.Println("hey2")
}
