package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"bigtwogame/bigtwo/entity"
	"bigtwogame/template"
)

func TestBigTwo(t *testing.T) {
	t.Parallel()

	t.Run("New game success and have 4 players", func(t *testing.T) {
		players := NewPlayers()
		game := NewBigTwoGame(players)

		assert.Equal(t, 4, len(game.Players))
	})

	t.Run("New a Deck and have it shuffled", func(t *testing.T) {
		deck := entity.NewBigTwoDeck()

		assert.Equal(t, 52, len(deck.Cards))

		deck.Shuffle()

		assert.NotEqual(t, entity.NewBigTwoDeck(), deck)
	})

	t.Run("New game and have card deal to all players", func(t *testing.T) {
		players := NewPlayers()
		game := NewBigTwoGame(players)

		game.DrawHands(game.NumCard)

		assert.Equal(t, 13, len(game.Players[0].GetHand()))
		assert.Equal(t, 13, len(game.Players[1].GetHand()))
		assert.Equal(t, 13, len(game.Players[2].GetHand()))
		assert.Equal(t, 13, len(game.Players[3].GetHand()))
		assert.Equal(t, 0, len(game.Deck.Cards))
	})

	t.Run("PreTakeTurn should play ♣3 from whoever had (single only)", func(t *testing.T) {
		players := NewPlayers()
		deck := entity.NewBigTwoDeck()
		playingGame := &BigTwoGame{Players: players, Deck: deck}
		game := &template.GameFramework[entity.BigTwoCard]{
			Deck:        deck,
			Players:     make([]template.IPlayer[entity.BigTwoCard], len(players)),
			NumCard:     13,
			PlayingGame: playingGame,
		}
		for i, player := range players {
			game.Players[i] = player
		}

		game.ShuffleDeck()
		game.DrawHands(game.NumCard)
		game.PreTakeTurns()

		assert.Equal(t, entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Three}, playingGame.DeskCard)
		//assert.Equal(t, 13-1, len(playingGame.GetCurrentPlayer().GetHand()))
	})

	t.Run("TakeTurnStep should respect the rule (single only)", func(t *testing.T) {
		players := NewPlayers()
		deck := entity.NewBigTwoDeck()
		playingGame := &BigTwoGame{Players: players, Deck: deck}
		game := &template.GameFramework[entity.BigTwoCard]{
			Deck:        deck,
			Players:     make([]template.IPlayer[entity.BigTwoCard], len(players)),
			NumCard:     13,
			PlayingGame: playingGame,
		}
		for i, player := range players {
			game.Players[i] = player
		}

		game.ShuffleDeck()
		game.DrawHands(game.NumCard)
		game.PreTakeTurns()

		firstTop := playingGame.DeskCard

		p := playingGame.GetCurrentPlayer()
		playingGame.TakeTurnStep(p)

		assert.Equal(t, 1, playingGame.DeskCard.Compare(firstTop))
		assert.Len(t, playingGame.GetCurrentPlayer().GetHand(), 13-1)
	})

	t.Run("TakeTurnStep no valid card should pass (single only)", func(t *testing.T) {
		players := NewPlayers()
		deck := entity.NewBigTwoDeck()
		playingGame := &BigTwoGame{Players: players, Deck: deck}
		game := &template.GameFramework[entity.BigTwoCard]{
			Deck:        deck,
			Players:     make([]template.IPlayer[entity.BigTwoCard], len(players)),
			NumCard:     13,
			PlayingGame: playingGame,
		}
		for i, player := range players {
			game.Players[i] = player
		}

		playingGame.DeskCard = entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Ten}
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Nine})
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Eight})

		firstTop := playingGame.DeskCard

		p := playingGame.GetCurrentPlayer()
		playingGame.TakeTurnStep(p)

		assert.Equal(t, 1, playingGame.Passed)
		assert.Equal(t, firstTop, playingGame.DeskCard)
		assert.Len(t, playingGame.GetCurrentPlayer().GetHand(), 2)
	})
}

func NewPlayers() []entity.IBigTwoPlayer {
	return []entity.IBigTwoPlayer{
		&entity.AiBigTwoPlayer{Name: "Computer 1"},
		&entity.AiBigTwoPlayer{Name: "Computer 2"},
		&entity.AiBigTwoPlayer{Name: "Computer 3"},
		&entity.AiBigTwoPlayer{Name: "Computer 4"},
	}
}
