package entity

import (
	"fmt"
	"math"
	"treasuremap/utils"
)

type Round int

type Damage int

type AttackDamageArea [10][10]Damage

type IAttackStrategy func(worldMap [10][10]*Position) AttackDamageArea

// AdventureGame
type AdventureGame struct {
	WorldMap  [10][10]*Position
	Character *Character
	round     Round
}

func NewAdventureGame() *AdventureGame {
	character := NewCharacter()
	game := &AdventureGame{Character: character}

	nonRepeatIntStack := utils.RandNonRepeatIntStack(0, 99, 5)

	num, _ := nonRepeatIntStack.Pop()
	game.AddObject(character, num%10, int(math.Floor(float64(num/10))), Up)

	num, _ = nonRepeatIntStack.Pop()
	game.AddObject(NewMonster(), num%10, int(math.Floor(float64(num/10))), Left)

	num, _ = nonRepeatIntStack.Pop()
	game.AddObject(NewTreasure(), num%10, int(math.Floor(float64(num/10))), None)

	num, _ = nonRepeatIntStack.Pop()
	game.AddObject(NewTreasure(), num%10, int(math.Floor(float64(num/10))), None)

	num, _ = nonRepeatIntStack.Pop()
	game.AddObject(NewTreasure(), num%10, int(math.Floor(float64(num/10))), None)

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

func (g *AdventureGame) Attack(attackStrategy IAttackStrategy) {

	var doAttack = func(attackRange AttackDamageArea) {
		for y, v := range attackRange {
			for x, damage := range v {
				if damage != 0 && g.WorldMap[y][x] != nil {
					if obj, ok := g.WorldMap[y][x].Object.(IStatefulMapObject); ok == true {
						if obj.TakeDamage(damage) <= 0 {
							g.WorldMap[y][x] = nil
						}
					}
				}
			}
		}
	}

	attackArea := attackStrategy(g.WorldMap)

	doAttack(attackArea)
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

func (g *AdventureGame) Run() {
	// TODO: To be implement
}
