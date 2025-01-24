package main

import (
	"showdown/entity"
	"showdown/service"
)

func main() {
	//p1 := entity.NewAIPlayer(entity.AIInput{}, entity.AIOutput{})
	//p2 := entity.NewAIPlayer(entity.AIInput{}, entity.AIOutput{})
	//p3 := entity.NewAIPlayer(entity.AIInput{}, entity.AIOutput{})
	//p4 := entity.NewAIPlayer(entity.AIInput{}, entity.AIOutput{})

	p1 := entity.NewHumanPlayer(entity.HumanInput{}, entity.HumanOutput{})
	p2 := entity.NewHumanPlayer(entity.HumanInput{}, entity.HumanOutput{})
	p3 := entity.NewHumanPlayer(entity.HumanInput{}, entity.HumanOutput{})
	p4 := entity.NewAIPlayer(entity.AIInput{}, entity.AIOutput{})
	deck := entity.NewDeck()

	game := service.NewGame(p1, p2, p3, p4, deck)

	game.Run()
}
