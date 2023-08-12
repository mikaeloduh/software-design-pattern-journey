package entity

import (
	"reflect"
)

type FullHousePattern CardPattern

func NewFullHousePattern(cards []BigTwoCard) FullHousePattern {
	if isMatchFullHouse(cards) {
		return cards
	}
	return nil
}

func (p FullHousePattern) Compare(tar ICardPattern) bool {
	return compareFullHouse(p, tar.GetThis())
}

func (p FullHousePattern) GetThis() CardPattern {
	return CardPattern(p)
}

type FullHousePatternConstructor struct {
	Next IPatternConstructor
}

func (h FullHousePatternConstructor) Do(cards []BigTwoCard) ICardPattern {
	if p := NewFullHousePattern(cards); p != nil {
		return p
	} else if h.Next != nil {
		return h.Next.Do(cards)
	} else {
		return nil
	}
}

type FullHouseComparator struct {
	Next IPatternComparator
}

func (v FullHouseComparator) Do(top ICardPattern, played ICardPattern) bool {
	if reflect.TypeOf(top) == reflect.TypeOf(FullHousePattern{}) && reflect.TypeOf(top) == reflect.TypeOf(played) {
		return played.Compare(top)
	} else if v.Next != nil {
		return v.Next.Do(top, played)
	} else {
		return false
	}
}
