package main

import (
	"cardgameframework/uno/entity"
	"cardgameframework/uno/service"
)

func main() {
	players := []entity.IUnoPlayer[entity.UnoCard]{
		//&entity.HumanUnoPlayer{name: "Player 1"},
		//&entity.HumanUnoPlayer{name: "Player 2"},
		//&entity.HumanUnoPlayer{name: "Player 3"},
		//&entity.HumanUnoPlayer{name: "Player 4"},
		&entity.AiUnoPlayer{Name: "Computer 1"},
		&entity.AiUnoPlayer{Name: "Computer 2"},
		&entity.AiUnoPlayer{Name: "Computer 3"},
		&entity.AiUnoPlayer{Name: "Computer 4"},
	}

	uno := service.NewUnoGame(players)

	uno.Run()
}
