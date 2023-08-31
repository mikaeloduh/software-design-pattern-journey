package entity

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Character struct {
	Writer   io.Writer
	MaxHp    int
	Hp       int
	State    IState
	Speed    int // actions per round
	Position *Position
}

func NewCharacter() *Character {
	var c *Character
	c = &Character{
		Writer: os.Stdout,
		MaxHp:  300,
		Hp:     300,
		State:  NewNormalState(c),
		Speed:  1,
	}
	return c
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

func (c *Character) MoveStep(direction Direction) {
	switch direction {
	case Up:
		c.Position.move(c.Position.x, c.Position.y+1, Up)

	case Down:
		c.Position.move(c.Position.x, c.Position.y-1, Down)

	case Left:
		c.Position.move(c.Position.x-1, c.Position.y, Left)

	case Right:
		c.Position.move(c.Position.x+1, c.Position.y, Right)

	}
}

func (c *Character) SetState(s IState) {
	c.State = s
}

func (c *Character) SetSpeed(speed int) {
	c.Speed = speed
}

func (c *Character) SetPosition(p *Position) {
	c.Position = p
}
