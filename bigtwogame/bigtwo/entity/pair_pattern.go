package entity

import "reflect"

// PairPattern
type PairPattern CardPattern

func NewPairPattern(cards []BigTwoCard) PairPattern {
	if isMatchPair(cards) {
		return cards
	}
	return nil
}

func (p PairPattern) Compare(tar ICardPattern) bool {
	// subject greater than target -> true
	//if p[0].Compare(tar.GetThis()[0]) == 1 || p[0].Compare(tar.GetThis()[1]) == 1 || p[1].Compare(tar.GetThis()[0]) == 1 || p[1].Compare(tar.GetThis()[1]) == 1 {
	//	return true
	//}
	//return false

	return comparePair(p, tar.GetThis())
}

func (p PairPattern) GetThis() CardPattern {
	return CardPattern(p)
}

type PairPatternConstructor struct {
	Next IPatternConstructor
}

func (h PairPatternConstructor) Do(cards []BigTwoCard) ICardPattern {
	if p := NewPairPattern(cards); p != nil {
		return p
	} else if h.Next != nil {
		return h.Next.Do(cards)
	} else {
		return nil
	}
}

type PairComparator struct {
	Next IPatternComparator
}

func (v PairComparator) Do(top ICardPattern, played ICardPattern) bool {
	if reflect.TypeOf(top) == reflect.TypeOf(PairPattern{}) && reflect.TypeOf(top) == reflect.TypeOf(played) {
		return played.Compare(top)
	} else if v.Next != nil {
		return v.Next.Do(top, played)
	} else {
		return false
	}
}

func isMatchPair(cards []BigTwoCard) bool {
	if len(cards) == 2 && cards[0].Rank == cards[1].Rank {
		return true
	}
	return false
}
