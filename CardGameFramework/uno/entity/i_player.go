package entity

import "cardgameframework/template"

// IUnoPlayer defines the methods required for a player in the UNO game.
type IUnoPlayer[T UnoCard] interface {
	template.IPlayer[UnoCard]
}
