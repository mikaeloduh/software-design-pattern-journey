package entity

type CardPattern []BigTwoCard

type ICardPattern interface {
	Compare(target ICardPattern) bool
	This() CardPattern
}

// IPatternConstructor CoR interface
type IPatternConstructor interface {
	Do(cards []BigTwoCard) ICardPattern
}

// IPatternComparator CoR interface
type IPatternComparator interface {
	Do(top ICardPattern, played ICardPattern) bool
}
