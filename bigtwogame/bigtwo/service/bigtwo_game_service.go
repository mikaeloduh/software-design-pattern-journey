package service

import (
	"bigtwogame/bigtwo/entity"
	"bigtwogame/template"
	"fmt"
)

type BigTwoGame struct {
	Players       []entity.IBigTwoPlayer
	Deck          *template.Deck[entity.BigTwoCard]
	TopCards      []entity.BigTwoCard
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
	b.TopCards = []entity.BigTwoCard{entity.InitCard()}

	for i, p := range b.Players {
		if b.haveValidMove(p.GetHand()) {
			var playerTakeTurn func() []entity.BigTwoCard
			playerTakeTurn = func() []entity.BigTwoCard {
				move := p.TakeTurnMove()
				if !b.isValidMove(move.DryRun()) {
					return playerTakeTurn()
				}
				return move.Play()
			}

			b.CurrentPlayer = i
			b.updateDeskCard(playerTakeTurn())
			break
		}
	}
	b.UpdateGameAndMoveToNext()
}

func (b *BigTwoGame) TakeTurnStep(player template.IPlayer[entity.BigTwoCard]) {

	if b.haveValidMove(player.GetHand()) {
		var playerPlay func() []entity.BigTwoCard
		playerPlay = func() []entity.BigTwoCard {
			move := player.(entity.IBigTwoPlayer).TakeTurnMove()
			if !b.isValidMove(move.DryRun()) {
				return playerPlay()
			}
			return move.Play()
		}

		playCard := playerPlay()
		b.updateDeskCard(playCard)
		b.Passed = 0
	} else {
		b.Passed += 1
	}

}

func (b *BigTwoGame) GetCurrentPlayer() template.IPlayer[entity.BigTwoCard] {
	return b.Players[b.CurrentPlayer]
}

func (b *BigTwoGame) UpdateGameAndMoveToNext() {
	b.CurrentPlayer = (b.CurrentPlayer + 1) % len(b.Players)

	if b.Passed == len(b.Players)-1 {
		b.TopCards = []entity.BigTwoCard{entity.PassCard()}
		b.Passed = 0
	}
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

func (b *BigTwoGame) haveValidMove(hand []entity.BigTwoCard) bool {
	for _, card := range hand {
		if b.isValidMove([]entity.BigTwoCard{card}) {
			return true
		}
	}
	return false
	//return b.isValidMove(hand)
}

func (b *BigTwoGame) isValidMove(cards []entity.BigTwoCard) bool {
	if len(b.TopCards) == 1 && b.TopCards[0] == entity.InitCard() {
		// IF pre take turn
		card := cards[0]
		return card.Compare(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Three}) == 0
	} else if len(b.TopCards) == 1 && b.TopCards[0] == entity.PassCard() {
		// IF all passed
		return true
	} else if len(b.TopCards) == 1 && len(cards) == 1 {
		// IF play single card
		card := cards[0]
		return card.Compare(b.TopCards[0]) == 1
	} else {
		fmt.Println("You should not pass!")
		return false
	}
}

func (b *BigTwoGame) updateDeskCard(card []entity.BigTwoCard) {
	b.TopCards = card
}
