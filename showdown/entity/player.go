package entity

import "errors"

type HumanPlayer struct {
	id        int
	name      string
	HandCards []Card
	point     int
}

func (p *HumanPlayer) YouExchangeMyCard(card Card) (Card, error) {
	if len(p.HandCards) < 1 {
		return Card{}, errors.New("not enough card")
	}

	// TODO: Choose a card
	myCard := p.HandCards[0]
	p.HandCards[0] = card

	return myCard, nil
}

func (p *HumanPlayer) MeExchangeYourCard(player IPlayer) error {
	if len(p.HandCards) < 1 {
		return errors.New("not enough card")
	}

	// TODO: Choose a card
	c := p.HandCards[0]

	ex, err := player.YouExchangeMyCard(c)
	if err != nil {
		return errors.New("not enough card")
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

func (p *HumanPlayer) TakeTurn() Card {
	// 1. exchange?

	// 2. show
	play := 0
	showCard := p.HandCards[play]
	p.HandCards = append([]Card{}, append(p.HandCards[0:play], p.HandCards[play+1:]...)...)

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
	return &HumanPlayer{id: id}
}
