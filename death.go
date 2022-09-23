package skillissue

func (b *Bot) handleDeath() error {
	return b.Player.Respawn()
}
