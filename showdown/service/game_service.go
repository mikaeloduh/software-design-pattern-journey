package service

import "showdown/entity"

type Game struct {
	Players []entity.IPlayer
	deck    entity.Deck
}

func NewGame(p1 entity.IPlayer, p2 entity.IPlayer, p3 entity.IPlayer, p4 entity.IPlayer, deck *entity.Deck) *Game {
	return &Game{
		Players: []entity.IPlayer{p1, p2, p3, p4},
		deck:    *deck,
	}
}
