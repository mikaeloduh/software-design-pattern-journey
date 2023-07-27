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

func (b *BigTwoGame) Init() {
	b.SetActionCards()
}

func (b *BigTwoGame) SetActionCards() {
	for _, v := range b.Players {
		v.SetActionCard(entity.PassCard())
	}
}

func (b *BigTwoGame) PreTakeTurns() {
	b.TopCards = []entity.BigTwoCard{entity.InitCard()}

	for i, p := range b.Players {
		if b.hasValidPreTakeTurnMove(p.GetHand()) {
			var playerTakeTurn func() []entity.BigTwoCard

			playerTakeTurn = func() []entity.BigTwoCard {
				move := p.TakeTurnMove()
				if !b.isValidTurnMove(move.DryRun()) {
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
	var playerPlay func() []entity.BigTwoCard

	playerPlay = func() []entity.BigTwoCard {
		move := player.(entity.IBigTwoPlayer).TakeTurnMove()
		if !b.isValidTurnMove(move.DryRun()) {
			return playerPlay()
		}
		return move.Play()
	}

	b.updateDeskCard(playerPlay())
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

func (b *BigTwoGame) hasValidPreTakeTurnMove(hand []entity.BigTwoCard) bool {
	for _, card := range hand {
		if b.isValidTurnMove([]entity.BigTwoCard{card}) {
			return true
		}
	}
	return false
}

// isValidTurnMove
func (b *BigTwoGame) isValidTurnMove(cards []entity.BigTwoCard) bool {
	if isPassCard(b.TopCards) {
		if isPassCard(cards) {
			return false
		} else if isMatchSingle(cards) {
			return true
		} else if isMatchPair(cards) {
			return true
		}
		return false
	} else if isInitCard(b.TopCards) {
		if !hasClubsThree(cards) {
			return false
		}
		if isMatchSingle(cards) {
			return true
		} else if isMatchPair(cards) {
			return true
		}
		return false
	} else if isPassCard(cards) {
		return true
	} else if isMatchSingle(b.TopCards) {
		return compareSingle(cards, b.TopCards)
	} else if isMatchPair(b.TopCards) {
		return ComparePair(cards, b.TopCards)
	}
	return false
}

func isInitCard(cards []entity.BigTwoCard) bool {
	return len(cards) == 1 && cards[0] == entity.InitCard()
}

func isPassCard(cards []entity.BigTwoCard) bool {
	return len(cards) == 1 && cards[0] == entity.PassCard()
}

func compareSingle(sub, tar []entity.BigTwoCard) bool {
	if !isMatchSingle(sub) || !isMatchSingle(tar) {
		return false
	}
	return sub[0].Compare(tar[0]) == 1
}

func ComparePair(subject, target []entity.BigTwoCard) bool {
	// subject greater than target -> true
	if !isMatchPair(subject) || !isMatchPair(target) {
		return false
	}
	if subject[0].Compare(target[0]) == 1 || subject[0].Compare(target[1]) == 1 || subject[1].Compare(target[0]) == 1 || subject[1].Compare(target[1]) == 1 {
		return true
	}
	return false
}

func (b *BigTwoGame) updateDeskCard(cards []entity.BigTwoCard) {
	if cards[0] == entity.PassCard() {
		b.Passed++
		return
	}
	b.TopCards = cards
}

func hasClubsThree(cards []entity.BigTwoCard) bool {
	for _, v := range cards {
		if v == (entity.BigTwoCard{Suit: entity.Clubs, Rank: entity.Three}) {
			return true
		}
	}
	return false
}

// helpers

func isMatchSingle(cards []entity.BigTwoCard) bool {
	return len(cards) == 1
}

func isMatchPair(cards []entity.BigTwoCard) bool {
	if len(cards) == 2 && cards[0].Rank == cards[1].Rank {
		fmt.Println("paiir!!!")
		return true
	}
	return false
}
