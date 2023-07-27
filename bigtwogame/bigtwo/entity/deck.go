package entity

import (
	"bigtwogame/template"
	"sort"
)

type BigTwoDeck struct {
	template.Deck[BigTwoCard]
}

// NewBigTwoDeck contains BigTwoCard
func NewBigTwoDeck() *BigTwoDeck {
	deck := &BigTwoDeck{}
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for rank := Three; rank <= Two; rank++ {
			deck.Cards = append(deck.Cards, BigTwoCard{Rank: rank, Suit: suit})
		}
	}
	return deck
}

func (d *BigTwoDeck) BigTwoHandler() PattenHandler {
	return InitCardHandler{
		AllPassHandler{
			PassCardHandler{
				SinglePattenComparator{
					PairPattenComparator{nil}}}}}
}

type PattenHandler interface {
	Do(topCards, playCards []BigTwoCard) bool
}

type InitCardHandler struct {
	Next PattenHandler
}

func (h InitCardHandler) Do(topCards, playCards []BigTwoCard) bool {
	if isInitCard(topCards) {
		return ClubsThreeValidator{SinglePattenValidator{PairPattenValidator{nil}}}.Do(playCards)
	} else if h.Next != nil {
		return h.Next.Do(topCards, playCards)
	} else {
		return false
	}
}

type AllPassHandler struct {
	Next PattenHandler
}

func (h AllPassHandler) Do(topCards, playCards []BigTwoCard) bool {
	if isPassCard(topCards) {
		return YouShallNotPass{SinglePattenValidator{PairPattenValidator{nil}}}.Do(playCards)
	} else if h.Next != nil {
		return h.Next.Do(topCards, playCards)
	} else {
		return false
	}
}

type PassCardHandler struct {
	Next PattenHandler
}

func (h PassCardHandler) Do(topCards, playCards []BigTwoCard) bool {
	if isPassCard(playCards) {
		return true
	} else if h.Next != nil {
		return h.Next.Do(topCards, playCards)
	} else {
		return false
	}
}

type SinglePattenComparator struct {
	Next PattenHandler
}

func (h SinglePattenComparator) Do(topCards, playCards []BigTwoCard) bool {
	if isMatchSingle(topCards) {
		return compareSingle(playCards, topCards)
	} else if h.Next != nil {
		return h.Next.Do(topCards, playCards)
	} else {
		return false
	}
}

type PairPattenComparator struct {
	Next PattenHandler
}

func (p PairPattenComparator) Do(topCards, playCards []BigTwoCard) bool {
	if isMatchPair(topCards) {
		return comparePair(playCards, topCards)
	} else if p.Next != nil {
		return p.Next.Do(topCards, playCards)
	} else {
		return false
	}
}

type PattenValidator interface {
	Do(cards []BigTwoCard) bool
}

type ClubsThreeValidator struct {
	Next PattenValidator
}

func (v ClubsThreeValidator) Do(cards []BigTwoCard) bool {
	if !hasClubsThree(cards) {
		return false
	} else if v.Next != nil {
		return v.Next.Do(cards)
	} else {
		return false
	}
}

type SinglePattenValidator struct {
	Next PattenValidator
}

func (v SinglePattenValidator) Do(cards []BigTwoCard) bool {
	if isMatchSingle(cards) {
		return true
	} else if v.Next != nil {
		return v.Next.Do(cards)
	} else {
		return false
	}
}

type PairPattenValidator struct {
	Next PattenValidator
}

func (v PairPattenValidator) Do(cards []BigTwoCard) bool {
	if isMatchPair(cards) {
		return true
	} else if v.Next != nil {
		return v.Next.Do(cards)
	} else {
		return false
	}
}

type StraightPattenValidator struct {
	Next PattenValidator
}

func (v StraightPattenValidator) Do(cards []BigTwoCard) bool {
	if isMatchStraight(cards) {
		return true
	} else if v.Next != nil {
		return v.Next.Do(cards)
	} else {
		return false
	}
}

type YouShallNotPass struct {
	Next PattenValidator
}

func (v YouShallNotPass) Do(cards []BigTwoCard) bool {
	if isPassCard(cards) {
		return false
	} else if v.Next != nil {
		return v.Next.Do(cards)
	} else {
		return false
	}
}

func isInitCard(cards []BigTwoCard) bool {
	return len(cards) == 1 && cards[0] == InitCard()
}

func isPassCard(cards []BigTwoCard) bool {
	return len(cards) == 1 && cards[0] == PassCard()
}

func isMatchSingle(cards []BigTwoCard) bool {
	return len(cards) == 1
}

func isMatchPair(cards []BigTwoCard) bool {
	if len(cards) == 2 && cards[0].Rank == cards[1].Rank {
		return true
	}
	return false
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

func hasClubsThree(cards []BigTwoCard) bool {
	return ContainsElement(cards, BigTwoCard{Suit: Clubs, Rank: Three})
}

func compareSingle(sub, tar []BigTwoCard) bool {
	if !isMatchSingle(sub) || !isMatchSingle(tar) {
		return false
	}
	return sub[0].Compare(tar[0]) == 1
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
