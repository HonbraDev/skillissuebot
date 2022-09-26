package skillissue

import (
	"errors"
	"strconv"
	"strings"
)

var errorUnwhisperable = errors.New("this command is unwhisperable")

var commands = map[string]Command{
	// public commands
	"shrug": {func(c *CommandContext) (string, error) {
		if c.IsWhisper {
			return "", errorUnwhisperable
		}
		return "", c.Bot.SendChatMsg("¯\\_(ツ)_/¯")
	}, 0, false},
	"lenny": {func(c *CommandContext) (string, error) {
		if c.IsWhisper {
			return "", errorUnwhisperable
		}
		return "", c.Bot.SendChatMsg("( ͡° ͜ʖ ͡°)")
	}, 0, false},
	"tableflip": {func(c *CommandContext) (string, error) {
		if c.IsWhisper {
			return "", errorUnwhisperable
		}
		return "", c.Bot.SendChatMsg("(╯°□°）╯︵ ┻━┻")
	}, 0, false},
	"skillissue": {func(c *CommandContext) (string, error) {
		return "", c.Bot.SendChatMsg("skill issue ↑")
	}, 0, false},
	"sleep": {func(c *CommandContext) (string, error) {
		if err := c.Bot.SendChatMsg(",sleep"); err != nil {
			return "", err
		}
		return "Attempting to sleep.", c.Bot.Interact(c.Bot.BedCoords, 1)
	}, 0, false},

	// formalities
	"help": {func(c *CommandContext) (string, error) {
		return "No one is around to help. (shrug, lenny, tableflip, skillissue, sleep)", nil
	}, 0, false},
	"about": {func(c *CommandContext) (string, error) {
		return "Bot written in Go by Honbra. Contact Honbra#0082/@honbra:honbra.com if the bot misbehaves.", nil
	}, 0, false},
	"bot": {func(c *CommandContext) (string, error) {
		return "[iambot] Skill issue. github.com/HonbraDev/skillissuebot", nil
	}, 0, false},

	// restricted commands
	"sudo": {func(c *CommandContext) (string, error) {
		c.Bot.SendChatMsg(strings.Join(c.Args, " "))
		return "", nil
	}, -1, true},
	"interact": {func(c *CommandContext) (string, error) {
		var d [3]int
		for i, arg := range c.Args {
			var err error
			if d[i], err = strconv.Atoi(arg); err != nil {
				return "", err
			}
		}
		return "", c.Bot.Interact(Position{d[0], d[1], d[2]}, 1)
	}, 3, true},
}
