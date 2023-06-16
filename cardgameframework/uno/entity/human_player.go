package entity

import "fmt"

// HumanPlayer represents a human player in the UNO game.
type HumanPlayer struct {
	Name string
	Hand []Card
}

// SetCard adds a card to the player's hand.
func (p *HumanPlayer) SetCard(card Card) {
	p.Hand = append(p.Hand, card)
}

// TakeTurn allows the player to choose a card to play.
func (p *HumanPlayer) TakeTurn() Card {
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
func (p *HumanPlayer) GetName() string {
	return p.Name
}

// GetHand returns the player's hand.
func (p *HumanPlayer) GetHand() []Card {
	return p.Hand
}
