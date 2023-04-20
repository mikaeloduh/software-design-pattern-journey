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

func (p *HumanPlayer) SetId(i int) {
	p.id = i
}

func (p *HumanPlayer) YouExchangeMyCard(card Card) (Card, error) {
	if len(p.HandCards) < 1 {
		err := errors.New(fmt.Sprintf("%s does not have enough cards to proceed with the exchange.", p.name))
		fmt.Printf("Error: %v", err)
		return Card{}, err
	}

	printCards(p.HandCards)
	fmt.Printf("%s, please select your card to exchange back: ", p.name)
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
	fmt.Printf("\n* Now is %s 's turn.\n", p.name)

	printCards(p.HandCards)

	// 1. exchange
	if !p.usedExchange {
		fmt.Printf("%s, do you want to exchange hand card? ", p.name)
		if p.InputBool() {
			var toExchangeCard func()
			toExchangeCard = func() {
				fmt.Printf("Which player do you want to exchange cards with? ")
				p.whoExchangeWith = players[p.InputNum(0, len(players)-1)]
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
	fmt.Printf("%s, please select a card to show ", p.name)
	toPlay := p.InputNum(0, len(p.HandCards)-1)
	showCard := p.HandCards[toPlay]
	p.HandCards = append([]Card{}, append(p.HandCards[0:toPlay], p.HandCards[toPlay+1:]...)...)

	return showCard
}

func printCards(cards []Card) {
	for i, c := range cards {
		if i%5 == 0 && i != 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%2d : [%4s ]  ", i, c.String())
	}
	fmt.Print("\n")
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
	fmt.Printf("%s, please enter your name: ", p.Name())

	s := p.InputString()
	if s != "?" {
		p.SetName(s)
	}
}

func NewHumanPlayer(input Input) *HumanPlayer {
	return &HumanPlayer{
		count:        3,
		usedExchange: false,
		Input:        input,
	}
}
