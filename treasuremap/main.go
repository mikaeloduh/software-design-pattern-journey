package main

import (
	"fmt"
	"io"
)

type AdventureGame struct {
	round     int
	character *Character
}

func NewAdventureGame(character *Character) *AdventureGame {
	return &AdventureGame{
		character: character,
	}
}

func (g *AdventureGame) StartRound() {
	g.round++
	g.character.OnRoundStart()
}

type Character struct {
	Writer io.Writer
	Hp     int
	State  IState
}

func NewCharacter(writer io.Writer) *Character {
	var c *Character
	c = &Character{Writer: writer, Hp: 300, State: &NormalState{c}}
	return c
}

func (c *Character) AddHp(num int) {
	c.Hp += num
}

func (c *Character) SetState(s IState) {
	c.State = s
	fmt.Fprint(c.Writer, "The character is poisoned.\n")
}

func (c *Character) OnRoundStart() {
	c.State.OnRoundStart()
}

// IState
type IState interface {
	OnRoundStart()
}

type NormalState struct {
	character *Character
}

func (s *NormalState) OnRoundStart() {
}

type PoisonedState struct {
	character *Character
}

func NewPoisonedState(character *Character) *PoisonedState {
	return &PoisonedState{character: character}
}

func (s *PoisonedState) OnRoundStart() {
	s.character.AddHp(-15)
}

func main() {
	fmt.Println("Hello world")
}
