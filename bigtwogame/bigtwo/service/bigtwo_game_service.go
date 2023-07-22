package service

import (
	"bigtwogame/bigtwo/entity"
	"bigtwogame/template"
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
		if b.haveValidPreTakeMove(p.GetHand()) {
			var playerTakeTurn func() []entity.BigTwoCard

			playerTakeTurn = func() []entity.BigTwoCard {
				move := p.TakeTurnMove()
				if !b.isValidPreTakeMove(move.DryRun()) {
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

	if b.haveValidTurnMove(player.GetHand()) {
		var playerPlay func() []entity.BigTwoCard

		playerPlay = func() []entity.BigTwoCard {
			move := player.(entity.IBigTwoPlayer).TakeTurnMove()
			if !b.isValidTurnMove(move.DryRun()) {
				return playerPlay()
			}
			return move.Play()
		}

		b.updateDeskCard(playerPlay())
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

// privates

func (b *BigTwoGame) haveValidPreTakeMove(hand []entity.BigTwoCard) bool {
	for _, card := range hand {
		if b.isValidPreTakeMove([]entity.BigTwoCard{card}) {
			return true
		}
	}
	return false
}

func (b *BigTwoGame) isValidPreTakeMove(cards []entity.BigTwoCard) bool {
	if !hasClubsThree(cards) {
		return false
	}
	if isMatchSingle(cards) {
		// IF playCard match single patten
		return cards[0].Compare(entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Three}) == 0
	}
	return false
}

func (b *BigTwoGame) haveValidTurnMove(hand []entity.BigTwoCard) bool {
	for _, card := range hand {
		if b.isValidTurnMove([]entity.BigTwoCard{card}) {
			return true
		}
	}
	return false
}

func (b *BigTwoGame) isValidTurnMove(cards []entity.BigTwoCard) bool {
	if len(b.TopCards) == 1 && b.TopCards[0] == entity.PassCard() {
		// IF IS PassCard : all passed
		if isMatchSingle(cards) {
			// IF playCard match single patten
			return true
		}
		return false
	} else {
		if isMatchSingle(b.TopCards) && isMatchSingle(cards) {
			// IF IS single card
			// IF playCard match single patten
			return cards[0].Compare(b.TopCards[0]) == 1
		}
		return false
	}
}

func (b *BigTwoGame) updateDeskCard(card []entity.BigTwoCard) {
	b.TopCards = card
}

// helpers

func hasClubsThree(cards []entity.BigTwoCard) bool {
	for _, v := range cards {
		if v == (entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Three}) {
			return true
		}
	}
	return false
}

func isMatchSingle(cards []entity.BigTwoCard) bool {
	return len(cards) == 1
}
