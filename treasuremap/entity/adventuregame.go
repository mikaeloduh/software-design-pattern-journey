package entity

import (
	"fmt"
	"math"
	"math/rand"
	"treasuremap/utils"
)

type Round int

type AttackRange [10][10]int

type IAttackStrategy func(worldMap [10][10]*Position) (attackRange AttackRange)

type AdventureGame struct {
	WorldMap  [10][10]*Position
	Character *Character
	round     Round
}

func NewAdventureGame(character *Character) *AdventureGame {
	game := &AdventureGame{
		Character: character,
	}

	nonRepeatIntStack := utils.RandNonRepeatIntStack(0, 99, 5)

	pop, _ := nonRepeatIntStack.Pop()
	x, y := randNonRepeatPosition(pop)
	game.AddObject(character, x, y, Up)

	pop, _ = nonRepeatIntStack.Pop()
	x, y = randNonRepeatPosition(pop)
	game.AddObject(NewTreasure(), x, y, Up)

	pop, _ = nonRepeatIntStack.Pop()
	x, y = randNonRepeatPosition(pop)
	game.AddObject(NewTreasure(), x, y, Up)

	pop, _ = nonRepeatIntStack.Pop()
	x, y = randNonRepeatPosition(pop)
	game.AddObject(NewTreasure(), x, y, Up)

	pop, _ = nonRepeatIntStack.Pop()
	x, y = randNonRepeatPosition(pop)
	game.AddObject(NewTreasure(), x, y, Up)

	return game
}

func randNonRepeatPosition(num int) (int, int) {
	return num % 10, int(math.Floor(float64(num / 10)))
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

	var doAttack = func(attackRange AttackRange) {
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

	attackRange := attackStrategy(g.WorldMap)

	doAttack(attackRange)
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

func randNewMapObject() IStatefulMapObject {
	return [3]func() IStatefulMapObject{
		func() IStatefulMapObject { return NewMonster() },
		func() IStatefulMapObject { return NewMonster() },
		func() IStatefulMapObject { return NewMonster() },
	}[rand.Intn(1)]()
}
