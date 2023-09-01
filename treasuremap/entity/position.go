package entity

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

type Position struct {
	game      *AdventureGame
	object    IMapObject
	x         int
	y         int
	direction Direction
}

func (p *Position) move(x, y int, d Direction) {
	px := p.x
	py := p.y

	if err := p.game.MovePosition(px, py, x, y); err != nil {
		return
	}

	p.x = x
	p.y = y
	p.direction = d
}
