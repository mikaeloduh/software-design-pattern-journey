package entity

import (
	"math/rand"
)

type AiBigTwoPlayer struct {
	Name        string
	Hand        []BigTwoCard
	ActionCards []BigTwoCard
}

func (p *AiBigTwoPlayer) GetName() string {
	return p.Name
}

func (p *AiBigTwoPlayer) Rename() {
	//TODO implement me
	panic("implement me")
}

func (p *AiBigTwoPlayer) AddPoint() {
	//TODO implement me
	panic("implement me")
}

func (p *AiBigTwoPlayer) GetHand() []BigTwoCard {
	return p.Hand
}

func (p *AiBigTwoPlayer) SetCard(card BigTwoCard) {
	p.Hand = append(p.Hand, card)
}

func (p *AiBigTwoPlayer) SetActionCard(card BigTwoCard) {
	p.ActionCards = append(p.ActionCards, card)
}

func (p *AiBigTwoPlayer) RemoveCard(index int) BigTwoCard {
	card := p.Hand[index]
	p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)
	return card
}

func (p *AiBigTwoPlayer) TakeTurnMove() *TurnMove {
	passibleMoves := findPassibleMove(p.Hand, p.ActionCards)
	selectedMove := passibleMoves[rand.Intn(len(passibleMoves))]

	return NewTurnMove(&p.Hand, selectedMove)
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
