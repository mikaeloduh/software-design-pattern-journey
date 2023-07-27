package service

import (
	"bigtwogame/showdown/entity"
	"bigtwogame/template"
	"fmt"
)

type ShowdownGame struct {
	Players       []entity.IShowdownPlayer
	Deck          *template.Deck[entity.ShowDownCard]
	CurrentPlayer int
	Record        entity.Record
}

const rounds int = 13

func NewShowdownGame(players []entity.IShowdownPlayer) *template.GameFramework[entity.ShowDownCard] {
	deck := entity.NewShowdownDeck()
	game := &template.GameFramework[entity.ShowDownCard]{
		Deck:    deck,
		Players: make([]template.IPlayer[entity.ShowDownCard], len(players)),
		NumCard: 13,
		PlayingGame: &ShowdownGame{
			Players:       players,
			Deck:          deck,
			CurrentPlayer: 0,
			Record:        entity.Record{nil},
		},
	}
	for i, player := range players {
		player.SetId(i)
		game.Players[i] = player
	}

	return game
}

func (g *ShowdownGame) Init() {}

func (g *ShowdownGame) PreTakeTurns() {
	fmt.Printf("Game Start")
}

func (g *ShowdownGame) TakeTurnStep(player template.IPlayer[entity.ShowDownCard]) {
	g.Record[len(g.Record)-1] = append(g.Record[len(g.Record)-1], entity.Result{
		Player: player,
		Card:   player.(entity.IShowdownPlayer).TakeTurn(),
		Win:    false,
	})
}

func (g *ShowdownGame) GetCurrentPlayer() template.IPlayer[entity.ShowDownCard] {
	return g.Players[g.CurrentPlayer]
}

func (g *ShowdownGame) UpdateGameAndMoveToNext() {
	// move to next player
	g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)

	// if hit next round
	if g.CurrentPlayer == 0 {
		currentRecord := g.Record[len(g.Record)-1]

		greatest := entity.ShowDownCard{Suit: entity.Clubs, Rank: entity.Three}
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

func (g *ShowdownGame) GameResult() template.IPlayer[entity.ShowDownCard] {
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
