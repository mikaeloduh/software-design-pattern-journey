package service

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"bigtwogame/bigtwo/entity"
	"bigtwogame/template"
)

func TestBigTwo(t *testing.T) {
	t.Parallel()

	t.Run("New game success and have 4 players", func(t *testing.T) {
		players := FakeNewAIPlayers()
		deck := entity.NewBigTwoDeck()
		game := NewBigTwoGame(deck, players)

		assert.Equal(t, 4, len(game.Players))
	})

	t.Run("New a Deck and have it shuffled", func(t *testing.T) {
		deck := entity.NewBigTwoDeck()

		assert.Equal(t, 52, len(deck.Cards))

		deck.Shuffle()

		assert.NotEqual(t, entity.NewBigTwoDeck(), deck)
	})

	t.Run("New game and have card deal to all players", func(t *testing.T) {
		players := FakeNewAIPlayers()
		deck := entity.NewBigTwoDeck()
		game := NewBigTwoGame(deck, players)

		game.DrawHands(game.NumCard)

		assert.Equal(t, 13, len(game.Players[0].GetHand()))
		assert.Equal(t, 13, len(game.Players[1].GetHand()))
		assert.Equal(t, 13, len(game.Players[2].GetHand()))
		assert.Equal(t, 13, len(game.Players[3].GetHand()))
		assert.Equal(t, 0, len(game.Deck.Cards))
	})

	t.Run("PreTakeTurn should play ♣3 from whoever had (single only)", func(t *testing.T) {
		players := FakeNewAIPlayers()
		deck := entity.NewBigTwoDeck()
		game, playingGame := FakeNewBigTwoGame(deck, players)

		playingGame.SetActionCards()
		game.ShuffleDeck()
		game.DrawHands(game.NumCard)
		game.PreTakeTurns()

		assert.Contains(t, playingGame.TopCards, entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Three})
	})

	t.Run("Testing TakeTurn while player not pass, played card should be a valid single", func(t *testing.T) {
		players := FakeNewAIPlayers()
		deck := entity.NewBigTwoDeck()
		_, playingGame := FakeNewBigTwoGame(deck, players)
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
		players := FakeNewAIPlayers()
		deck := entity.NewBigTwoDeck()
		_, playingGame := FakeNewBigTwoGame(deck, players)
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
		players := FakeNewAIPlayers()
		deck := entity.NewBigTwoDeck()
		game, playingGame := FakeNewBigTwoGame(deck, players)

		game.ShuffleDeck()
		game.DrawHands(game.NumCard)
		oldTopCards := []entity.BigTwoCard{
			{Suit: entity.Clubs, Rank: entity.Three},
			{Suit: entity.Diamonds, Rank: entity.Three}}
		playingGame.TopCards = oldTopCards

		playingGame.CurrentPlayer = 3
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())

		assert.Equal(t, playingGame.TopCards[0].Rank, playingGame.TopCards[1].Rank)
		assert.Greater(t, playingGame.TopCards[0].Rank, oldTopCards[0].Rank)
		assert.Len(t, playingGame.GetCurrentPlayer().GetHand(), 13-2)
	})

	t.Run("Testing TakeTurn with no valid card player should pass (single only)", func(t *testing.T) {
		players := FakeNewAIPlayers()
		deck := entity.NewBigTwoDeck()
		_, playingGame := FakeNewBigTwoGame(deck, players)
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
		players := FakeNewAIPlayers()
		deck := entity.NewBigTwoDeck()
		_, playingGame := FakeNewBigTwoGame(deck, players)

		playingGame.SetActionCards()
		// TopCards is Club 10
		playingGame.TopCards = []entity.BigTwoCard{{Suit: entity.Clubs, Rank: entity.Ten}}
		players[0].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Nine})
		players[1].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Eight})
		players[2].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Seven})
		players[3].SetCard(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Two})

		// Player 0 should pass
		playingGame.CurrentPlayer = 0
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())
		playingGame.UpdateGameAndMoveToNext()
		// Player 1 should pass
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())
		playingGame.UpdateGameAndMoveToNext()
		// Player 2 should pass
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())
		playingGame.UpdateGameAndMoveToNext()
		// Player 3 should play hand
		playingGame.TakeTurnStep(playingGame.GetCurrentPlayer())
		playingGame.UpdateGameAndMoveToNext()

		assert.Equal(t, []entity.BigTwoCard{{Suit: entity.Clubs, Rank: entity.Two}}, playingGame.TopCards)
		assert.Len(t, players[0].GetHand(), 1)
		assert.Len(t, players[1].GetHand(), 1)
		assert.Len(t, players[2].GetHand(), 1)
		assert.Len(t, players[3].GetHand(), 0)
	})
}

