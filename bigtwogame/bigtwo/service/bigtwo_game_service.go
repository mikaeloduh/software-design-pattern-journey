package service

import (
	"bigtwogame/bigtwo/entity"
	"bigtwogame/template"
)

type BigTwoGame struct {
	Players  []entity.IBigTwoPlayer[entity.BigTwoCard]
	Deck     template.Deck[entity.BigTwoCard]
	DeskCard entity.BigTwoCard
}

func NewBigTwoGame(players []entity.IBigTwoPlayer[entity.BigTwoCard]) *template.GameFramework[entity.BigTwoCard] {
	deck := entity.NewBigTwoDeck()
	game := &template.GameFramework[entity.BigTwoCard]{
		Deck:        deck,
		Players:     make([]template.IPlayer[entity.BigTwoCard], len(players)),
		NumCard:     13,
		PlayingGame: &BigTwoGame{Players: players, Deck: deck},
	}
	for i, player := range players {
		game.Players[i] = player
	}

	return game
}

func (b *BigTwoGame) PreTakeTurns() {
	//TODO implement me
	b.DeskCard = entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Three}
}

func (b *BigTwoGame) TakeTurnStep(player template.IPlayer[entity.BigTwoCard]) {
	//TODO implement me
	panic("implement me")
}

func (b *BigTwoGame) GetCurrentPlayer() template.IPlayer[entity.BigTwoCard] {
	//TODO implement me
	panic("implement me")
}

func (b *BigTwoGame) UpdateGameAndMoveToNext() {
	//TODO implement me
	panic("implement me")
}

func (b *BigTwoGame) IsGameFinished() bool {
	//TODO implement me
	panic("implement me")
}

func (b *BigTwoGame) GameResult() template.IPlayer[entity.BigTwoCard] {
	//TODO implement me
	panic("implement me")
}
