package service

import (
	"fmt"
	"showdown/entity"
)

type Game struct {
	Players []entity.IPlayer
	Deck    *entity.Deck
}

const rounds int = 13

func (g *Game) Run() {
	g.init()
	g.drawLoop()
	g.takeTurnLoop()
	g.gameResult()
}

func (g *Game) init() {
	g.Deck.Shuffle()
}

func (g *Game) drawLoop() {
	for i := 0; i < rounds; i++ {
		for i := range g.Players {
			g.Players[i].GetCard(g.Deck.DrawCard())
		}
	}
}

func (g *Game) takeTurnLoop() {
	muckedCards := make([]entity.Card, len(g.Players))

	for i := 0; i < rounds; i++ {
		fmt.Printf("==Round %d ==\n", i)
		for i := range g.Players {
			muckedCards[i] = g.Players[i].TakeTurn(g.Players)
		}

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

	winner := g.Players[win]
	fmt.Printf("The Winner is: %s\n", winner.Name())

	return winner
}

func NewGame(p1 entity.IPlayer, p2 entity.IPlayer, p3 entity.IPlayer, p4 entity.IPlayer, deck *entity.Deck) *Game {
	return &Game{
		Players: []entity.IPlayer{p1, p2, p3, p4},
		Deck:    deck,
	}
}

// helper
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
