package main

import (
	"cardgameframework/showdown/entity"
	"cardgameframework/showdown/service"
)

func main() {
	players := []entity.IShowdownPlayer[entity.ShowDownCard]{
		entity.NewAIPlayer(entity.AIPlayerInput{}, entity.AIPlayerOutput{}),
		entity.NewAIPlayer(entity.AIPlayerInput{}, entity.AIPlayerOutput{}),
		entity.NewAIPlayer(entity.AIPlayerInput{}, entity.AIPlayerOutput{}),
		entity.NewAIPlayer(entity.AIPlayerInput{}, entity.AIPlayerOutput{}),
		//entity.NewHumanPlayer(entity.HumanPlayerInput{}, entity.HumanPlayerOutput{})
		//entity.NewHumanPlayer(entity.HumanPlayerInput{}, entity.HumanPlayerOutput{})
		//entity.NewHumanPlayer(entity.HumanPlayerInput{}, entity.HumanPlayerOutput{})
		//entity.NewHumanPlayer(entity.HumanPlayerInput{}, entity.HumanPlayerOutput{})
	}

	game := service.NewShowdownGame(players)

	game.Run()
}
