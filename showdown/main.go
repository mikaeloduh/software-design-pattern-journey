package main

import (
	"showdown/entity"
	"showdown/service"
)

func main() {
	const (
		userId1 int = iota
		userId2
		userId3
		userId4
	)
	p1 := entity.NewHumanPlayer(userId1, entity.UserInput{})
	p2 := entity.NewHumanPlayer(userId2, entity.UserInput{})
	p3 := entity.NewHumanPlayer(userId3, entity.UserInput{})
	p4 := entity.NewHumanPlayer(userId4, entity.AIInput{})
	deck := entity.NewDeck()

	game := service.NewGame(p1, p2, p3, p4, deck)

	game.Run()
}
