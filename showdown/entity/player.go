package entity

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type HumanPlayer struct {
	id           int
	name         string
	HandCards    []Card
	point        int
	usedExchange bool
	count        int
	who          IPlayer
}

func (p *HumanPlayer) YouExchangeMyCard(card Card) (Card, error) {
	if len(p.HandCards) < 1 {
		fmt.Printf("Player %d us fucking up", p.id)
		return Card{}, errors.New(fmt.Sprintf("Player %d does not have enough cards to proceed with the exchange.", p.id))
	}

	// TODO: Choose a card input
	myCard := p.HandCards[0]
	p.HandCards[0] = card

	return myCard, nil
}

func (p *HumanPlayer) MeExchangeYourCard(player IPlayer) error {
	if len(p.HandCards) < 1 {
		fmt.Println("yr fucking up")
		return errors.New(fmt.Sprintf("Player %d (You) does not have enough cards to proceed with the exchange.", p.id))
	}

	// TODO: Choose a card input
	c := p.HandCards[0]

	ex, err := player.YouExchangeMyCard(c)
	if err != nil {
		return err
	}
	p.HandCards[0] = ex

	return nil
}

func (p *HumanPlayer) Point() int {
	return p.point
}

func (p *HumanPlayer) AddPoint() {
	p.point += 1
}

func (p *HumanPlayer) TakeTurn(players []IPlayer) Card {
	// 1. exchange?
	fmt.Printf("Player %d taking turn...\n", p.id)

	if !p.usedExchange {
		wantExchange := randomBool() // TODO: input
		if wantExchange {
			var toExchangeCard func()
			toExchangeCard = func() {
				fmt.Println("before exchange")
				p.who = players[(p.id+1)%4] // TODO: input
				if err := p.MeExchangeYourCard(p.who); err != nil {
					fmt.Println("try another exchange")
					toExchangeCard()
				}
			}
			toExchangeCard()
			p.usedExchange = true
		}
	} else {
		p.count--
		if p.count == 0 {
			fmt.Println("before exchange back")
			_ = p.MeExchangeYourCard(p.who)
		}
	}

	// 2. show
	toPlay := 0 // TODO: input
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

func NewHumanPlayer(id int) *HumanPlayer {
	return &HumanPlayer{id: id, count: 3, usedExchange: false}
}

func randomBool() bool {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) == 1
}
