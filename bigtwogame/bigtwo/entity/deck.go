package entity

import "bigtwogame/template"

// NewBigTwoDeck contains BigTwoCard
func NewBigTwoDeck() *template.Deck[BigTwoCard] {
	deck := &template.Deck[BigTwoCard]{}
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for rank := Three; rank <= Two; rank++ {
			deck.Cards = append(deck.Cards, BigTwoCard{Rank: rank, Suit: suit})
		}
	}
	return deck
}

// Pattern Handler
// challenger, topCard --> [IF both valid single, DO compare] --> [IF both valid pair, DO compare]
// challenger, null    --> [IF topCard is null & challenger valid single, DO pass]
//                     --> [IF topCard is null $ challenger valid pair, DO pass]
// challenger, Club 3  --> [IF topCard is Club 3 & challenger is Club 3, DO pass]
//                     --> [IF topCard is Club 3 & challenger contains C3 and valid pair, DO pass]

/**
func (h TemplateHandler) PatternHandler(c1Ptr, c2Ptr *Sprite) bool {
	if h.isMatch(deskCard) {
		return compare()
	} else if isInitCard(deskCard) {

	} else if h.Next != nil {
		h.Next.Handle(c1Ptr, c2Ptr)
	}
}
*/

/* single pattern handler

# IF deskCard IS single
if isMatch(deskCard):
	if isMatch(playCard):
		return playCard.compare(deskCard) == 1
	elif hasMatch(playCard):
		compare(deskCard, playCard)


# IF deskCard IS InitCard
if isInitCard(deskCard):
	if isMatch(playCard):
		return playCard.compare(Club3) == 0

# IF deskCard IS PassCard
if isPassCard(deskCard):
	if isMatch(playCard):
		return True


func isMatch(card) :
	if len(card)
	if match pattern ...

func hasMatch(card) :



*/
