package entity

type IPlayer interface {
	Id() int
	Name() string
	GetDrawCard(deck *Deck)
}
