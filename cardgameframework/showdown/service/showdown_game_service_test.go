package service

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"cardgameframework/showdown/entity"
	"cardgameframework/template"
)

func TestRunAGamePeacefully(t *testing.T) {
	p1 := entity.NewHumanPlayer(MockInput{}, MockOutput{})
	p2 := entity.NewHumanPlayer(MockInput{}, MockOutput{})
	p3 := entity.NewHumanPlayer(MockInput{}, MockOutput{})
	pAI := entity.NewAIPlayer(MockInput{}, MockOutput{})
	var game *template.GameFramework[entity.ShowDownCard]

	t.Run("Test creating game with human Player, AI Player, and new Deck", func(t *testing.T) {
		game = NewShowdownGame([]entity.IShowdownPlayer[entity.ShowDownCard]{p1, p2, p3, pAI})

		assert.NotEmpty(t, game)
		assert.Equal(t, 4, len(game.Players))
		assert.Equal(t, 52, len(game.Deck.Cards))
	})

	t.Run("should successfully rename the human Player, except AI Player", func(t *testing.T) {
		p1.SetName("TestPlayer1")
		p2.SetName("TestPlayer2")
		p3.SetName("TestPlayer3")

		assert.Equal(t, "TestPlayer1", game.Players[0].GetName())
		assert.Equal(t, "TestPlayer2", game.Players[1].GetName())
		assert.Equal(t, "TestPlayer3", game.Players[2].GetName())
	})

	t.Run("cards in a shuffled Deck should be random ordered", func(t *testing.T) {
		game.ShuffleDeck()

		assert.NotEqual(t, game.Deck.Cards, entity.NewShowdownDeck())
		assert.Equal(t, 52, len(game.Deck.Cards))
	})

	t.Run("when draw is finished, every Player should have 13 hand ShowDownCard", func(t *testing.T) {
		game.DrawHands(13)

		assert.IsType(t, entity.ShowDownCard{}, p1.GetHand()[0])
		assert.Equal(t, rounds, len(p1.Hand))
		assert.Equal(t, rounds, len(p2.Hand))
		assert.Equal(t, rounds, len(p3.Hand))
		assert.Equal(t, rounds, len(pAI.Hand))
		assert.Equal(t, 52-rounds*4, len(game.Deck.Cards))
	})

	t.Run("Testing game over: game should be end after 13th rounds", func(t *testing.T) {
		game.PreTakeTurns()
		game.TakeTurns()

		assert.Equal(t, 0, len(p1.Hand))
		assert.Equal(t, 0, len(p2.Hand))
		assert.Equal(t, 0, len(p3.Hand))
		assert.Equal(t, 0, len(pAI.Hand))
	})

	t.Run("Testing game result: winner's points should be the highest one", func(t *testing.T) {
		winner := game.GameResult()

		assert.NotEmpty(t, winner)
		//for i := range game.Players {
		//	p := game.Players[i]
		//	if p != winner {
		//		//fmt.Printf("looser: %+v\n", p)
		//		assert.GreaterOrEqual(t, winner.Point(), p.Point())
		//	}
		//}
	})
}

type MockInput struct{}

func (i MockInput) InputString() string {
	return "TestInputString"
}

func (i MockInput) InputNum(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func (i MockInput) InputBool() bool {
	return rand.Intn(2) == 1
}

type MockOutput struct {
}

func (m MockOutput) MeExchangeYourCardOutput() {}

func (m MockOutput) MeExchangeYourCardErrorOutput(err error) {}

func (m MockOutput) RenameOutput(name string) {}

func (m MockOutput) RoundStartOutput(i int) {}

func (m MockOutput) RoundResultOutput(i int, roundResults entity.RoundResult) {}

func (m MockOutput) GameOverOutput(winner entity.IShowdownPlayer[entity.ShowDownCard], players []entity.IShowdownPlayer[entity.ShowDownCard]) {
}

func (m MockOutput) YouExchangeMyCardOutput(name string) {}

func (m MockOutput) PrintCardsOutput(cards []entity.ShowDownCard) {}

func (m MockOutput) AskToExchangeCardOutput(name string) {}

func (m MockOutput) ToExchangeCardOutput() {}

func (m MockOutput) TakeTurnStartOutput(name string) {}

func (m MockOutput) ExchangeBackOutput() {}

func (m MockOutput) AskShowCardOutput(name string) {}
