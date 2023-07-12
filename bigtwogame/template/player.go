package template

type IPlayer[T ICard] interface {
	Rename()
	SetCard(card T)
	RemoveCard(idx int) T
	GetName() string
	GetHand() []T
	AddPoint()
}
