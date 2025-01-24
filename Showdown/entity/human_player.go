package entity

import (
	"errors"
	"fmt"
)

type HumanPlayer struct {
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

func (p *HumanPlayer) SetId(i int) {
	p.id = i
}

func (p *HumanPlayer) YouExchangeMyCard(card Card) (Card, error) {
	if len(p.HandCards) < 1 {
		err := errors.New(fmt.Sprintf("%s does not have enough cards to proceed with the exchange.", p.name))
		return Card{}, err
	}

	p.PrintCardsOutput(p.HandCards)
	p.YouExchangeMyCardOutput(p.name)
	toPlay := p.InputNum(0, len(p.HandCards)-1)
	myCard := p.HandCards[toPlay]
	p.HandCards[toPlay] = card

	return myCard, nil
}

func (p *HumanPlayer) MeExchangeYourCard(otherPlayer IPlayer) error {
	if len(p.HandCards) < 1 {
		err := errors.New(fmt.Sprintf("%s (You) don't have enough cards to proceed with the exchange.", p.name))
		fmt.Printf("Error: %v", err)
		return err
	}

	p.MeExchangeYourCardOutput()
	toPlay := p.InputNum(0, len(p.HandCards)-1)
	c := p.HandCards[toPlay]

	ex, err := otherPlayer.YouExchangeMyCard(c)
	if err != nil {
		p.MeExchangeYourCardErrorOutput(err)
		return err
	}
	p.HandCards[toPlay] = ex

	return nil
}

func (p *HumanPlayer) Point() int {
	return p.point
}

func (p *HumanPlayer) AddPoint() {
	p.point += 1
}

func (p *HumanPlayer) TakeTurn(players []IPlayer) Card {
	p.TakeTurnStartOutput(p.name)
	p.PrintCardsOutput(p.HandCards)

	var toExchangeCard func()
	toExchangeCard = func() {
		p.ToExchangeCardOutput()
		p.whoExchangeWith = players[p.InputNum(0, len(players)-1)]
		if err := p.MeExchangeYourCard(p.whoExchangeWith); err != nil {
			toExchangeCard()
		}
	}

	// 1. exchange
	if !p.usedExchange {
		p.AskToExchangeCardOutput(p.name)
		if p.InputBool() {
			toExchangeCard()
			p.usedExchange = true
		}
	} else {
		p.count--
		if p.count == 0 {
			p.ExchangeBackOutput()
			_ = p.MeExchangeYourCard(p.whoExchangeWith)
		}
	}

	// 2. show
	p.AskShowCardOutput(p.name)
	toPlay := p.InputNum(0, len(p.HandCards)-1)
	showCard := p.HandCards[toPlay]
	p.HandCards = append([]Card{}, append(p.HandCards[0:toPlay], p.HandCards[toPlay+1:]...)...)

	return showCard
}

func (p *HumanPlayer) GetCard(card Card) {
	p.HandCards = append(p.HandCards, card)
}

func (p *HumanPlayer) Id() int {
	return p.id
}

func (p *HumanPlayer) Name() string {
	return p.name
}

func (p *HumanPlayer) SetName(name string) {
	p.name = name
}

func (p *HumanPlayer) Rename() {
	p.RenameOutput(p.name)

	s := p.InputString()
	if s != "?" {
		p.SetName(s)
	}
}

func NewHumanPlayer(input IInput, output IOutput) *HumanPlayer {
	return &HumanPlayer{
		count:        3,
		usedExchange: false,
		IInput:       input,
		IOutput:      output,
	}
}
