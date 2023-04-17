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
	g.renamePlayer()
	g.drawLoop()
	g.takeTurnLoop()
	g.gameResult()
}

func (g *Game) init() {
	g.Deck.Shuffle()
}

func (g *Game) renamePlayer() {
	for i := range g.Players {
		g.Players[i].ReName()
	}
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
		fmt.Printf("\n======== Round %d ========\n", i)
		for i := range g.Players {
			muckedCards[i] = g.Players[i].TakeTurn(g.Players)
		}

		win := showDown(muckedCards)

		g.Players[win].AddPoint()
	}
}

func (g *Game) gameResult() entity.IPlayer {
	var winner entity.IPlayer
	max := 0
	for i := range g.Players {
		if g.Players[i].Point() > max {
			max = g.Players[i].Point()
			winner = g.Players[i]
		}
	}

	fmt.Printf("\n======== Game Over ========\nThe Winner is: %s\n", winner.Name())

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
