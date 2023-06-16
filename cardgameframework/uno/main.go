package main

import (
	"cardgameframework/uno/entity"
	"cardgameframework/uno/service"
)

func main() {
	deck := entity.NewDeck()

	players := []entity.IPlayer{
		&entity.HumanPlayer{Name: "Player 1"},
		&entity.HumanPlayer{Name: "Player 2"},
		&entity.HumanPlayer{Name: "Player 3"},
		&entity.HumanPlayer{Name: "Player 4"},
		//&service.ComputerPlayer{Name: "Computer 1"},
		// Add more players here...
	}

	game := service.NewUnoGame(players, deck)
	game.ShuffleDeck()
	game.DealHands(5)
	game.TakeTurns()
	game.GameResult()
}
