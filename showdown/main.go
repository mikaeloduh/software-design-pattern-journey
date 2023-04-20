package main

import (
	"showdown/entity"
	"showdown/service"
)

func main() {
	p1 := entity.NewHumanPlayer(entity.HumanInput{})
	p2 := entity.NewHumanPlayer(entity.HumanInput{})
	p3 := entity.NewHumanPlayer(entity.HumanInput{})
	p4 := entity.NewAIPlayer(entity.AIInput{})
	//p1 := entity.NewAIPlayer(entity.AIInput{})
	//p2 := entity.NewAIPlayer(entity.AIInput{})
	//p3 := entity.NewAIPlayer(entity.AIInput{})
	//p4 := entity.NewAIPlayer(entity.AIInput{})

	deck := entity.NewDeck()

	game := service.NewGame(p1, p2, p3, p4, deck)

	game.Run()
}
