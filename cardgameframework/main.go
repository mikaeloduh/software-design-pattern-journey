package main

import (
	"cardgameframework/uno/entity"
	"cardgameframework/uno/service"
)

//func main() {
//p1 := entity.NewAIPlayer(entity.Showdown_AIInput{}, entity.ShowdownAIOutput{})
//p2 := entity.NewAIPlayer(entity.Showdown_AIInput{}, entity.ShowdownAIOutput{})
//p3 := entity.NewAIPlayer(entity.Showdown_AIInput{}, entity.ShowdownAIOutput{})
//p4 := entity.NewAIPlayer(entity.Showdown_AIInput{}, entity.ShowdownAIOutput{})

//p1 := entity.NewHumanPlayer(entity.ShowdownHumanInput{}, entity.ShowdownHumanOutput{})
//p2 := entity.NewHumanPlayer(entity.ShowdownHumanInput{}, entity.ShowdownHumanOutput{})
//p3 := entity.NewHumanPlayer(entity.ShowdownHumanInput{}, entity.ShowdownHumanOutput{})
//p4 := entity.NewAIPlayer(entity.ShowdownAIInput{}, entity.ShowdownAIOutput{})
//deck := entity.NewShowdownDeck()
//
//game := template.NewShowdownGame(p1, p2, p3, p4, deck)
//
//game.Run()
//}

func main() {
	deck := entity.NewDeck()

	players := []entity.IPlayer{
		&entity.HumanPlayer{Name: "IPlayer 1"},
		&entity.HumanPlayer{Name: "IPlayer 2"},
		&entity.HumanPlayer{Name: "IPlayer 3"},
		//&template.HumanPlayer{Name: "IPlayer 4"},
		&entity.ComputerPlayer{Name: "Computer 1"},
		// Add more players here...
	}

	game := service.NewUnoGame(players, deck)
	game.ShuffleDeck()
	game.DealHands(7)
	game.TakeTurns()
	game.GameResult()
}
