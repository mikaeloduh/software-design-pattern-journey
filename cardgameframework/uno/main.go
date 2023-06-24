package main

import (
	"cardgameframework/template"
	"cardgameframework/uno/entity"
	"cardgameframework/uno/service"
)

func main() {
	//players := []template.IPlayer[entity.UnoCard]{
	//	&entity.HumanUnoPlayer{Name: "Player 1"},
	//	&entity.HumanUnoPlayer{Name: "Player 2"},
	//	&entity.HumanUnoPlayer{Name: "Player 3"},
	//	&entity.HumanUnoPlayer{Name: "Player 4"},
	//	//&service.ComputerPlayer{Name: "Computer 1"},
	//	// Add more players here...
	//}

	players := []entity.UnoPlayer[entity.UnoCard]{
		&entity.HumanUnoPlayer{Name: "Player 1"},
		&entity.HumanUnoPlayer{Name: "Player 2"},
		&entity.HumanUnoPlayer{Name: "Player 3"},
		&entity.HumanUnoPlayer{Name: "Player 4"},
	}
	deck := entity.NewUnoDeck()

	//game := service.NewUnoGame(players, deck)

	playGame := service.UnoGame[entity.UnoCard]{Players: players}
	base := template.GameFramework[entity.UnoCard]{
		Deck:        deck,
		Players:     make([]template.IPlayer[entity.UnoCard], len(players)), // <- Cannot use 'players' (type []entity.UnoPlayer[entity.UnoCard]) as the type []IPlayer[T]
		PlayingGame: &playGame,
	}
	for i, player := range players {
		base.Players[i] = player
	}

	base.Run()
}
