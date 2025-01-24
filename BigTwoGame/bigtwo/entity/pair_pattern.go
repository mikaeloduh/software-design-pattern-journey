package entity

// PairPattern
type PairPattern CardPattern

func NewPairPattern(cards []BigTwoCard) PairPattern {
	if isMatchPair(cards) {
		return cards
	}
	return nil
}

func (p PairPattern) Compare(tar ICardPattern) bool {
	return comparePair(p, tar.This())
}

func (p PairPattern) This() CardPattern {
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

type PairPatternComparator struct {
	Next IPatternComparator
}

func (v PairPatternComparator) Do(top ICardPattern, played ICardPattern) bool {
	if IsSameType(top, PairPattern{}) && IsSameType(top, played) {
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

func comparePair(subject, target []BigTwoCard) bool {
	// subject greater than target -> true
	if !isMatchPair(subject) || !isMatchPair(target) {
		return false
	}
	if subject[0].Compare(target[0]) == 1 || subject[0].Compare(target[1]) == 1 || subject[1].Compare(target[0]) == 1 || subject[1].Compare(target[1]) == 1 {
		return true
	}
	return false
}
