package service

import "showdown/entity"

type Game struct {
	Players []entity.IPlayer
	deck    entity.Deck
}

func (g *Game) Init() {
	g.deck.Shuffle()
}

func (g *Game) DrawLoop() {
	for i := 0; i < 13; i++ {
		g.Players[0].GetDrawCard(&g.deck)
		g.Players[1].GetDrawCard(&g.deck)
		g.Players[2].GetDrawCard(&g.deck)
		g.Players[3].GetDrawCard(&g.deck)
	}
}

func NewGame(p1 entity.IPlayer, p2 entity.IPlayer, p3 entity.IPlayer, p4 entity.IPlayer, deck *entity.Deck) *Game {
	return &Game{
		Players: []entity.IPlayer{p1, p2, p3, p4},
		deck:    *deck,
	}
}
