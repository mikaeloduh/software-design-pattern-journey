package entity

type IPlayer interface {
	Id() int
	SetId(int)
	Name() string
	SetName(name string)
	Rename()
	AssignCard(card ShowdownCard)
	TakeTurn(players []IPlayer) ShowdownCard
	AddPoint()
	Point() int
	IInput
	IOutput
}
