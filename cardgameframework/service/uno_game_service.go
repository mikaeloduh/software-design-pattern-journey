package service

import (
	"cardgameframework/entity"
)

type UnoGame struct {
	Players []entity.IUnoPlayer
	Deck    entity.UnoDeck
}

func NewUnoGame(p1 entity.IUnoPlayer, p2 entity.IUnoPlayer, p3 entity.IUnoPlayer, p4 entity.IUnoPlayer) *UnoGame {
	for i, p := range []entity.IUnoPlayer{p1, p2, p3, p4} {
		p.SetId(i)
		//p.SetName(fmt.Sprintf("P%d", i))
	}
	return &UnoGame{
		Players: []entity.IUnoPlayer{p1, p2, p3, p4},
		Deck:    entity.NewUnoDeck(),
	}
}

func (g *UnoGame) Run() {
	g.Init()
}

func (g *UnoGame) Init() {
	g.Deck.Shuffle()
}

func (g *UnoGame) Draw() {
	//TODO implement me
	panic("implement me")
}

func (g *UnoGame) TakeTurn() {
	//TODO implement me
	panic("implement me")
}

func (g *UnoGame) GameResult() entity.IShowdownPlayer {
	//TODO implement me
	panic("implement me")
}
