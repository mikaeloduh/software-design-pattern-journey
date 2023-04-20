package main

import (
	"showdown/entity"
	"showdown/service"
)

func main() {
	p1 := entity.NewHumanPlayer(entity.UserInput{})
	p2 := entity.NewHumanPlayer(entity.UserInput{})
	p3 := entity.NewHumanPlayer(entity.UserInput{})
	p4 := entity.NewAIPlayer(entity.AIInput{})

	deck := entity.NewDeck()

	game := service.NewGame(p1, p2, p3, p4, deck)

	game.Run()
}
