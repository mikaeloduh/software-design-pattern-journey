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

	game := AdventureGame{
		characters: character,
	}

	p := &Position{
		game:      &game,
		character: character,
		x:         5,
		y:         5,
		direction: Up,
	}

	game.WorldMap[5][5] = p

	//for _, num := range commons.RandNonRepeatInt(0, 99, 5) {
	//	x, y := num%10, int(math.Floor(float64(num/10)))
	//
	//	game.WorldMap[y][x] = p
	//}

	character.SetPosition(p)

	return &game
}

func (g *AdventureGame) StartRound() {
	g.round++
	g.characters.OnRoundStart()
}

func (g *AdventureGame) MovePosition(x1, y1, x2, y2 int) error {
	if g.WorldMap[y2][x2] != nil {
		return fmt.Errorf("invalid position")
	}
	g.WorldMap[y2][x2] = g.WorldMap[y1][x1]
	g.WorldMap[y1][x1] = nil

	return nil
}
