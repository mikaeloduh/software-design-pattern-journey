package entity

type FullHousePattern CardPattern

func NewFullHousePattern(cards []BigTwoCard) FullHousePattern {
	if isMatchFullHouse(cards) {
		return cards
	}
	return nil
}

func (p FullHousePattern) Compare(tar ICardPattern) bool {
	return compareFullHouse(p, tar.This())
}

func (p FullHousePattern) This() CardPattern {
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

type FullHousePatternComparator struct {
	Next IPatternComparator
}

func (v FullHousePatternComparator) Do(top ICardPattern, played ICardPattern) bool {
	if IsSameType(top, FullHousePattern{}) && IsSameType(top, played) {
		return played.Compare(top)
	} else if v.Next != nil {
		return v.Next.Do(top, played)
	} else {
		return false
	}
}

func isMatchFullHouse(cards []BigTwoCard) bool {
	if len(cards) != 5 {
		return false
	}

	rankCount := make(map[Rank]int)
	for _, card := range cards {
		rankCount[card.Rank]++
	}

	var hasTwo, hasThree bool
	for _, count := range rankCount {
		switch count {
		case 2:
			hasTwo = true
		case 3:
			hasThree = true
		}
	}

	return hasTwo && hasThree
}

func compareFullHouse(subject, target []BigTwoCard) bool {
	if !isMatchFullHouse(subject) || !isMatchFullHouse(target) {
		return false
	}

	subjectRankCounts := make(map[Rank]int)
	for _, card := range subject {
		subjectRankCounts[card.Rank]++
	}

	targetRankCounts := make(map[Rank]int)
	for _, card := range target {
		targetRankCounts[card.Rank]++
	}

	var subjectThreeRank Rank
	for rank, count := range subjectRankCounts {
		if count == 3 {
			subjectThreeRank = rank
		}
	}

	var targetThreeRank Rank
	for rank, count := range targetRankCounts {
		if count == 3 {
			targetThreeRank = rank
		}
	}

	if subjectThreeRank > targetThreeRank {
		return true
	} else {
		return false
	}
}
