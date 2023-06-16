package main

import (
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
	deck := service.NewDeck()

	players := []service.Player{
		&service.HumanPlayer{Name: "Player 1"},
		&service.HumanPlayer{Name: "Player 2"},
		&service.HumanPlayer{Name: "Player 3"},
		//&template.HumanPlayer{Name: "Player 4"},
		&service.ComputerPlayer{Name: "Computer 1"},
		// Add more players here...
	}

	game := service.NewUnoGame(players, deck)
	game.ShuffleDeck()
	game.DealHands(7)
	game.TakeTurns()
	game.GameResult()
}