func TestBigTwoAcceptanceTest(t *testing.T) {
	t.Skip()

	inputData, _ := os.Open("./test_data/always-play-first-card.in")
	defer inputData.Close()
	inputDataReader := bufio.NewScanner(inputData)
	scanInputDataText := func() string {
		inputDataReader.Scan()
		return inputDataReader.Text()
	}

	outputData, _ := os.Open("./test_data/always-play-first-card.out")
	defer outputData.Close()
	outputDataReader := bufio.NewScanner(inputData)
	scanOutputDataText := func() string {
		outputDataReader.Scan()
		return outputDataReader.Text()
	}

	// Create a temporary buffer to capture the output
	var writer *bytes.Buffer
	writerWrite := func() string {
		return writer.String()
	}
	// Set up custom input data using a bytes.Buffer
	var reader *bytes.Buffer
	readerRead := func(s string) {
		reader = bytes.NewBufferString(s)
	}

	deck := entity.NewBigTwoDeck()
	game, _ := FakeNewBigTwoGame(deck, FakeNewHumanPlayers(reader, writer))

	// Set up deck
	scanInputDataText()
	//playingGame.Deck = NewMockDeck(fileReader.Text())

	// Player rename
	game.Init()
	readerRead(scanInputDataText())
	readerRead(scanInputDataText())
	readerRead(scanInputDataText())
	readerRead(scanInputDataText())

	// Draw
	game.DrawHands(game.NumCard)
	expectedOut := scanOutputDataText()
	actualOut := writerWrite()

	assert.Equal(t, expectedOut, actualOut)

	// PreTakeTurn
	game.PreTakeTurns()
}

func FakeNewBigTwoGame(deck *entity.BigTwoDeck, players []entity.IBigTwoPlayer) (*template.GameFramework[entity.BigTwoCard], *BigTwoGame) {
	playingGame := &BigTwoGame{Players: players, Deck: deck}
	game := &template.GameFramework[entity.BigTwoCard]{
		Deck:        &deck.Deck,
		Players:     make([]template.IPlayer[entity.BigTwoCard], len(players)),
		NumCard:     13,
		PlayingGame: playingGame,
	}
	for i, player := range players {
		game.Players[i] = player
	}

	return game, playingGame
}

func FakeNewAIPlayers() []entity.IBigTwoPlayer {
	return []entity.IBigTwoPlayer{
		&entity.AiBigTwoPlayer{Name: "Computer 1"},
		&entity.AiBigTwoPlayer{Name: "Computer 2"},
		&entity.AiBigTwoPlayer{Name: "Computer 3"},
		&entity.AiBigTwoPlayer{Name: "Computer 4"},
	}
}

func FakeNewHumanPlayers(r io.Reader, w io.Writer) []entity.IBigTwoPlayer {
	return []entity.IBigTwoPlayer{
		&entity.HumanPlayer{Name: "Player 1", Reader: r, Writer: w},
		&entity.HumanPlayer{Name: "Player 2", Reader: r, Writer: w},
		&entity.HumanPlayer{Name: "Player 3", Reader: r, Writer: w},
		&entity.HumanPlayer{Name: "Player 4", Reader: r, Writer: w},
	}
}
