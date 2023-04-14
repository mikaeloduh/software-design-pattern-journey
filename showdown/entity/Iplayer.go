package entity

type IPlayer interface {
	Id() int
	Name() string
	GetCard(card Card)
	TakeTurn() Card
	AddPoint()
	Point() int
	MeExchangeYourCard(player IPlayer) error
	YouExchangeMyCard(card Card) (Card, error)
}
