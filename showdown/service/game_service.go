package service

import (
	"fmt"
	"showdown/entity"
)

type Game struct {
	Players []entity.IPlayer
	deck    entity.Deck
	Winner  entity.IPlayer
}

const rounds int = 13

func (g *Game) Init() {
	g.deck.Shuffle()
}

func (g *Game) DrawLoop() {
	for i := 0; i < rounds; i++ {
		g.Players[0].GetCard(g.deck.DrawCard())
		g.Players[1].GetCard(g.deck.DrawCard())
		g.Players[2].GetCard(g.deck.DrawCard())
		g.Players[3].GetCard(g.deck.DrawCard())
	}
}

func (g *Game) takeTurnLoop() {
	muckedCards := make([]entity.Card, len(g.Players))

	for i := 0; i < rounds; i++ {
		fmt.Printf("==Round %d ==\n", i)
		muckedCards[0] = g.Players[0].TakeTurn(g.Players)
		muckedCards[1] = g.Players[1].TakeTurn(g.Players)
		muckedCards[2] = g.Players[2].TakeTurn(g.Players)
		muckedCards[3] = g.Players[3].TakeTurn(g.Players)

		win := showDown(muckedCards)

		g.Players[win].AddPoint()
	}
}

func (g *Game) gameResult() entity.IPlayer {
	max := 0
	var win int
	for i, p := range g.Players {
		if p.Point() > max {
			max = p.Point()
			win = i
		}
	}
	return g.Players[win]
}

func NewGame(p1 entity.IPlayer, p2 entity.IPlayer, p3 entity.IPlayer, p4 entity.IPlayer, deck *entity.Deck) *Game {
	return &Game{
		Players: []entity.IPlayer{p1, p2, p3, p4},
		deck:    *deck,
	}
}

func showDown(cards []entity.Card) int {
	if len(cards) == 0 {
		return 0
	}
	max := cards[0]
	imax := 0
	for i, card := range cards {
		if card.IsGreater(max) {
			max = card
			imax = i
		}
	}
	return imax
}
