package entity

type IPlayer interface {
	Id() int
	SetId(int)
	Name() string
	SetName(name string)
	Rename()
	AssignCard(card Card)
	TakeTurn(players []IPlayer) Card
	AddPoint()
	Point() int
	IPlayerInput
	IPlayerOutput
}
