package main

import (
	"fmt"
	"io"
)

type Character struct {
	Writer io.Writer
}

func NewCharacter(writer io.Writer) *Character {
	return &Character{Writer: writer}
}

func (c *Character) SetState(s IState) {
	fmt.Fprint(c.Writer, "The character is poisoned.\n")
}

type IState interface {
	OnRoundStart()
}

type PoisonedState struct {
	character *Character
}

func NewPoisonedState(character *Character) *PoisonedState {
	return &PoisonedState{character: character}
}

func (s *PoisonedState) OnRoundStart() {
	//s.character.AddHp(-15)
	fmt.Println("Poisoned: -15 HP")
}

func main() {
	fmt.Println("Hello world")
}
