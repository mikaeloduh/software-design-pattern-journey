package entity

import (
	"errors"
	"fmt"
)

type AIPlayer struct {
	id           int
	name         string
	HandCards    []Card
	point        int
	usedExchange bool
	count        int
	who          IPlayer
	Input
}

func (ai *AIPlayer) YouExchangeMyCard(card Card) (Card, error) {
	if len(ai.HandCards) < 1 {
		err := errors.New(fmt.Sprintf("Player %d (AI) does not have enough cards to proceed with the exchange.", ai.id))
		fmt.Printf("Error: %v", err)
		return Card{}, err
	}

	fmt.Printf("AI is selection card...")
	toPlay := ai.InputNum(0, len(ai.HandCards)-1)
	myCard := ai.HandCards[toPlay]
	ai.HandCards[toPlay] = card

	return myCard, nil
}

func (ai *AIPlayer) MeExchangeYourCard(player IPlayer) error {
	if len(ai.HandCards) < 1 {
		err := errors.New(fmt.Sprintf("Player %d (AI) does not have enough cards to proceed with the exchange.", ai.id))
		fmt.Printf("Error: %v", err)
		return err
	}

	fmt.Printf("AI is selection card...")
	toPlay := ai.InputNum(0, len(ai.HandCards)-1)
	c := ai.HandCards[toPlay]

	ex, err := player.YouExchangeMyCard(c)
	if err != nil {
		return err
	}
	ai.HandCards[toPlay] = ex

	return nil
}

func (ai *AIPlayer) Point() int {
	return ai.point
}

func (ai *AIPlayer) AddPoint() {
	ai.point += 1
}

func (ai *AIPlayer) TakeTurn(players []IPlayer) Card {
	fmt.Printf("Player %d (AI) is taking turn...\n", ai.id)

	// 1. exchange?
	if !ai.usedExchange {
		if ai.InputBool() {
			fmt.Printf("Player %d (AI) want to exchange card ", ai.id)
			var toExchangeCard func()
			toExchangeCard = func() {
				ai.who = players[ai.InputNum(0, 3)]
				if err := ai.MeExchangeYourCard(ai.who); err != nil {
					toExchangeCard()
				}
			}
			toExchangeCard()
			ai.usedExchange = true
		}
	} else {
		ai.count--
		if ai.count == 0 {
			fmt.Println("Exchange back")
			_ = ai.MeExchangeYourCard(ai.who)
		}
	}

	// 2. Show card
	fmt.Printf("AI is selecting card to show...")
	toPlay := ai.InputNum(0, len(ai.HandCards)-1)
	showCard := ai.HandCards[toPlay]
	ai.HandCards = append([]Card{}, append(ai.HandCards[0:toPlay], ai.HandCards[toPlay+1:]...)...)

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

func NewAIPlayer(id int, input Input) *AIPlayer {
	return &AIPlayer{
		id:    id,
		name:  "AI has no name",
		Input: input,
	}
}
