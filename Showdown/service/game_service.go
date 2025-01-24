package service

import (
	"fmt"
	"showdown/entity"
)

type Game struct {
	Players []entity.IPlayer
	Deck    *entity.Deck
}

const rounds int = 13

func (g *Game) Run() {
	g.init()
	g.renamePlayer()
	g.drawLoop()
	g.takeTurnLoop()
	g.gameResult()
}

func (g *Game) init() {
	g.Deck.Shuffle()
}

func (g *Game) renamePlayer() {
	for i := range g.Players {
		g.Players[i].Rename()
	}
}

func (g *Game) drawLoop() {
	for i := 0; i < rounds; i++ {
		for i := range g.Players {
			g.Players[i].GetCard(g.Deck.DrawCard())
		}
	}
}

func (g *Game) takeTurnLoop() {

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

func (g *Game) gameResult() entity.IPlayer {
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

func NewGame(p1 entity.IPlayer, p2 entity.IPlayer, p3 entity.IPlayer, p4 entity.IPlayer, deck *entity.Deck) *Game {
	for i, p := range []entity.IPlayer{p1, p2, p3, p4} {
		p.SetId(i)
		p.SetName(fmt.Sprintf("P%d", i))
	}

	return &Game{
		Players: []entity.IPlayer{p1, p2, p3, p4},
		Deck:    deck,
	}
}
