package main

import (
	"cardgameframework/uno/service"
)

func main() {
	deck := service.NewDeck()

	players := []service.Player{
		&service.HumanPlayer{Name: "Player 1"},
		&service.HumanPlayer{Name: "Player 2"},
		&service.HumanPlayer{Name: "Player 3"},
		&service.HumanPlayer{Name: "Player 4"},
		//&service.ComputerPlayer{Name: "Computer 1"},
		// Add more players here...
	}

	game := service.NewUnoGame(players, deck)
	game.ShuffleDeck()
	game.DealHands(5)
	game.TakeTurns()
	game.GameResult()
}
