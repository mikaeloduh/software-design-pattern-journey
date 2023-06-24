package entity

import "cardgameframework/template"

// Deck represents the UNO deck.
//type Deck[T Card] struct {
//	Cards []T
//}

func NewUnoDeck() template.Deck[UnoCard] {
	deck := template.Deck[UnoCard]{}
	for _, color := range []Color{Red, Blue, Green, Yellow} {
		for value := Zero; value <= Nine; value++ {
			card := UnoCard{Color: color, Value: value}
			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}

//// Shuffle randomly shuffles the deck of cards.
//func (d *Deck[T]) Shuffle() {
//	rand.Seed(time.Now().UnixNano())
//	rand.Shuffle(len(d.Cards), func(i, j int) {
//		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
//	})
//}
//
//// DealCard deals a card from the deck.
//func (d *Deck[T]) DealCard() T {
//	card := d.Cards[0]
//	d.Cards = d.Cards[1:]
//	return card
//}
