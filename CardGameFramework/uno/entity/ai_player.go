package entity

import (
	"fmt"
	"math/rand"
)

// AiUnoPlayer represents a computer player in the UNO game.
type AiUnoPlayer struct {
	Name string
	Hand []UnoCard
}

func (p *AiUnoPlayer) Rename() {}

// SetCard adds a card to the player's hand.
func (p *AiUnoPlayer) SetCard(card UnoCard) {
	p.Hand = append(p.Hand, card)
}

// TakeTurn randomly selects a card to play.
func (p *AiUnoPlayer) TakeTurn() UnoCard {
	fmt.Print("\nComputerPlayer's turn.\n")
	cardIndex := rand.Intn(len(p.Hand))
	card := p.Hand[cardIndex]
	p.Hand = append(p.Hand[:cardIndex], p.Hand[cardIndex+1:]...)
	return card
}

// GetName returns the player's name.
func (p *AiUnoPlayer) GetName() string {
	return p.Name
}

// GetHand returns the player's hand.
func (p *AiUnoPlayer) GetHand() []UnoCard {
	return p.Hand
}

func (p *AiUnoPlayer) AddPoint() {
}
