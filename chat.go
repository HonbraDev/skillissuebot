package skillissue

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/Tnze/go-mc/chat"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
)

var uptimeCheckUUID = uuid.MustParse("8be486a9-ea4d-4265-b74a-5c3fa24e6648")

func (b *Bot) SendChatMsg(msg string) error {
	if len(msg) > 256 {
		return errors.New("message too long")
	}
	return b.Client.Conn.WritePacket(pk.Marshal(0x03, pk.String(msg)))
}

func (b *Bot) Whisper(player, msg string) error {
	return b.SendChatMsg("/w " + player + " " + msg)
}

func (b *Bot) handleChat(c chat.Message, pos byte, u uuid.UUID) error {
	b.Logger.Println(pos, c.String())

	switch pos {
	case 0:
		return b.handleChatPlayer(c, u)
	case 1:
		return b.handleChatSystem(c)
	}

	return nil
}

func (b *Bot) handleChatPlayer(c chat.Message, u uuid.UUID) error {
	sender, msg := splitChatMsgText(c)
	if strings.HasPrefix(msg, b.CommandPrefix) {
		return b.handleCommand(sender, msg, false)
	}
	if u == uptimeCheckUUID && msg == "chat test" {
		return b.SendChatMsg("skill test")
	}
	return nil
}

func (b *Bot) handleChatSystem(c chat.Message) error {
	t := c.Translate
	switch {
	case strings.HasPrefix(t, "death."):
		return b.handleChatDeath(c)
	case t == "commands.message.display.incoming":
		return b.handleChatWhisper(splitChatMsgText(c))
	}
	return nil
}

func (b *Bot) handleChatWhisper(sender, msg string) error {
	return b.handleCommand(sender, msg, true)
}

func splitChatMsgText(c chat.Message) (string, string) {
	var m1, m2 chat.Message
	if err := json.Unmarshal(c.With[0], &m1); err != nil {
		return "", ""
	}
	if err := json.Unmarshal(c.With[1], &m2); err != nil {
		return "", ""
	}
	name := m1.Text
	for _, c := range m1.Extra {
		name = string(c.Text)
	}
	return name, m2.Text
}
