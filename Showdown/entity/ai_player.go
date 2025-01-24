package entity

import (
	"errors"
	"fmt"
)

type AIPlayer struct {
	id              int
	name            string
	HandCards       []Card
	point           int
	usedExchange    bool
	count           int
	whoExchangeWith IPlayer
	IInput
	IOutput
}

func (ai *AIPlayer) YouExchangeMyCard(card Card) (Card, error) {
	if len(ai.HandCards) < 1 {
		err := errors.New(fmt.Sprintf("%s (AI) does not have enough cards to proceed with the exchange.", ai.name))
		fmt.Printf("Error: %v", err)
		return Card{}, err
	}

	ai.YouExchangeMyCardOutput(ai.name)
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

	ai.MeExchangeYourCardOutput()
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
	ai.TakeTurnStartOutput(ai.name)

	var toExchangeCard func()
	toExchangeCard = func() {
		ai.whoExchangeWith = players[ai.InputNum(0, len(players)-1)]
		ai.ToExchangeCardOutput()
		if err := ai.MeExchangeYourCard(ai.whoExchangeWith); err != nil {
			toExchangeCard()
		}
	}

	// 1. exchange?
	if !ai.usedExchange {
		if ai.InputBool() {
			ai.AskToExchangeCardOutput(ai.name)
			toExchangeCard()
			ai.usedExchange = true
		}
	} else {
		ai.count--
		if ai.count == 0 {
			ai.ExchangeBackOutput()
			_ = ai.MeExchangeYourCard(ai.whoExchangeWith)
		}
	}

	// 2. Show card
	ai.AskShowCardOutput(ai.name)
	toPlay := ai.InputNum(0, len(ai.HandCards)-1)
	showCard := ai.HandCards[toPlay]
	ai.HandCards = append([]Card{}, append(ai.HandCards[0:toPlay], ai.HandCards[toPlay+1:]...)...)

	return showCard
}

func (ai *AIPlayer) Id() int {
	return ai.id
}

func (ai *AIPlayer) SetId(i int) {
	ai.id = i
}

func (ai *AIPlayer) Name() string {
	return ai.name
}

func (ai *AIPlayer) SetName(s string) {
	ai.name = s + "_AI"
}

func (ai *AIPlayer) Rename() {
}

func (ai *AIPlayer) GetCard(card Card) {
	ai.HandCards = append(ai.HandCards, card)
}

func NewAIPlayer(input IInput, output IOutput) *AIPlayer {
	return &AIPlayer{
		count:   3,
		name:    "PlayerAI",
		IInput:  input,
		IOutput: output,
	}
}
