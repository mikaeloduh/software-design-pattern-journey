package entity

import "bigtwogame/template"

type IBigTwoPlayer[T BigTwoCard] interface {
	template.IPlayer[BigTwoCard]
}

type AiBigTwoPlayer struct {
	Name string
	Hand []BigTwoCard
}

func (a *AiBigTwoPlayer) Rename() {
	//TODO implement me
	panic("implement me")
}

func (a *AiBigTwoPlayer) SetCard(card BigTwoCard) {
	a.Hand = append(a.Hand, card)
}

func (a *AiBigTwoPlayer) TakeTurn() BigTwoCard {
	//TODO implement me
	panic("implement me")
}

func (a *AiBigTwoPlayer) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (a *AiBigTwoPlayer) GetHand() []BigTwoCard {
	return a.Hand
}

func (a *AiBigTwoPlayer) AddPoint() {
	//TODO implement me
	panic("implement me")
}
