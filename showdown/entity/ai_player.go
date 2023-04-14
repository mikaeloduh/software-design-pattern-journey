package entity

import (
	"errors"
	"fmt"
)

type AIPlayer struct {
	id        int
	name      string
	HandCards []Card
	point     int
}

func (ai *AIPlayer) YouExchangeMyCard(card Card) (Card, error) {
	if len(ai.HandCards) < 1 {
		fmt.Printf("AI %d us fucking up", ai.id)
		return Card{}, errors.New(fmt.Sprintf("AI %d does not have enough cards to proceed with the exchange.", ai.id))
	}

	// TODO: Choose a card input
	myCard := ai.HandCards[0]
	ai.HandCards[0] = card

	return myCard, nil
}

func (ai *AIPlayer) MeExchangeYourCard(player IPlayer) error {
	if len(ai.HandCards) < 1 {
		fmt.Println("yr fucking up")
		return errors.New(fmt.Sprintf("Player %d (AI) does not have enough cards to proceed with the exchange.", ai.id))
	}

	// TODO: Choose a card input
	c := ai.HandCards[0]

	ex, err := player.YouExchangeMyCard(c)
	if err != nil {
		return err
	}
	ai.HandCards[0] = ex

	return nil
}

func (ai *AIPlayer) Point() int {
	return ai.point
}

func (ai *AIPlayer) AddPoint() {
	ai.point += 1
}

func (ai *AIPlayer) TakeTurn(players []IPlayer) Card {
	// TODO: 1. exchange?

	// 2. Show card
	play := 0
	showCard := ai.HandCards[play]
	ai.HandCards = append([]Card{}, append(ai.HandCards[0:play], ai.HandCards[play+1:]...)...)

	return showCard
}

func (ai *AIPlayer) Id() int {
	return ai.id
}

func (ai *AIPlayer) Name() string {
	return ai.name
}

func (ai *AIPlayer) GetCard(card Card) {
	ai.HandCards = append(ai.HandCards, card)
}

func NewAIPlayer(id int) *AIPlayer {
	return &AIPlayer{
		id:   id,
		name: "AI has no name",
	}
}
