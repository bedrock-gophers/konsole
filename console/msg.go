package console

import "github.com/sandertv/gophertunnel/minecraft/text"

var (
	MessageLoginRequest  = text.Colourf("<red>Please enter the password.</red>")
	MessageWrongPassword = text.Colourf("<red>The password you have entered is incorrect.</red>")

	MessageLoginSuccess   = text.Colourf("<green>Successfuly logged into console.</green>")
	MessageUnknownCommand = "<red>Command with the name '%s' does not exist</red>"
)
