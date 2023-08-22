package entity

import (
	"fmt"
	"io"
)

type Character struct {
	Writer io.Writer
	Hp     int
	State  IState
}

func NewCharacter(writer io.Writer) *Character {
	var c *Character
	c = &Character{Writer: writer, Hp: 300, State: NewNormalState(c)}
	return c
}

func (c *Character) TakeDamage(d int) {
	c.Hp -= c.State.OnTakeDamage(d)
}

func (c *Character) SetState(s IState) {
	c.State = s
	fmt.Fprint(c.Writer, "The character is poisoned.\n")
}

func (c *Character) OnRoundStart() {
	c.State.OnRoundStart()
}
