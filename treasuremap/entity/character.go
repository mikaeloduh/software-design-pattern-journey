package entity

import (
	"fmt"
	"io"
	"math"
)

type Character struct {
	Writer io.Writer
	MaxHp  int
	Hp     int
	State  IState
}

func NewCharacter(writer io.Writer) *Character {
	var c *Character
	c = &Character{Writer: writer, MaxHp: 300, Hp: 300, State: NewNormalState(c)}
	return c
}

func (c *Character) SetState(s IState) {
	c.State = s
	fmt.Fprint(c.Writer, "The character is poisoned.\n")
}

func (c *Character) OnRoundStart() {
	c.State.OnRoundStart()
}

func (c *Character) TakeDamage(damage int) {
	c.Hp -= c.State.OnTakeDamage(damage)
}

func (c *Character) Heal(health int) {
	c.Hp = int(math.Min(float64(float32(c.MaxHp)), float64(c.Hp+health)))
}
