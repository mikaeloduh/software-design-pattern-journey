package template

import (
	"cardgameframework/showdown/entity"
)

type Game interface {
	Run()
	ShuffleDeck()
	DrawHands()
	TakeTurns()
	GameResult() entity.IPlayer
}
