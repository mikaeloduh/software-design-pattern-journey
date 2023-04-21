package entity

type IPlayer interface {
	Id() int
	SetId(int)
	Name() string
	SetName(name string)
	Rename()
	GetCard(card Card)
	TakeTurn(players []IPlayer) Card
	AddPoint()
	Point() int
	MeExchangeYourCard(player IPlayer) error
	YouExchangeMyCard(card Card) (Card, error)
	IInput
	IOutput
}
