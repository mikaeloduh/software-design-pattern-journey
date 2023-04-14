package main

import (
	"showdown/entity"
	"showdown/service"
)

func main() {
	p1 := entity.NewHumanPlayer(0, entity.UserInput{})
	p2 := entity.NewHumanPlayer(1, entity.UserInput{})
	p3 := entity.NewHumanPlayer(2, entity.UserInput{})
	p4 := entity.NewHumanPlayer(3, entity.AIInput{})
	deck := entity.NewDeck()
	game := service.NewGame(p1, p2, p3, p4, deck)

	game.Run()
}
