package skillissue

import "fmt"

type Position struct {
	X, Y, Z int
}

func (p Position) Add(x, y, z int) Position {
	return Position{p.X + x, p.Y + y, p.Z + z}
}

func (p Position) Sub(x, y, z int) Position {
	return Position{p.X - x, p.Y - y, p.Z - z}
}

func (p Position) AddPos(pos Position) Position {
	return Position{p.X + pos.X, p.Y + pos.Y, p.Z + pos.Z}
}

func (p Position) SubPos(pos Position) Position {
	return Position{p.X - pos.X, p.Y - pos.Y, p.Z - pos.Z}
}

func (p Position) Equals(x, y, z int) bool {
	return p.X == x && p.Y == y && p.Z == z
}

func (p Position) EqualsPos(pos Position) bool {
	return p.X == pos.X && p.Y == pos.Y && p.Z == pos.Z
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.X, p.Y, p.Z)
}

func (p Position) GetCoords() (int, int, int) {
	return p.X, p.Y, p.Z
}
