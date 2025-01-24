package entity

type SinglePattern CardPattern

func NewSinglePattern(cards []BigTwoCard) SinglePattern {
	if isMatchSingle(cards) {
		return cards
	}
	return nil
}

func (p SinglePattern) Compare(tar ICardPattern) bool {
	return compareSingle(p, tar.This())
}

func (p SinglePattern) This() CardPattern {
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

type SinglePatternComparator struct {
	Next IPatternComparator
}

func (v SinglePatternComparator) Do(top ICardPattern, played ICardPattern) bool {
	if IsSameType(top, SinglePattern{}) && IsSameType(top, played) {
		return played.Compare(top)
	} else if v.Next != nil {
		return v.Next.Do(top, played)
	} else {
		return false
	}
}

func isMatchSingle(cards []BigTwoCard) bool {
	return len(cards) == 1
}

func compareSingle(sub, tar []BigTwoCard) bool {
	if !isMatchSingle(sub) || !isMatchSingle(tar) {
		return false
	}
	return sub[0].Compare(tar[0]) == 1
}
