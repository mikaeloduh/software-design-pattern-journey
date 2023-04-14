package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"

	"showdown/entity"
)

func TestRunAGamePeacefully(t *testing.T) {
	p1 := entity.NewHumanPlayer(0, MockInput{})
	p2 := entity.NewHumanPlayer(1, MockInput{})
	p3 := entity.NewHumanPlayer(2, MockInput{})
	pAI := entity.NewAIPlayer(3, MockInput{})
	var deck *entity.Deck
	var game *Game

	t.Run("Test creating game with human player, AI player, and new Deck", func(t *testing.T) {
		deck = entity.NewDeck()
		game = NewGame(p1, p2, p3, pAI, deck)

		assert.IsType(t, &Game{}, game)
		assert.Equal(t, 4, len(game.Players))
		assert.Equal(t, 52, len(deck.Cards))
	})

	t.Run("should successfully rename the human player", func(t *testing.T) {
		p1.SetName("TestPlayer1")
		p2.SetName("TestPlayer2")
		p3.SetName("TestPlayer3")

		assert.Equal(t, "TestPlayer1", p1.Name())
		assert.Equal(t, "TestPlayer2", p2.Name())
		assert.Equal(t, "TestPlayer3", p3.Name())
		assert.Equal(t, "AI has no name", pAI.Name())
	})

	t.Run("cards in a shuffled Deck should be random ordered", func(t *testing.T) {
		game.Init()
		c1 := game.Deck.Cards[0]
		game.Init()
		c2 := game.Deck.Cards[0]

		assert.NotEqual(t, c1, c2)
	})

	t.Run("when draw is finished, every player should have 13 hand card", func(t *testing.T) {
		game.DrawLoop()

		assert.IsType(t, entity.Card{}, p1.HandCards[0])
		assert.Equal(t, rounds, len(p1.HandCards))
		assert.Equal(t, rounds, len(p2.HandCards))
		assert.Equal(t, rounds, len(p3.HandCards))
		assert.Equal(t, rounds, len(pAI.HandCards))
		assert.Equal(t, 52-rounds*4, len(game.Deck.Cards))
	})

	t.Run("Testing game over: game should be end after 13th rounds", func(t *testing.T) {
		game.takeTurnLoop()

		assert.Equal(t, 0, len(p1.HandCards))
		assert.Equal(t, 0, len(p2.HandCards))
		assert.Equal(t, 0, len(p3.HandCards))
		assert.Equal(t, 0, len(pAI.HandCards))
	})

	t.Run("Testing game result: winner's points should be the highest one", func(t *testing.T) {
		winner := game.gameResult()

		fmt.Printf("winner: %+v\n", winner)

		assert.NotEmpty(t, winner)
		for i := range game.Players {
			p := game.Players[i]
			if p != winner {
				fmt.Printf("looser: %+v\n", p)

				assert.GreaterOrEqual(t, winner.Point(), p.Point())
			}
		}
	})
}

func TestRunAGameBloodily(t *testing.T) {

	t.Run("testing exchange cards: two cards should be exchanged from a player to another", func(t *testing.T) {
		p1 := entity.NewHumanPlayer(0, MockInput{})
		p2 := entity.NewHumanPlayer(1, MockInput{})
		bigBlackTwo := entity.Card{Suit: entity.Spades, Rank: entity.Two}
		diamondThree := entity.Card{Suit: entity.Diamonds, Rank: entity.Three}
		p1.GetCard(bigBlackTwo)
		p2.GetCard(diamondThree)

		assert.Equal(t, bigBlackTwo, p1.HandCards[0])
		assert.Equal(t, diamondThree, p2.HandCards[0])

		_ = p1.MeExchangeYourCard(p2)

		assert.Equal(t, diamondThree, p1.HandCards[0])
		assert.Equal(t, bigBlackTwo, p2.HandCards[0])
	})

	t.Run("testing exchange cards: exchange card should not proceed if player has run out of hand cards", func(t *testing.T) {

		p3 := entity.NewHumanPlayer(2, MockInput{})
		p4 := entity.NewHumanPlayer(3, MockInput{})
		bigBlackTwo := entity.Card{Suit: entity.Spades, Rank: entity.Two}
		p3.GetCard(bigBlackTwo)

		assert.Equal(t, bigBlackTwo, p3.HandCards[0])
		assert.Empty(t, p4.HandCards)

		_ = p3.MeExchangeYourCard(p4)

		assert.Equal(t, bigBlackTwo, p3.HandCards[0])
		assert.Empty(t, p4.HandCards)
	})

	t.Run("test takeTurnLoop with exchange step and have no problem", func(t *testing.T) {
		p1 := entity.NewHumanPlayer(0, MockInput{})
		p2 := entity.NewHumanPlayer(1, MockInput{})
		p3 := entity.NewHumanPlayer(2, MockInput{})
		p4 := entity.NewHumanPlayer(3, MockInput{})
		deck := entity.NewDeck()
		game := NewGame(p1, p2, p3, p4, deck)

		game.Init()
		game.DrawLoop()
		game.takeTurnLoop()

		assert.Equal(t, 0, len(p1.HandCards))
		assert.Equal(t, 0, len(p2.HandCards))
		assert.Equal(t, 0, len(p3.HandCards))
		assert.Equal(t, 0, len(p3.HandCards))

		winner := game.gameResult()

		fmt.Printf("winner: %+v\n", winner)

		assert.NotEmpty(t, winner)
		for i := range game.Players {
			p := game.Players[i]
			if p != winner {
				fmt.Printf("looser: %+v\n", p)

				assert.GreaterOrEqual(t, winner.Point(), p.Point())
			}
		}
	})
}

type MockInput struct{}

func (i MockInput) InputNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return min + rand.Intn(max-min+1)
}

func (i MockInput) InputBool() bool {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) == 1
}
