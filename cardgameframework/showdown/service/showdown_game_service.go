package service

import (
	"cardgameframework/showdown/entity"
	"fmt"
)

type ShowdownGame struct {
	Players       []entity.IPlayer
	Deck          *entity.Deck
	CurrentPlayer int
	Record        entity.Record
}

const rounds int = 13

func NewShowdownGame(p1 entity.IPlayer, p2 entity.IPlayer, p3 entity.IPlayer, p4 entity.IPlayer, deck *entity.Deck) *ShowdownGame {
	for i, p := range []entity.IPlayer{p1, p2, p3, p4} {
		p.SetId(i)
		p.SetName(fmt.Sprintf("P%d", i))
	}

	return &ShowdownGame{
		Players:       []entity.IPlayer{p1, p2, p3, p4},
		Deck:          deck,
		CurrentPlayer: 0,
		Record:        entity.Record{nil},
	}
}

func (g *ShowdownGame) PreTakeTurns() {
	fmt.Printf("Game Start")
}

func (g *ShowdownGame) TakeTurnStep(player entity.IPlayer) {
	g.Record[len(g.Record)-1] = append(g.Record[len(g.Record)-1], entity.Result{
		Player: player,
		Card:   player.TakeTurn(g.Players),
		Win:    false,
	})
}

func (g *ShowdownGame) GetCurrentPlayer() entity.IPlayer {
	return g.Players[g.CurrentPlayer]
}

func (g *ShowdownGame) UpdateGameAndMoveToNext() {
	// move to next player
	g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)

	// if hit next round
	if g.CurrentPlayer == 0 {
		currentRecord := g.Record[len(g.Record)-1]

		greatest := entity.Card{Suit: entity.Clubs, Rank: entity.Three}
		greatestIdx := 0
		for i, r := range currentRecord {
			if r.Card.IsGreater(greatest) {
				greatest = r.Card
				greatestIdx = i
			}
		}
		currentRecord[greatestIdx].Win = true
		currentRecord[greatestIdx].Player.AddPoint()

		g.Players[0].RoundResultOutput(0, currentRecord)

		g.Record = append(g.Record, nil)
	}
}

func (g *ShowdownGame) IsGameFinished() bool {
	// if all players ran out their hands
	var num int
	for _, player := range g.Players {
		num += len(player.GetHand())
	}
	return num == 0
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
