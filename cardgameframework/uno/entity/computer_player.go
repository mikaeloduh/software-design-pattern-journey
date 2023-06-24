package entity

import (
	"fmt"
	"math/rand"
)

// ComputerUnoPlayer represents a computer player in the UNO game.
type ComputerUnoPlayer struct {
	Name string
	Hand []UnoCard
}

func (p *ComputerUnoPlayer) Rename() {}

// SetCard adds a card to the player's hand.
func (p *ComputerUnoPlayer) SetCard(card UnoCard) {
	p.Hand = append(p.Hand, card)
}

// TakeTurn randomly selects a card to play.
func (p *ComputerUnoPlayer) TakeTurn() UnoCard {
	fmt.Print("\nComputerPlayer's turn.\n")
	cardIndex := rand.Intn(len(p.Hand))
	card := p.Hand[cardIndex]
	p.Hand = append(p.Hand[:cardIndex], p.Hand[cardIndex+1:]...)
	return card
}

// GetName returns the player's name.
func (p *ComputerUnoPlayer) GetName() string {
	return p.Name
}

// GetHand returns the player's hand.
func (p *ComputerUnoPlayer) GetHand() []UnoCard {
	return p.Hand
}
