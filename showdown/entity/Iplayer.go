package entity

type IPlayer interface {
	Id() int
	Name() string
	GetDrawCard(deck *Deck)
	TakeTurn() Card
	AddPoint()
	Point() int
}
