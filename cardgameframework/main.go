package main

import (
	"cardgameframework/entity"
	"cardgameframework/service"
)

func main() {
	//p1 := entity.NewAIPlayer(entity.Showdown_AIInput{}, entity.ShowdownAIOutput{})
	//p2 := entity.NewAIPlayer(entity.Showdown_AIInput{}, entity.ShowdownAIOutput{})
	//p3 := entity.NewAIPlayer(entity.Showdown_AIInput{}, entity.ShowdownAIOutput{})
	//p4 := entity.NewAIPlayer(entity.Showdown_AIInput{}, entity.ShowdownAIOutput{})

	p1 := entity.NewHumanPlayer(entity.ShowdownHumanInput{}, entity.ShowdownHumanOutput{})
	p2 := entity.NewHumanPlayer(entity.ShowdownHumanInput{}, entity.ShowdownHumanOutput{})
	p3 := entity.NewHumanPlayer(entity.ShowdownHumanInput{}, entity.ShowdownHumanOutput{})
	p4 := entity.NewAIPlayer(entity.ShowdownAIInput{}, entity.ShowdownAIOutput{})
	deck := entity.NewShowdownDeck()

	game := service.NewShowdownGame(p1, p2, p3, p4, deck)

	game.Run()
}
