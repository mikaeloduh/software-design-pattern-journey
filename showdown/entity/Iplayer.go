package entity

type IPlayer interface {
	Id() int
	Name() string
	ReName()
	GetCard(card Card)
	TakeTurn(players []IPlayer) Card
	AddPoint()
	Point() int
	MeExchangeYourCard(player IPlayer) error
	YouExchangeMyCard(card Card) (Card, error)
	Input
}
