package entity

import "bigtwogame/template"

type IBigTwoPlayer interface {
	template.IPlayer[BigTwoCard]
	TakeTurnMove() *TurnMove
	SetActionCard(card BigTwoCard)
}
