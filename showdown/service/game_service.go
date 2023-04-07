package service

import "showdown/entity"

type Game struct {
	Players []*entity.Player
}

func NewGame(p1 *entity.Player, p2 *entity.Player, p3 *entity.Player, p4 *entity.Player) *Game {
	return &Game{
		Players: []*entity.Player{p1, p2, p3, p4},
	}
}
