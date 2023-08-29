package entity

import (
	"fmt"
	"io"
	"math"
)

type Direction string

const (
	Up Direction = "up"
)

type Position struct {
	game      *AdventureGame
	character *Character
	x         int
	y         int
	direction Direction
}

func (p *Position) move(x, y int, d Direction) {
	px := p.x
	py := p.y

	p.x = x
	p.y = y
	p.direction = d

	p.game.WorldMap[px][py], p.game.WorldMap[p.x][p.y] = p.game.WorldMap[p.x][p.y], p.game.WorldMap[px][py]
}

type Character struct {
	Writer   io.Writer
	MaxHp    int
	Hp       int
	State    IState
	Speed    int // actions per round
	position *Position
}

func (c *Character) MoveStep(direction Direction) {
	switch direction {
	case Up:
		//c.position.move(c.position.x+1, c.position.y, direction)
		fmt.Fprint(c.Writer, "move up\n")
	}
}

func NewCharacter(writer io.Writer) *Character {
	var c *Character
	c = &Character{
		Writer: writer,
		MaxHp:  300,
		Hp:     300,
		State:  NewNormalState(c),
		Speed:  1,
	}
	return c
}

func (c *Character) SetState(s IState) {
	c.State = s
}

func (c *Character) OnRoundStart() {
	c.State.OnRoundStart()

	for !c.isRoundEnd() {
		c.AfterRoundStart()
	}

	c.OnRoundEnd()
}

func (c *Character) AfterRoundStart() {
	fmt.Fprint(c.Writer, "take action\n")

	c.Speed--
}

func (c *Character) OnRoundEnd() {
	c.Speed = 1
}

func (c *Character) isRoundEnd() bool {
	return c.Speed <= 0
}

func (c *Character) TakeDamage(damage int) {
	c.Hp -= c.State.OnTakeDamage(damage)
}

func (c *Character) Heal(health int) {
	c.Hp = int(math.Min(float64(float32(c.MaxHp)), float64(c.Hp+health)))
}

func (c *Character) SetSpeed(speed int) {
	c.Speed = speed
}
