package entity

import (
	"fmt"
)

type Round int

type AdventureGame struct {
	WorldMap   [10][10]*Position
	characters *Character
	round      Round
}

func NewAdventureGame(character *Character) *AdventureGame {
	game := &AdventureGame{
		characters: character,
	}

	//for _, num := range commons.RandNonRepeatInt(0, 99, 5) {
	//	x, y := num%10, int(math.Floor(float64(num/10)))
	//
	//	game.WorldMap[y][x] = p
	//}

	game.AddObject(character, 5, 5, Up)

	return game
}

func (g *AdventureGame) AddObject(object IMapObject, x, y int, d Direction) {
	p := &Position{
		Game:      g,
		Object:    object,
		X:         x,
		Y:         y,
		Direction: d,
	}
	g.WorldMap[y][x] = p
	object.SetPosition(p)
}

func (g *AdventureGame) StartRound() {
	g.round++
	g.characters.OnRoundStart()
}

type XY struct {
	X int
	Y int
}

func (g *AdventureGame) AttackMap(damage int, area []XY) {
	for _, a := range area {
		if g.WorldMap[a.Y][a.X] != nil {
			if g.WorldMap[a.Y][a.X].Object.TakeDamage(damage) <= 0 {
				g.WorldMap[a.Y][a.X] = nil
			}
		}
	}
}

func (g *AdventureGame) MovePosition(x1, y1, x2, y2 int) error {
	if x2 < 0 || y2 < 0 {
		return fmt.Errorf("invalid input")
	}
	if g.WorldMap[y2][x2] != nil {
		return fmt.Errorf("invalid position")
	}
	g.WorldMap[y2][x2] = g.WorldMap[y1][x1]
	g.WorldMap[y1][x1] = nil

	return nil
}
