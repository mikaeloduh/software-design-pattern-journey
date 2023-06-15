package service

import (
	"cardgameframework/entity"
	"fmt"
)

type ShowdownGame struct {
	Players []entity.IShowdownPlayer
	Deck    *entity.ShowdownDeck
}

const rounds int = 13

func NewShowdownGame(p1 entity.IShowdownPlayer, p2 entity.IShowdownPlayer, p3 entity.IShowdownPlayer, p4 entity.IShowdownPlayer, deck *entity.ShowdownDeck) *ShowdownGame {
	for i, p := range []entity.IShowdownPlayer{p1, p2, p3, p4} {
		p.SetId(i)
		p.SetName(fmt.Sprintf("P%d", i))
	}

	return &ShowdownGame{
		Players: []entity.IShowdownPlayer{p1, p2, p3, p4},
		Deck:    deck,
	}
}

func (g *ShowdownGame) Run() {
	g.Init()
	g.Draw()
	g.TakeTurn()
	g.GameResult()
}

func (g *ShowdownGame) Init() {
	for i := range g.Players {
		g.Players[i].Rename()
	}

	g.Deck.Shuffle()
}

func (g *ShowdownGame) Draw() {
	for i := 0; i < rounds; i++ {
		for i := range g.Players {
			g.Players[i].AssignCard(g.Deck.DrawCard())
		}
	}
}

func (g *ShowdownGame) TakeTurn() {

	for i := 0; i < rounds; i++ {
		g.Players[0].RoundStartOutput(i)

		roundResults := make(entity.RoundResults, len(g.Players))
		for r := range roundResults {
			roundResults[r] = entity.ShowdownRoundResult{
				Player: g.Players[r],
				Card:   g.Players[r].TakeTurn(g.Players),
				Win:    false,
			}
		}

		greatest := entity.ShowdownCard{Suit: entity.Clubs, Rank: entity.Three}
		for _, rr := range roundResults {
			if rr.Card.IsGreater(greatest) {
				greatest = rr.Card
			}
		}
		for j, rr := range roundResults {
			if rr.Card == greatest {
				roundResults[j].Win = true
				roundResults[j].Player.AddPoint()
			}
		}

		g.Players[0].RoundResultOutput(i, roundResults)
	}
}

func (g *ShowdownGame) GameResult() entity.IShowdownPlayer {
	var winner entity.IShowdownPlayer
	max := 0
	for i := range g.Players {
		if g.Players[i].Point() > max {
			max = g.Players[i].Point()
			winner = g.Players[i]
		}
	}

	g.Players[0].GameOverOutput(winner, g.Players)

	return winner
}
