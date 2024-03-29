package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"cardgameframework/template"
	"cardgameframework/uno/entity"
)

func TestUnoGame_ShuffleDeck(t *testing.T) {
	deck := entity.NewUnoDeck()

	game := NewUnoGame(nil)
	game.ShuffleDeck()

	// Since shuffling the deck is a random process, we can only assert that the number
	// of cards remains the same after shuffling.
	assert.Len(t, game.Deck.Cards, len(deck.Cards))

	// Assert that the deck cards have been shuffled
	assert.NotEqual(t, game.Deck.Cards, entity.NewUnoDeck().Cards)
}

func TestUnoGame_DealHands(t *testing.T) {
	players := []entity.IUnoPlayer[entity.UnoCard]{
		&entity.HumanUnoPlayer{Name: "UnoPlayer 1"},
		&entity.AiUnoPlayer{Name: "Computer 1"},
		&entity.AiUnoPlayer{Name: "Computer 2"},
		&entity.AiUnoPlayer{Name: "Computer 3"},
	}

	game := NewUnoGame(players)
	numCards := 5
	game.DrawHands(numCards)

	// Each player should have received 5 cards, dealt cards should be removed from deck.
	assert.Len(t, players[0].GetHand(), numCards)
	assert.Len(t, players[1].GetHand(), numCards)
	assert.Len(t, game.Deck.Cards, 20)
}

func TestUnoGame_Result(t *testing.T) {
	players := []entity.IUnoPlayer[entity.UnoCard]{
		&entity.AiUnoPlayer{Name: "Computer 1"},
		&entity.AiUnoPlayer{Name: "Computer 2"},
		&entity.AiUnoPlayer{Name: "Computer 3"},
		&entity.AiUnoPlayer{Name: "Computer 4"},
	}

	game := NewUnoGame(players)
	game.ShuffleDeck()
	game.DrawHands(5)
	game.PreTakeTurns()
	game.TakeTurns()
	winner := game.GameResult()

	// UnoPlayer who won the game should have their hand empty
	assert.Equal(t, 0, len(winner.GetHand()))
}

func TestGameLogic(t *testing.T) {
	players := []entity.IUnoPlayer[entity.UnoCard]{
		&entity.AiUnoPlayer{Name: "Computer 1"},
		&entity.AiUnoPlayer{Name: "Computer 2"},
		&entity.AiUnoPlayer{Name: "Computer 3"},
		&entity.AiUnoPlayer{Name: "Computer 4"},
	}
	deck := entity.NewUnoDeck()
	playingGame := &UnoGame[entity.UnoCard]{Players: players, Deck: deck}
	game := &template.GameFramework[entity.UnoCard]{
		Deck:        deck,
		Players:     make([]template.IPlayer[entity.UnoCard], len(players)),
		NumCard:     5,
		PlayingGame: playingGame,
	}
	for i, player := range players {
		game.Players[i] = player
	}

	game.ShuffleDeck()
	game.DrawHands(5)
	game.PreTakeTurns()

	expectedDeskCard := playingGame.DeskCard

	p := playingGame.GetCurrentPlayer()
	playingGame.TakeTurnStep(p)

	actualDeskCard := playingGame.DeskCard

	// Played card should be a valid move
	if expectedDeskCard.Value != actualDeskCard.Value && expectedDeskCard.Color != actualDeskCard.Color {
		t.Errorf("invalid move: expected: %v, actual: %v", expectedDeskCard, actualDeskCard)
	}
}
