package entity

import (
	"fmt"
	"math"
	"math/rand"
	"treasuremap/utils"
)

type Round int

type AdventureGame struct {
	WorldMap  [10][10]*Position
	Character *Character
	round     Round
}

func NewAdventureGame(character *Character) *AdventureGame {
	game := &AdventureGame{
		Character: character,
	}

	nonRepeatInt := utils.RandNonRepeatInt(0, 99, 5)
	num := nonRepeatInt[0]
	x, y := num%10, int(math.Floor(float64(num/10)))
	game.AddObject(character, x, y, Up)

	for _, num := range nonRepeatInt[1:] {
		x, y := num%10, int(math.Floor(float64(num/10)))
		game.AddObject(randNewMapObject(), x, y, Up)
	}

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
	g.Character.OnRoundStart()
}

type AttackMap [10][10]int

func (g *AdventureGame) Attack(attack AttackMap) {
	for y, v := range attack {
		for x, w := range v {
			if w != 0 && g.WorldMap[y][x] != nil {
				if g.WorldMap[y][x].Object.TakeDamage(w) <= 0 {
					g.WorldMap[y][x] = nil
				}
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

func randNewMapObject() IMapObject {
	return [3]func() IMapObject{
		func() IMapObject { return NewMonster() },
	}[rand.Intn(1)]()
}
