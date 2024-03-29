package main

import (
	"bigtwogame/showdown/entity"
	"bigtwogame/showdown/service"
)

func main() {
	players := []entity.IShowdownPlayer{
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
