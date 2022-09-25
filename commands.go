package skillissue

import (
	"strings"
)

var commands = map[string]Command{
	// public commands
	"shrug": {func(c *CommandContext) (string, error) {
		if c.IsWhisper {
			return "¯\\_(ツ)_/¯", nil
		}
		return "", c.Bot.SendChatMsg("¯\\_(ツ)_/¯")
	}, 0, false},

	// formalities
	"help": {func(c *CommandContext) (string, error) {
		return "No one is around to help. (!shrug)", nil
	}, 0, false},
	"about": {func(c *CommandContext) (string, error) {
		return "Bot written in Go by Honbra. Contact Honbra#0082/@honbra:honbra.com if the bot misbehaves.", nil
	}, 0, false},
	"bot": {func(c *CommandContext) (string, error) {
		return "[iambot] Skill issue", nil
	}, 0, false},

	// restricted commands
	"sudo": {func(c *CommandContext) (string, error) {
		c.Bot.SendChatMsg(strings.Join(c.Args, " "))
		return "", nil
	}, -1, true},
}
