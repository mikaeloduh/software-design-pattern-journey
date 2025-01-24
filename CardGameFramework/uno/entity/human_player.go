package entity

import "fmt"

// HumanUnoPlayer represents a human player in the UNO game.
type HumanUnoPlayer struct {
	Name string
	Hand []UnoCard
}

func (p *HumanUnoPlayer) Rename() {}

// SetCard adds a card to the player's hand.
func (p *HumanUnoPlayer) SetCard(card UnoCard) {
	p.Hand = append(p.Hand, card)
}

// TakeTurn allows the player to choose a card to play.
func (p *HumanUnoPlayer) TakeTurn() UnoCard {
	fmt.Printf("\n%s's turn. Your hand: %v\n", p.GetName(), p.GetHand())
	var cardIndex int
	for {
		fmt.Print("Enter the index of the card you want to play: ")
		_, err := fmt.Scan(&cardIndex)
		if err == nil && cardIndex >= 0 && cardIndex < len(p.Hand) {
			break
		}
		fmt.Println("Invalid input. Try again.")
	}
	card := p.Hand[cardIndex]
	p.Hand = append(p.Hand[:cardIndex], p.Hand[cardIndex+1:]...)

	return card
}

// GetName returns the player's name.
func (p *HumanUnoPlayer) GetName() string {
	return p.Name
}

// GetHand returns the player's hand.
func (p *HumanUnoPlayer) GetHand() []UnoCard {
	return p.Hand
}

func (p *HumanUnoPlayer) AddPoint() {
}
