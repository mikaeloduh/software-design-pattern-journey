package template

type IPlayer[T ICard] interface {
	Rename()
	SetCard(card T)
	TakeTurn() T
	GetName() string
	GetHand() []T
	AddPoint()
}
