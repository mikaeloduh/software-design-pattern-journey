package entity

import (
	"fmt"
	"math/rand"
)

// ComputerPlayer represents a computer player in the UNO game.
type ComputerPlayer struct {
	Name string
	Hand []Card
}

// SetCard adds a card to the player's hand.
func (p *ComputerPlayer) SetCard(card Card) {
	p.Hand = append(p.Hand, card)
}

// TakeTurn randomly selects a card to play.
func (p *ComputerPlayer) TakeTurn() Card {
	fmt.Print("\nComputerPlayer's turn.\n")
	cardIndex := rand.Intn(len(p.Hand))
	card := p.Hand[cardIndex]
	p.Hand = append(p.Hand[:cardIndex], p.Hand[cardIndex+1:]...)
	return card
}

// GetName returns the player's name.
func (p *ComputerPlayer) GetName() string {
	return p.Name
}

// GetHand returns the player's hand.
func (p *ComputerPlayer) GetHand() []Card {
	return p.Hand
}
