package skillissue

import (
	"errors"

	"github.com/Tnze/go-mc/chat"
)

func (b *Bot) handleDisconnect(reason chat.Message) error {
	return errors.New(reason.String())
}
