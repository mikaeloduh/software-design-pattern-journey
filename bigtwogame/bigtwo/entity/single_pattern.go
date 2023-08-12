package entity

import (
	"reflect"
)

type SinglePattern CardPattern

func NewSinglePattern(cards []BigTwoCard) SinglePattern {
	if isMatchSingle(cards) {
		return cards
	}
	return nil
}

func (p SinglePattern) Compare(tar ICardPattern) bool {
	return compareSingle(p, tar.GetThis())
}

func (p SinglePattern) GetThis() CardPattern {
	return CardPattern(p)
}

type SinglePatternConstructor struct {
	Next IPatternConstructor
}

func (h SinglePatternConstructor) Do(cards []BigTwoCard) ICardPattern {
	if p := NewSinglePattern(cards); p != nil {
		return p
	} else if h.Next != nil {
		return h.Next.Do(cards)
	} else {
		return nil
	}
}

type SingleComparator struct {
	Next IPatternComparator
}

func (v SingleComparator) Do(top ICardPattern, played ICardPattern) bool {
	if reflect.TypeOf(top) == reflect.TypeOf(SinglePattern{}) && reflect.TypeOf(top) == reflect.TypeOf(played) {
		return played.Compare(top)
	} else if v.Next != nil {
		return v.Next.Do(top, played)
	} else {
		return false
	}
}
