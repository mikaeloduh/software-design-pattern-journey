package entity

import "cardgameframework/template"

// UnoPlayer defines the methods required for a player in the UNO game.
//type UnoPlayer[T UnoCard] interface {
//	Rename()
//	SetCard(card T)
//	TakeTurn() T
//	GetName() string
//	GetHand() []T
//}

// UnoPlayer defines the methods required for a player in the UNO game.
type UnoPlayer[T UnoCard] interface {
	template.IPlayer[UnoCard]
}
