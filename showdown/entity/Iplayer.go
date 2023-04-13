package entity

type IPlayer interface {
	Id() int
	Name() string
	GetCard(card Card)
	TakeTurn() Card
	AddPoint()
	Point() int
}
