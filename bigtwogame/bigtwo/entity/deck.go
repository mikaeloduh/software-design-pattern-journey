package entity

import (
	"bigtwogame/template"
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

func (d *BigTwoDeck) PatternConstructor() IPatternConstructor {
	return InitCardConstructor{
		PassCardConstructor{
			SinglePatternConstructor{
				PairPatternConstructor{
					StraightPatternConstructor{
						FullHousePatternConstructor{nil}}}}}}
}

func (d *BigTwoDeck) PatternComparator() IPatternComparator {
	return InitCardComparator{
		AllPassComparator{
			PassCardComparator{
				SinglePatternComparator{
					PairPatternComparator{
						StraightPatternComparator{
							FullHousePatternComparator{nil}}}}}}}
}
