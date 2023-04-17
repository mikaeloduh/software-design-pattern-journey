package main

import (
	"showdown/entity"
	"showdown/service"
)

func main() {
	userInput := entity.UserInput{}
	p1 := entity.NewHumanPlayer(0, userInput)
	p2 := entity.NewHumanPlayer(1, userInput)
	p3 := entity.NewHumanPlayer(2, userInput)
	p4 := entity.NewHumanPlayer(3, entity.AIInput{})
	deck := entity.NewDeck()
	game := service.NewGame(p1, p2, p3, p4, deck)

	game.Run()
}
