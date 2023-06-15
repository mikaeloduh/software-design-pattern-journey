package entity

type IShowdownPlayer interface {
	Id() int
	SetId(int)
	Name() string
	SetName(name string)
	Rename()
	AssignCard(card ShowdownCard)
	TakeTurn(players []IShowdownPlayer) ShowdownCard
	AddPoint()
	Point() int
	IShowdownInput
	IShowdownOutput
}
