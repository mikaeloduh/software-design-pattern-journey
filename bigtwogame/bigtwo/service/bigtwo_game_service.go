package service

import (
	"bigtwogame/bigtwo/entity"
	"bigtwogame/template"
)

type BigTwoGame struct {
	Players       []entity.IBigTwoPlayer
	Deck          *template.Deck[entity.BigTwoCard]
	DeskCard      entity.BigTwoCard
	CurrentPlayer int
	Passed        int
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
		if b.haveValidTop(p.GetHand()) {
			var playerPlay func() entity.BigTwoCard
			playerPlay = func() entity.BigTwoCard {
				move := p.TakeTurnMove()
				if !b.isValidTop(move.DryRun()) {
					return playerPlay()
				}
				return move.Play()
			}

			b.CurrentPlayer = i
			playCard := playerPlay()
			b.updateDeskCard(playCard)
			break
		}
	}
	b.UpdateGameAndMoveToNext()
}

func (b *BigTwoGame) TakeTurnStep(player template.IPlayer[entity.BigTwoCard]) {

	if b.haveValidMove(player.GetHand()) {
		var playerPlay func() entity.BigTwoCard
		playerPlay = func() entity.BigTwoCard {
			move := player.(entity.IBigTwoPlayer).TakeTurnMove()
			if !b.isValidMove(move.DryRun()) {
				return playerPlay()
			}
			return move.Play()
		}

		playCard := playerPlay()
		b.updateDeskCard(playCard)
	} else {
		b.Passed += 1
	}

}

func (b *BigTwoGame) GetCurrentPlayer() template.IPlayer[entity.BigTwoCard] {
	return b.Players[b.CurrentPlayer]
}

func (b *BigTwoGame) UpdateGameAndMoveToNext() {
	b.CurrentPlayer = (b.CurrentPlayer + 1) % len(b.Players)
}

func (b *BigTwoGame) IsGameFinished() bool {
	//TODO implement me
	panic("implement me")
}

func (b *BigTwoGame) GameResult() template.IPlayer[entity.BigTwoCard] {
	//TODO implement me
	panic("implement me")
}

// Helpers

func (b *BigTwoGame) haveValidTop(hand []entity.BigTwoCard) bool {
	for _, card := range hand {
		if b.isValidTop(card) {
			return true
		}
	}
	return false
}

func (b *BigTwoGame) isValidTop(card entity.BigTwoCard) bool {
	return card.Compare(entity.BigTwoCard{
		Suit: entity.Clubs,
		Rank: entity.Three,
	}) == 0
}

func (b *BigTwoGame) haveValidMove(hand []entity.BigTwoCard) bool {
	for _, card := range hand {
		if b.isValidMove(card) {
			return true
		}
	}
	return false
}

func (b *BigTwoGame) isValidMove(card entity.BigTwoCard) bool {
	return card.Compare(b.DeskCard) == 1
}

func (b *BigTwoGame) updateDeskCard(card entity.BigTwoCard) {
	b.DeskCard = card
}
