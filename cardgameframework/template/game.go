package template

import (
	"cardgameframework/showdown/entity"
)

type Game interface {
	Run()
	Init()
	Draw()
	TakeTurn()
	GameResult() entity.IPlayer
}
