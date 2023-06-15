package service

import "cardgameframework/entity"

type Game interface {
	Run()
	Init()
	Draw()
	TakeTurn()
	GameResult() entity.IShowdownPlayer
}
