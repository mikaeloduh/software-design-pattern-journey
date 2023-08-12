package service

import (
	"bigtwogame/bigtwo/entity"
	"bigtwogame/template"
	"fmt"
)

type BigTwoGame struct {
	Players       []entity.IBigTwoPlayer
	Deck          *entity.BigTwoDeck
	TopCards      []entity.BigTwoCard
	CurrentPlayer int
	Passed        int
}

func NewBigTwoGame(deck *entity.BigTwoDeck, players []entity.IBigTwoPlayer) *template.GameFramework[entity.BigTwoCard] {
	game := &template.GameFramework[entity.BigTwoCard]{
		Deck:        &deck.Deck,
		Players:     make([]template.IPlayer[entity.BigTwoCard], len(players)),
		NumCard:     13,
		PlayingGame: &BigTwoGame{Players: players, Deck: deck},
	}
	for i, player := range players {
		game.Players[i] = player
	}

	return game
}

func (b *BigTwoGame) Init() {
	b.SetActionCards()
}

func (b *BigTwoGame) SetActionCards() {
	for _, v := range b.Players {
		v.SetActionCard(entity.PassCard())
	}
}

func (b *BigTwoGame) PreTakeTurns() {
	b.TopCards = []entity.BigTwoCard{entity.InitCard()}

	for i, p := range b.Players {
		if b.hasValidPreTakeTurnMove(p.GetHand()) {
			var playerTakeTurn func() []entity.BigTwoCard

			playerTakeTurn = func() []entity.BigTwoCard {
				move := p.TakeTurnMove()
				if !b.isValidTurnMove(move.DryRun()) {
					return playerTakeTurn()
				}
				return move.Play()
			}

			b.CurrentPlayer = i
			b.updateDeskCard(playerTakeTurn())
			break
		}
	}
	b.UpdateGameAndMoveToNext()
}

func (b *BigTwoGame) TakeTurnStep(player template.IPlayer[entity.BigTwoCard]) {
	var playerPlay func() []entity.BigTwoCard

	playerPlay = func() []entity.BigTwoCard {
		move := player.(entity.IBigTwoPlayer).TakeTurnMove()
		if !b.isValidTurnMove(move.DryRun()) {
			return playerPlay()
		}
		return move.Play()
	}

	b.updateDeskCard(playerPlay())
}

func (b *BigTwoGame) GetCurrentPlayer() template.IPlayer[entity.BigTwoCard] {
	return b.Players[b.CurrentPlayer]
}

func (b *BigTwoGame) UpdateGameAndMoveToNext() {
	b.CurrentPlayer = (b.CurrentPlayer + 1) % len(b.Players)

	if b.Passed == len(b.Players)-1 {
		b.TopCards = []entity.BigTwoCard{entity.PassCard()}
		b.Passed = 0
	}
}

func (b *BigTwoGame) IsGameFinished() bool {
	for _, player := range b.Players {
		if len(player.GetHand()) == 0 {
			return true
		}
	}
	return false
}

func (b *BigTwoGame) GameResult() (winner template.IPlayer[entity.BigTwoCard]) {
	for _, player := range b.Players {
		if len(player.GetHand()) == 0 {
			fmt.Printf("%s is the winner!\n", player.GetName())
			winner = player
		}
	}
	return winner
}

// Privates
func (b *BigTwoGame) hasValidPreTakeTurnMove(hand []entity.BigTwoCard) bool {
	for _, card := range hand {
		if b.isValidTurnMove([]entity.BigTwoCard{card}) {
			return true
		}
	}
	return false
}

func (b *BigTwoGame) isValidTurnMove(playCards []entity.BigTwoCard) bool {
	played := b.Deck.AllPatternConstructor().Do(playCards)
	if played == nil {
		return false
	}
	top := b.Deck.AllPatternConstructor().Do(b.TopCards)

	return b.Deck.AllPatternComparator().Do(top, played)
}

//func (b *BigTwoGame) isValidTurnMove(playCards []entity.BigTwoCard) bool {
//	// get current deskCards status (initCard, allPassed, passCard, others...)
//	// is playCards a valid pattern? (the validator CoR)
//	// is playCards pattern greater than deskCards? ()
//	status := typeofTop(b.TopCards)
//	switch status {
//	case "InitCard":
//		// InitCardValidator (CoR)
//		b.Deck.InitCardValidator().Do(playCards)
//	case "PassCard":
//		// AllPassedHandler (CoR)
//		b.Deck.AllPassValidator().Do(playCards)
//	default:
//		p := b.Deck.CardPatternHandler() // (CoR)
//
//		return p.Do(b.TopCards, playCards)
//	}
//
//	return b.Deck.PatternHandler().Do(b.TopCards, playCards)
//}

func typeofTop(cards []entity.BigTwoCard) string {
	if len(cards) == 1 && cards[0] == entity.PassCard() {
		return "PassCard"
	} else if len(cards) == 1 && cards[0] == entity.InitCard() {
		return "InitCard"
	} else {
		return "Default"
	}
}

//func (b *BigTwoGame) isValidTurnMove(playCards []entity.BigTwoCard) bool {
//	return b.Deck.PatternHandler().Do(b.TopCards, playCards)
//}

func (b *BigTwoGame) updateDeskCard(cards []entity.BigTwoCard) {
	if cards[0] == entity.PassCard() {
		b.Passed++
		return
	}
	b.TopCards = cards
}
