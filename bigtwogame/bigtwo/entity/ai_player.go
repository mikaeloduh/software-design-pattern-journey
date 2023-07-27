package entity

import (
	"bigtwogame/template"
	"math/rand"
)

type IBigTwoPlayer interface {
	template.IPlayer[BigTwoCard]
	TakeTurnMove() *TurnMove
	SetActionCard(card BigTwoCard)
}

type AiBigTwoPlayer struct {
	Name        string
	Hand        []BigTwoCard
	ActionCards []BigTwoCard
}

func (a *AiBigTwoPlayer) GetName() string {
	return a.Name
}

func (a *AiBigTwoPlayer) Rename() {
	//TODO implement me
	panic("implement me")
}

func (a *AiBigTwoPlayer) AddPoint() {
	//TODO implement me
	panic("implement me")
}

func (a *AiBigTwoPlayer) GetHand() []BigTwoCard {
	return a.Hand
}

func (a *AiBigTwoPlayer) SetCard(card BigTwoCard) {
	a.Hand = append(a.Hand, card)
}

func (a *AiBigTwoPlayer) SetActionCard(card BigTwoCard) {
	a.ActionCards = append(a.ActionCards, card)
}

func (a *AiBigTwoPlayer) RemoveCard(index int) BigTwoCard {
	card := a.Hand[index]
	a.Hand = append(a.Hand[:index], a.Hand[index+1:]...)
	return card
}

func (a *AiBigTwoPlayer) TakeTurnMove() *TurnMove {
	passibleMoves := findPassibleMove(a.Hand, a.ActionCards)
	selectedMove := passibleMoves[rand.Intn(len(passibleMoves))]

	return NewTurnMove(&a.Hand, selectedMove)
}

func findPassibleMove(hand, actionCards []BigTwoCard) [][]BigTwoCard {
	var moves [][]BigTwoCard
	moves = append(moves, findSingle(hand)...)
	moves = append(moves, findPair(hand)...)
	moves = append(moves, actionCards)

	return moves
}

func findSingle(cards []BigTwoCard) [][]BigTwoCard {
	var singles [][]BigTwoCard
	for _, card := range cards {
		singles = append(singles, []BigTwoCard{card})
	}

	return singles
}

func findPair(cards []BigTwoCard) [][]BigTwoCard {
	var pairs [][]BigTwoCard
	for i, cardA := range cards {
		for j, cardB := range cards {
			if i != j && cardA.Rank == cardB.Rank {
				pair := []BigTwoCard{cardA, cardB}
				pairs = append(pairs, pair)
			}
		}
	}

	return pairs
}
