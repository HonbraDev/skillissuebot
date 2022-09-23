package skillissue

func (b *Bot) handleGameStart() error {
	return b.SendChatMsg("Skill issue bot v2 activated.")
}
