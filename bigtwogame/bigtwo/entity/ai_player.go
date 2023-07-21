package entity

import (
	"bigtwogame/template"
	"math/rand"
)

type IBigTwoPlayer interface {
	template.IPlayer[BigTwoCard]
	TakeTurnMove() *TurnMove
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

func (a *AiBigTwoPlayer) RemoveCard(index int) BigTwoCard {
	card := a.Hand[index]
	a.Hand = append(a.Hand[:index], a.Hand[index+1:]...)
	return card
}

func (a *AiBigTwoPlayer) TakeTurnMove() *TurnMove {
	selectCard := rand.Intn(len(a.Hand))

	return NewTurnMove(&a.Hand, selectCard)
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
