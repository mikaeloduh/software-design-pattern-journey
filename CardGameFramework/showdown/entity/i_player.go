package entity

import "cardgameframework/template"

type IShowdownPlayer[T ShowDownCard] interface {
	template.IPlayer[ShowDownCard]

	Id() int
	SetId(int)
	SetName(name string)
	AddPoint()
	Point() int
	IPlayerInput
	IPlayerOutput
}
