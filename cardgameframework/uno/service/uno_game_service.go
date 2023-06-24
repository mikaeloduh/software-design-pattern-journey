package service

import (
	"cardgameframework/template"
	"cardgameframework/uno/entity"
	"fmt"
)

// UnoGame represents the UNO game.
type UnoGame[T entity.UnoCard] struct {
	Players       []entity.IUnoPlayer[entity.UnoCard]
	Deck          template.Deck[entity.UnoCard]
	DeskCard      entity.UnoCard
	CurrentPlayer int
}

// NewUnoGame creates a new instance of the UnoGame.
func NewUnoGame(players []entity.IUnoPlayer[entity.UnoCard]) *template.GameFramework[entity.UnoCard] {
	deck := entity.NewUnoDeck()
	base := &template.GameFramework[entity.UnoCard]{
		Deck:        deck,
		Players:     make([]template.IPlayer[entity.UnoCard], len(players)),
		PlayingGame: &UnoGame[entity.UnoCard]{Players: players, Deck: deck},
	}
	for i, player := range players {
		base.Players[i] = player
	}

	return base
}

// PreTakeTurns run before TakeTurns
func (u *UnoGame[T]) PreTakeTurns() {
	// Start the game by placing a card from the deck on the desk
	fmt.Printf("Starting card on the desk. ")
	u.DeskCard = u.Deck.DealCard()
}

func (u *UnoGame[T]) TakeTurnStep(player template.IPlayer[entity.UnoCard]) {
	fmt.Printf("The desk card is %v\n", u.DeskCard)
	// Check if the player has a valid card to play
	haveValidCards := u.haveValidCards(player.GetHand())
	if haveValidCards {
		// Player has valid cards, let them choose a card to play
		var playerTakeTurn func() entity.UnoCard
		playerTakeTurn = func() entity.UnoCard {
			card := player.TakeTurn()
			if !u.isValidMove(card) {
				fmt.Printf("Invalid move, try again.\n")

				player.SetCard(card)
				playerTakeTurn()
			}
			return card
		}

		playedCard := playerTakeTurn()
		// Update the desk card
		u.updateDeskCard(playedCard)

		fmt.Printf("%s played %v\n", player.GetName(), playedCard)
	} else {
		fmt.Printf("\n%s's turn. \nYou have no valid cards\n", player.GetName())

		// Player has no valid cards, they need to draw a card from the deck
		u.Deck.Shuffle()
		drawnCard := u.Deck.DealCard()
		player.SetCard(drawnCard)
		fmt.Printf("%s drew a card: %v\n", player.GetName(), drawnCard)
	}
}

// GetCurrentPlayer returns the current turn player.
func (u *UnoGame[T]) GetCurrentPlayer() template.IPlayer[entity.UnoCard] {
	return u.Players[u.CurrentPlayer]
}

// UpdateGameAndMoveToNext moves the turn to the next player.
func (u *UnoGame[T]) UpdateGameAndMoveToNext() {
	u.CurrentPlayer = (u.CurrentPlayer + 1) % len(u.Players)
}

// IsGameFinished checks if the game is finished.
func (u *UnoGame[T]) IsGameFinished() bool {
	for _, player := range u.Players {
		if len(player.GetHand()) == 0 {
			return true
		}
	}
	return false
}

// GameResult processes the final result of the game.
func (u *UnoGame[T]) GameResult() (winner template.IPlayer[entity.UnoCard]) {
	for _, player := range u.Players {
		fmt.Printf("%s's hand: %v\n", player.GetName(), player.GetHand())
		if len(player.GetHand()) == 0 {
			fmt.Printf("%s has won the game!\n", player.GetName())
			winner = player
		}
	}
	return winner
}

/// Helpers: ///

// haveValidCards returns the valid cards that the player can play based on the current desk card.
func (u *UnoGame[T]) haveValidCards(hand []entity.UnoCard) bool {
	for _, card := range hand {
		if u.isValidMove(card) {
			return true
		}
	}
	return false
}

// isValidMove checks if a card is a valid move.
func (u *UnoGame[T]) isValidMove(card entity.UnoCard) bool {
	return true
}

// updateDeskCard updates the desk card and moves the old desk card to the deck.
func (u *UnoGame[T]) updateDeskCard(card entity.UnoCard) {
	u.Deck.Cards = append(u.Deck.Cards, u.DeskCard)
	u.DeskCard = card
}
