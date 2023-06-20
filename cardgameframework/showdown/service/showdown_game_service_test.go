package service

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"cardgameframework/showdown/entity"
)

func TestRunAGamePeacefully(t *testing.T) {
	p1 := entity.NewHumanPlayer(MockInput{}, MockOutput{})
	p2 := entity.NewHumanPlayer(MockInput{}, MockOutput{})
	p3 := entity.NewHumanPlayer(MockInput{}, MockOutput{})
	pAI := entity.NewAIPlayer(MockInput{}, MockOutput{})
	var deck *entity.Deck
	var game *ShowdownGame

	t.Run("Test creating game with human Player, AI Player, and new Deck", func(t *testing.T) {
		deck = entity.NewDeck()
		game = NewShowdownGame(p1, p2, p3, pAI, deck)

		assert.IsType(t, &ShowdownGame{}, game)
		assert.Equal(t, 4, len(game.Players))
		assert.Equal(t, 52, len(game.Deck.Cards))
	})

	t.Run("should successfully rename the human Player, except AI Player", func(t *testing.T) {
		p1.SetName("TestPlayer1")
		p2.SetName("TestPlayer2")
		p3.SetName("TestPlayer3")

		assert.Equal(t, "TestPlayer1", p1.Name())
		assert.Equal(t, "TestPlayer2", p2.Name())
		assert.Equal(t, "TestPlayer3", p3.Name())
	})

	t.Run("cards in a shuffled Deck should be random ordered", func(t *testing.T) {
		game.ShuffleDeck()
		c1 := game.Deck.Cards[0]
		game.ShuffleDeck()
		c2 := game.Deck.Cards[0]

		assert.NotEqual(t, c1, c2)
	})

	t.Run("when draw is finished, every Player should have 13 hand Card", func(t *testing.T) {
		game.DrawHands(13)

		assert.IsType(t, entity.Card{}, p1.HandCards[0])
		assert.Equal(t, rounds, len(p1.HandCards))
		assert.Equal(t, rounds, len(p2.HandCards))
		assert.Equal(t, rounds, len(p3.HandCards))
		assert.Equal(t, rounds, len(pAI.HandCards))
		assert.Equal(t, 52-rounds*4, len(game.Deck.Cards))
	})

	t.Run("Testing game over: game should be end after 13th rounds", func(t *testing.T) {
		game.TakeTurns()

		assert.Equal(t, 0, len(p1.HandCards))
		assert.Equal(t, 0, len(p2.HandCards))
		assert.Equal(t, 0, len(p3.HandCards))
		assert.Equal(t, 0, len(pAI.HandCards))
	})

	t.Run("Testing game result: winner's points should be the highest one", func(t *testing.T) {
		winner := game.GameResult()

		assert.NotEmpty(t, winner)
		for i := range game.Players {
			p := game.Players[i]
			if p != winner {
				//fmt.Printf("looser: %+v\n", p)
				assert.GreaterOrEqual(t, winner.Point(), p.Point())
			}
		}
	})
}

type MockInput struct{}

func (i MockInput) InputString() string {
	return "TestInputString"
}

func (i MockInput) InputNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return min + rand.Intn(max-min+1)
}

func (i MockInput) InputBool() bool {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) == 1
}

type MockOutput struct{}

func (m MockOutput) MeExchangeYourCardOutput() {}

func (m MockOutput) MeExchangeYourCardErrorOutput(err error) {}

func (m MockOutput) RenameOutput(name string) {}

func (m MockOutput) RoundStartOutput(i int) {}

func (m MockOutput) RoundResultOutput(i int, roundResults entity.RoundResult) {}

func (m MockOutput) GameOverOutput(winner entity.IPlayer, players []entity.IPlayer) {}

func (m MockOutput) YouExchangeMyCardOutput(name string) {}

func (m MockOutput) PrintCardsOutput(cards []entity.Card) {}

func (m MockOutput) AskToExchangeCardOutput(name string) {}

func (m MockOutput) ToExchangeCardOutput() {}

func (m MockOutput) TakeTurnStartOutput(name string) {}

func (m MockOutput) ExchangeBackOutput() {}

func (m MockOutput) AskShowCardOutput(name string) {}
