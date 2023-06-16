package entity

// IPlayer defines the methods required for a player in the UNO game.
type IPlayer interface {
	SetCard(card Card)
	TakeTurn() Card
	GetName() string
	GetHand() []Card
}
