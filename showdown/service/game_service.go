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
		g.Players[i].ReName()
	}
}

func (g *Game) drawLoop() {
	for i := 0; i < rounds; i++ {
		for i := range g.Players {
			g.Players[i].GetCard(g.Deck.DrawCard())
		}
	}
}

type RoundResult struct {
	player entity.IPlayer
	card   entity.Card
}

type RoundResults []RoundResult

func (g *Game) takeTurnLoop() {

	for i := 0; i < rounds; i++ {
		fmt.Printf("\n============== Round %d ==============\n", i)

		roundResults := make(RoundResults, len(g.Players))
		for i := range roundResults {
			roundResults[i] = RoundResult{g.Players[i], g.Players[i].TakeTurn(g.Players)}
		}

		win := g.showDown(roundResults)

		// Printing round win
		fmt.Printf("\n* Round %d end\n", i)
		for _, rr := range roundResults {
			fmt.Printf("[%4s ]   ", rr.card.String())
		}
		fmt.Print("\n")
		for _, rr := range roundResults {
			fmt.Printf("Player %d  ", rr.player.Id())
		}
		fmt.Printf("\n Player %d: %s win!\n", win.player.Id(), win.player.Name())

		win.player.AddPoint()
	}
}

func (g *Game) showDown(rrs RoundResults) RoundResult {
	if len(rrs) == 0 {
		return RoundResult{}
	}

	greatest := rrs[0]
	for _, rr := range rrs {
		if rr.card.IsGreater(greatest.card) {
			greatest = rr
		}
	}

	return greatest
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

	fmt.Printf("\n============== Game Over ==============\nThe Winner is P%d: %s\n", winner.Id(), winner.Name())

	for _, p := range g.Players {
		fmt.Printf("P%d:%d point\n", p.Id(), p.Point())
	}

	return winner
}

func NewGame(p1 entity.IPlayer, p2 entity.IPlayer, p3 entity.IPlayer, p4 entity.IPlayer, deck *entity.Deck) *Game {
	return &Game{
		Players: []entity.IPlayer{p1, p2, p3, p4},
		Deck:    deck,
	}
}
