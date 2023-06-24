package main

import (
	"cardgameframework/uno/entity"
	"cardgameframework/uno/service"
)

func main() {
	players := []entity.IUnoPlayer[entity.UnoCard]{
		&entity.HumanUnoPlayer{Name: "Player 1"},
		&entity.HumanUnoPlayer{Name: "Player 2"},
		&entity.HumanUnoPlayer{Name: "Player 3"},
		&entity.HumanUnoPlayer{Name: "Player 4"},
		//&entity.ComputerUnoPlayer{Name: "Computer 1"},
	}

	uno := service.NewUnoGame(players)

	uno.Run()
}
