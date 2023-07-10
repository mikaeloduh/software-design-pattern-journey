package service

import "bigtwogame/bigtwo/entity"

type BigTwoGame struct {
	Players []entity.IBigTwoPlayer[entity.BigTwoCard]
}

func NewBigTwoGame(players []entity.IBigTwoPlayer[entity.BigTwoCard]) *BigTwoGame {
	return &BigTwoGame{players}
}
