package template

type IPlayer[T ICard] interface {
	SetCard(card T)
	TakeTurn() T
	GetName() string
	GetHand() []T
}
