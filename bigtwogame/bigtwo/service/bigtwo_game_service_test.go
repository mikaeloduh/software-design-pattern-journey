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
		players := FakeNewPlayers()
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
		players := FakeNewPlayers()
		game := NewBigTwoGame(players)

		game.DrawHands(game.NumCard)

		assert.Equal(t, 13, len(game.Players[0].GetHand()))
		assert.Equal(t, 13, len(game.Players[1].GetHand()))
		assert.Equal(t, 13, len(game.Players[2].GetHand()))
		assert.Equal(t, 13, len(game.Players[3].GetHand()))
		assert.Equal(t, 0, len(game.Deck.Cards))
	})

	t.Run("PreTakeTurn should play ♣3 from whoever had (single only)", func(t *testing.T) {
		players := FakeNewPlayers()
		game, playingGame := FakeNewBigTwoGame(players)

		playingGame.SetActionCards()
		game.ShuffleDeck()
		game.DrawHands(game.NumCard)
		game.PreTakeTurns()

		assert.Contains(t, playingGame.TopCards, entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Three})
	})

	t.Run("Testing TakeTurn while player not pass, played card should be a valid single", func(t *testing.T) {
		players := FakeNewPlayers()
		_, playingGame := FakeNewBigTwoGame(players)
		playingGame.TopCards = []entity.BigTwoCard{{
			Suit: entity.Spades,
			Rank: entity.Eight}}
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Spades, Rank: entity.Three})
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Hearts, Rank: entity.Four})
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Diamonds, Rank: entity.Five})
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Hearts, Rank: entity.Jack})

		playingGame.CurrentPlayer = 0
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())

		assert.Len(t, players[0].GetHand(), 3)
		assert.Equal(t, []entity.BigTwoCard{{Suit: entity.Hearts, Rank: entity.Jack}}, playingGame.TopCards)
		assert.Equal(t, 0, playingGame.Passed)
	})

	t.Run("Testing TakeTurn while player not pass, played card should be a valid pair", func(t *testing.T) {
		players := FakeNewPlayers()
		_, playingGame := FakeNewBigTwoGame(players)
		playingGame.TopCards = []entity.BigTwoCard{
			{Suit: entity.Clubs, Rank: entity.Three},
			{Suit: entity.Diamonds, Rank: entity.Three}}
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.King})
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Spades, Rank: entity.King})

		playingGame.CurrentPlayer = 0
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())

		assert.Contains(t, playingGame.TopCards, entity.BigTwoCard{Suit: entity.Spades, Rank: entity.King})
		assert.Contains(t, playingGame.TopCards, entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.King})
		assert.Len(t, playingGame.GetCurrentPlayer().GetHand(), 0)
	})

	t.Run("Testing TakeTurn while player not pass, played card should be a valid pair (full)", func(t *testing.T) {
		players := FakeNewPlayers()
		game, playingGame := FakeNewBigTwoGame(players)

		game.ShuffleDeck()
		game.DrawHands(game.NumCard)
		playingGame.TopCards = []entity.BigTwoCard{
			{Suit: entity.Clubs, Rank: entity.Three},
			{Suit: entity.Diamonds, Rank: entity.Three}}

		playingGame.CurrentPlayer = 3
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())

		assert.Equal(t, true, isMatchPair(playingGame.TopCards))
		assert.Len(t, playingGame.GetCurrentPlayer().GetHand(), 13-2)
	})

	t.Run("Testing TakeTurn with no valid card player should pass (single only)", func(t *testing.T) {
		players := FakeNewPlayers()
		_, playingGame := FakeNewBigTwoGame(players)
		playingGame.SetActionCards()
		expectedTopCards := []entity.BigTwoCard{{Suit: entity.Clubs, Rank: entity.Ten}}
		playingGame.TopCards = expectedTopCards
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Nine})
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Eight})

		playingGame.CurrentPlayer = 0
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())

		assert.Equal(t, 1, playingGame.Passed)
		assert.Equal(t, expectedTopCards, playingGame.TopCards)
		assert.Len(t, playingGame.GetCurrentPlayer().GetHand(), 2)
	})

	t.Run("Testing TakeTurn, when three pass in a line, a new turn should start (single only)", func(t *testing.T) {
		players := FakeNewPlayers()
		_, playingGame := FakeNewBigTwoGame(players)

		playingGame.SetActionCards()
		// TopCards is Club 10
		playingGame.TopCards = []entity.BigTwoCard{{Suit: entity.Clubs, Rank: entity.Ten}}
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Nine})
		players[1].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Eight})
		players[2].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Seven})
		players[3].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Two})

		// Player 0
		playingGame.CurrentPlayer = 0
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())
		playingGame.UpdateGameAndMoveToNext()
		// Player 1
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())
		playingGame.UpdateGameAndMoveToNext()
		// Player 2
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())
		playingGame.UpdateGameAndMoveToNext()
		// Player 3
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())
		playingGame.UpdateGameAndMoveToNext()

		assert.Equal(t, []entity.BigTwoCard{{Suit: entity.Clubs, Rank: entity.Two}}, playingGame.TopCards)
		assert.Len(t, players[0].GetHand(), 1)
		assert.Len(t, players[1].GetHand(), 1)
		assert.Len(t, players[2].GetHand(), 1)
		assert.Len(t, players[3].GetHand(), 0)
	})
}

func FakeNewBigTwoGame(players []entity.IBigTwoPlayer) (*template.GameFramework[entity.BigTwoCard], *BigTwoGame) {
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

	return game, playingGame
}

func FakeNewPlayers() []entity.IBigTwoPlayer {
	return []entity.IBigTwoPlayer{
		&entity.AiBigTwoPlayer{Name: "Computer 1"},
		&entity.AiBigTwoPlayer{Name: "Computer 2"},
		&entity.AiBigTwoPlayer{Name: "Computer 3"},
		&entity.AiBigTwoPlayer{Name: "Computer 4"},
	}
}
