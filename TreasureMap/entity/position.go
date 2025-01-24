package entity

type Direction string

const (
	None  Direction = ""
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

type Position struct {
	Game      *AdventureGame
	Object    IMapObject
	X         int
	Y         int
	Direction Direction
}

func (p *Position) Move(x, y int, d Direction) {
	px := p.X
	py := p.Y

	if err := p.Game.MovePosition(px, py, x, y); err != nil {
		return
	}

	p.X = x
	p.Y = y
	p.Direction = d
}
