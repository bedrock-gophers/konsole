package main

import (
	"fmt"
	"github.com/bedrock-gophers/konsole/konsole"
	"github.com/bedrock-gophers/konsole/konsole/app"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
)

/*
 */
func main() {
	a := app.New("/n87asdh867adb68g63gd326g7d23g67dg23asiuhdhuiiusdhhuasdhuiauisauidgiugh2861286")
	go a.ListenAndServe(":6969")

	ws := konsole.NewWebSocketServer(chat.StdoutSubscriber{}, "test", testFormatter{})
	go ws.ListenAndServe(":8080")

	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(ws)
	conf, _ := server.DefaultConfig().Config(log)

	srv := conf.New()
	srv.CloseOnProgramEnd()
	cmd.Register(cmd.New("test", "", nil, testCommand{}, testCommand2{}))

	srv.Listen()
	for srv.Accept(func(p *player.Player) {
		p.Handle(testHandler{p: p})
	}) {
		// Do nothing
	}
}

type testHandler struct {
	player.NopHandler
	p *player.Player
}

func (h testHandler) HandleChat(ctx *event.Context, message *string) {
	ctx.Cancel()
	_, _ = chat.Global.WriteString(text.Colourf("<grey>%s: %s</grey>", h.p.Name(), *message))
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

func (testCommand) Run(_ cmd.Source, _ *cmd.Output) {
	fmt.Println("hey")
}

type testCommand2 struct {
	Sub cmd.SubCommand `cmd:"test2"`
}

func (testCommand2) Run(_ cmd.Source, _ *cmd.Output) {
	fmt.Println("hey2")
}
