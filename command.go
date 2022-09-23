package skillissue

import (
	"strings"
)

type Command struct {
	Handler    func(c *CommandContext) (string, error)
	ArgNum     int
	Restricted bool
}

type CommandContext struct {
	Bot       *Bot
	Args      []string
	Invoker   string
	IsWhisper bool
}

func (b *Bot) handleCommand(sender, msg string, isWhisper bool) error {
	if sender == b.Client.Name {
		return nil
	}
	str := msg
	if isWhisper {
		if strings.HasPrefix(str, b.CommandPrefix) {
			return b.Whisper(sender, "silly goose uses command prefix in whispers")
		}
	} else {
		str = strings.TrimPrefix(msg, b.CommandPrefix)
	}
	tokens := strings.Split(str, " ")
	cmdName := tokens[0]
	args := tokens[1:]
	if cmd, ok := commands[cmdName]; ok {
		if cmd.ArgNum != -1 && len(args) != cmd.ArgNum {
			if isWhisper {
				return b.Whisper(sender, "Error: invalid number of arguments")
			} else {
				return nil
			}
		}
		if cmd.Restricted && !b.CoolPeople.IsIgnored(sender) {
			return b.Whisper(sender, "Error: silly goose")
		}
		ctx := CommandContext{
			Bot:     b,
			Args:    args,
			Invoker: sender,
		}
		msg, err := cmd.Handler(&ctx)
		if err != nil {
			return b.Whisper(sender, "Error: "+err.Error())
		}
		return b.Whisper(sender, msg)
	}
	return nil
}
