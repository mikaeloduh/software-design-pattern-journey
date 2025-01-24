package entity

import "reflect"

type PassCardPattern CardPattern

func NewPassCardPattern(cards []BigTwoCard) PassCardPattern {
	if isPassCard(cards) {
		return cards
	}
	return nil
}

func (p PassCardPattern) Compare(tar ICardPattern) bool {
	return true
}

func (p PassCardPattern) This() CardPattern {
	return CardPattern(p)
}

type PassCardConstructor struct {
	Next IPatternConstructor
}

func (h PassCardConstructor) Do(cards []BigTwoCard) ICardPattern {
	if p := NewPassCardPattern(cards); p != nil {
		return p
	} else if h.Next != nil {
		return h.Next.Do(cards)
	} else {
		return nil
	}
}

type PassCardComparator struct {
	Next IPatternComparator
}

func (v PassCardComparator) Do(top ICardPattern, played ICardPattern) bool {
	if reflect.TypeOf(played) == reflect.TypeOf(PassCardPattern{}) {
		return true
	} else if v.Next != nil {
		return v.Next.Do(top, played)
	} else {
		return false
	}
}

type AllPassComparator struct {
	Next IPatternComparator
}

func (v AllPassComparator) Do(top ICardPattern, played ICardPattern) bool {
	if reflect.TypeOf(top) == reflect.TypeOf(PassCardPattern{}) {
		if reflect.TypeOf(played) == reflect.TypeOf(PassCardPattern{}) {
			return false
		}
		return true
	} else if v.Next != nil {
		return v.Next.Do(top, played)
	} else {
		return false
	}
}

func isPassCard(cards []BigTwoCard) bool {
	return len(cards) == 1 && cards[0] == PassCard()
}
