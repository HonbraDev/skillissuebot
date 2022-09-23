package skillissue

import (
	"errors"
	"strings"
)

var commands = map[string]Command{
	// public commands
	"help": {func(c *CommandContext) (string, error) {
		return "h", nil
	}, 0, false},
	"ping": {func(c *CommandContext) (string, error) {
		return "pong", nil
	}, 0, false},
	"fail": {func(c *CommandContext) (string, error) {
		return "", errors.New("fail")
	}, 0, false},
	"about": {func(c *CommandContext) (string, error) {
		return "Bot written in Go by Honbra. Contact Honbra#0082/@honbra:honbra.com if the bot misbehaves.", nil
	}, 0, false},
	"echo": {func(c *CommandContext) (string, error) {
		msg := strings.Join(c.Args, " ")
		if len(msg) > 200 {
			return "", errors.New("message too long, silly")
		}
		return msg, nil
	}, -1, false},
	"bot": {func(c *CommandContext) (string, error) {
		return "[iambot] Skill issue", nil
	}, 0, false},

	// restricted commands
	"sudo": {func(c *CommandContext) (string, error) {
		c.Bot.SendChatMsg(strings.Join(c.Args, " "))
		return "", nil
	}, -1, true},
}
