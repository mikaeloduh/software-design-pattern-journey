package main

import (
	"cardgameframework/showdown/entity"
	"cardgameframework/showdown/service"
)

func main() {
	p1 := entity.NewAIPlayer(entity.AIPlayerInput{}, entity.AIPlayerOutput{})
	p2 := entity.NewAIPlayer(entity.AIPlayerInput{}, entity.AIPlayerOutput{})
	p3 := entity.NewAIPlayer(entity.AIPlayerInput{}, entity.AIPlayerOutput{})
	p4 := entity.NewAIPlayer(entity.AIPlayerInput{}, entity.AIPlayerOutput{})
	//p1 := entity.NewHumanPlayer(entity.HumanPlayerInput{}, entity.HumanPlayerOutput{})
	//p2 := entity.NewHumanPlayer(entity.HumanPlayerInput{}, entity.HumanPlayerOutput{})
	//p3 := entity.NewHumanPlayer(entity.HumanPlayerInput{}, entity.HumanPlayerOutput{})
	//p4 := entity.NewAIPlayer(entity.AIPlayerInput{}, entity.AIPlayerOutput{})

	deck := entity.NewDeck()
	game := service.NewShowdownGame(p1, p2, p3, p4, deck)

	game.Run()
}
