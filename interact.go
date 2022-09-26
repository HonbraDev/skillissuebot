package skillissue

import (
	pkid "github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

func (b *Bot) Interact(p Position, face int) error {
	return b.Client.Conn.WritePacket(pk.Marshal(
		pkid.ServerboundUseItemOn,
		pk.VarInt(0),      // main hand
		pk.Position(p),    // block position
		pk.VarInt(face),   // face (top)
		pk.Float(0.5),     // cursor position
		pk.Float(0.5),     //
		pk.Float(0.5),     //
		pk.Boolean(false), // inside block
	))
}
