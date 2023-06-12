package service

import "cardgameframework/entity"

type UnoGame struct {
	Deck entity.UnoDeck
}

func NewUnoGame() *UnoGame {
	return &UnoGame{Deck: entity.NewUnoDeck()}
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

func (g *UnoGame) GameResult() entity.IPlayer {
	//TODO implement me
	panic("implement me")
}
