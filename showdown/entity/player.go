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
	Input
}

func (p *HumanPlayer) YouExchangeMyCard(card Card) (Card, error) {
	if len(p.HandCards) < 1 {
		err := errors.New(fmt.Sprintf("Player %d does not have enough cards to proceed with the exchange.", p.id))
		fmt.Printf("Error: %v", err)
		return Card{}, err
	}

	fmt.Printf("Player %s, please select your card to exchange back: ", p.name)
	toPlay := p.InputNum(0, len(p.HandCards)-1)
	myCard := p.HandCards[toPlay]
	p.HandCards[toPlay] = card

	return myCard, nil
}

func (p *HumanPlayer) MeExchangeYourCard(otherPlayer IPlayer) error {
	if len(p.HandCards) < 1 {
		err := errors.New(fmt.Sprintf("Player %d (You) don't have enough cards to proceed with the exchange.", p.id))
		fmt.Printf("Error: %v", err)
		return err
	}

	fmt.Printf("Please select your card to exchange: ")
	toPlay := p.InputNum(0, len(p.HandCards)-1)
	c := p.HandCards[toPlay]

	ex, err := otherPlayer.YouExchangeMyCard(c)
	if err != nil {
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
	fmt.Printf("\n* Now is player %s 's turn.\n", p.name)

	// 1. exchange
	if !p.usedExchange {
		fmt.Printf("Player %s, do you want to exchange hand card? ", p.name)
		if p.InputBool() {
			var toExchangeCard func()
			toExchangeCard = func() {
				fmt.Printf("Which player do you want to exchange cards with? ")
				p.whoExchangeWith = players[p.InputNum(0, 3)]
				if err := p.MeExchangeYourCard(p.whoExchangeWith); err != nil {
					toExchangeCard()
				}
			}
			toExchangeCard()
			p.usedExchange = true
		}
	} else {
		p.count--
		if p.count == 0 {
			fmt.Println("Exchange back")
			_ = p.MeExchangeYourCard(p.whoExchangeWith)
		}
	}

	// 2. show
	fmt.Printf("Player %s, please select a card to show ", p.name)
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

func (p *HumanPlayer) ReName() {
	fmt.Printf("Player %d enter you name: ", p.id)
	p.SetName(p.InputString())
}

func NewHumanPlayer(id int, input Input) *HumanPlayer {
	return &HumanPlayer{
		id:           id,
		name:         fmt.Sprintf("Player-%d", id),
		count:        3,
		usedExchange: false,
		Input:        input,
	}
}
