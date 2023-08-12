package entity

import "reflect"

type InitCardPattern CardPattern

func NewInitCardPattern(cards []BigTwoCard) InitCardPattern {
	if isInitCard(cards) {
		return cards
	}
	return nil
}

func (p InitCardPattern) Compare(tar ICardPattern) bool {
	return true
}

func (p InitCardPattern) GetThis() CardPattern {
	return CardPattern(p)
}

type InitCardConstructor struct {
	Next IPatternConstructor
}

func (h InitCardConstructor) Do(cards []BigTwoCard) ICardPattern {
	if p := NewInitCardPattern(cards); p != nil {
		return p
	} else if h.Next != nil {
		return h.Next.Do(cards)
	} else {
		return nil
	}
}

type InitCardComparator struct {
	Next IPatternComparator
}

func (v InitCardComparator) Do(top ICardPattern, played ICardPattern) bool {
	if reflect.TypeOf(top) == reflect.TypeOf(InitCardPattern{}) {
		return ContainsElement(played.GetThis(), BigTwoCard{Suit: Clubs, Rank: Three})
	} else if v.Next != nil {
		return v.Next.Do(top, played)
	} else {
		return false
	}
}
