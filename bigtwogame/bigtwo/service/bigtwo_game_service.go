package service

import (
	"bigtwogame/bigtwo/entity"
	"bigtwogame/template"
)

type BigTwoGame struct {
	Players       []entity.IBigTwoPlayer
	Deck          template.Deck[entity.BigTwoCard]
	DeskCard      entity.BigTwoCard
	CurrentPlayer int
}

func NewBigTwoGame(players []entity.IBigTwoPlayer) *template.GameFramework[entity.BigTwoCard] {
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

	for i, p := range b.Players {
		if b.haveValidCards(p.GetHand()) {
			var playerPlay func() entity.BigTwoCard
			playerPlay = func() entity.BigTwoCard {
				move := p.TakeTurnMove()
				if !b.isValidMove(move.DryRun()) {
					return playerPlay()
				}
				return move.Play()
			}

			b.CurrentPlayer = i
			playCard := playerPlay()
			b.updateDeskCard(playCard)
		}
	}
}

func (b *BigTwoGame) haveValidCards(hand []entity.BigTwoCard) bool {
	for _, card := range hand {
		if b.isValidMove(card) {
			return true
		}
	}
	return false
}

func (b *BigTwoGame) isValidMove(card entity.BigTwoCard) bool {
	return (card.Rank == entity.Three) && (card.Suit == entity.Clubs)
}

func (b *BigTwoGame) TakeTurnStep(player template.IPlayer[entity.BigTwoCard]) {
	//TODO implement me
	panic("implement me")
}

func (b *BigTwoGame) GetCurrentPlayer() template.IPlayer[entity.BigTwoCard] {
	return b.Players[b.CurrentPlayer]
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

func (b *BigTwoGame) updateDeskCard(card entity.BigTwoCard) {
	b.DeskCard = card
}
