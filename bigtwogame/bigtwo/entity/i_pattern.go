package entity

type CardPattern []BigTwoCard

type ICardPattern interface {
	//Match(target ICardPattern) bool
	Compare(target ICardPattern) bool
	GetThis() CardPattern
}

// IPatternConstructor CoR interface
type IPatternConstructor interface {
	Do(cards []BigTwoCard) ICardPattern
}

// IPatternComparator
type IPatternComparator interface {
	Do(top ICardPattern, played ICardPattern) bool
}
