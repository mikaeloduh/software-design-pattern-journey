package entity

import "bigtwogame/template"

type IShowdownPlayer interface {
	template.IPlayer[ShowDownCard]

	Id() int
	SetId(int)
	SetName(name string)
	TakeTurn() ShowDownCard
	AddPoint()
	Point() int
	IPlayerInput
	IPlayerOutput
}
