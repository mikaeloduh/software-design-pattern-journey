package entity

import (
	"sort"
)

// StraightPattern
type StraightPattern CardPattern

func NewStraightPattern(cards []BigTwoCard) StraightPattern {
	if isMatchStraight(cards) {
		return cards
	}
	return nil
}

func (p StraightPattern) Compare(tar ICardPattern) bool {
	return compareStraight(p, tar.This())
}

func (p StraightPattern) This() CardPattern {
	return CardPattern(p)
}

type StraightPatternConstructor struct {
	Next IPatternConstructor
}

func (h StraightPatternConstructor) Do(cards []BigTwoCard) ICardPattern {
	if p := NewStraightPattern(cards); p != nil {
		return p
	} else if h.Next != nil {
		return h.Next.Do(cards)
	} else {
		return nil
	}
}

type StraightPatternComparator struct {
	Next IPatternComparator
}

func (v StraightPatternComparator) Do(top ICardPattern, played ICardPattern) bool {
	if IsSameType(top, StraightPattern{}) && IsSameType(top, played) {
		return played.Compare(top)
	} else if v.Next != nil {
		return v.Next.Do(top, played)
	} else {
		return false
	}
}

func isMatchStraight(cards []BigTwoCard) bool {
	if len(cards) < 5 {
		return false
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	for i := 0; i < len(cards)-1; i++ {
		if cards[i].Rank+1 != cards[i+1].Rank {
			return false
		}
	}
	return true
}

func compareStraight(cards []BigTwoCard, cards2 []BigTwoCard) bool {
	if !isMatchStraight(cards) || !isMatchStraight(cards2) {
		return false
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})
	sort.Slice(cards2, func(i, j int) bool {
		return cards2[i].Rank < cards2[j].Rank
	})

	return cards[len(cards)-1].Compare(cards2[len(cards2)-1]) == 1
}
