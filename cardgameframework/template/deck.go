package template

type IDeck[T ICard] struct {
	Shuffle()
	DealCard() T
}
