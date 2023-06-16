package service

import (
	"cardgameframework/showdown/entity"
	"fmt"
)

type ShowdownGame struct {
	Players []entity.IPlayer
	Deck    *entity.Deck
}

const rounds int = 13

func NewShowdownGame(p1 entity.IPlayer, p2 entity.IPlayer, p3 entity.IPlayer, p4 entity.IPlayer, deck *entity.Deck) *ShowdownGame {
	for i, p := range []entity.IPlayer{p1, p2, p3, p4} {
		p.SetId(i)
		p.SetName(fmt.Sprintf("P%d", i))
	}

	return &ShowdownGame{
		Players: []entity.IPlayer{p1, p2, p3, p4},
		Deck:    deck,
	}
}

func (g *ShowdownGame) Run() {
	g.Init()
	g.ShuffleDeck()
	g.DrawHands()
	g.TakeTurns()
	g.GameResult()
}

func (g *ShowdownGame) Init() {
	for i := range g.Players {
		g.Players[i].Rename()
	}
}

func (g *ShowdownGame) ShuffleDeck() {
	g.Deck.Shuffle()
}

func (g *ShowdownGame) DrawHands() {
	for i := 0; i < rounds; i++ {
		for i := range g.Players {
			g.Players[i].AssignCard(g.Deck.DrawCard())
		}
	}
}

func (g *ShowdownGame) TakeTurns() {

	for i := 0; i < rounds; i++ {
		g.Players[0].RoundStartOutput(i)

		roundResults := make(entity.RoundResults, len(g.Players))
		for r := range roundResults {
			roundResults[r] = entity.RoundResult{
				Player: g.Players[r],
				Card:   g.Players[r].TakeTurn(g.Players),
				Win:    false,
			}
		}

		greatest := entity.Card{Suit: entity.Clubs, Rank: entity.Three}
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

func (g *ShowdownGame) GameResult() entity.IPlayer {
	var winner entity.IPlayer
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
