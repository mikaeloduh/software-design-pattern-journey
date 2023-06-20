package service

import (
	"cardgameframework/uno/entity"
	"fmt"
)

// UnoGame represents the UNO game.
type UnoGame struct {
	Players       []entity.IPlayer
	Deck          entity.Deck
	DeskCard      entity.Card
	CurrentPlayer int
}

// NewUnoGame creates a new instance of the UnoGame.
func NewUnoGame(players []entity.IPlayer, deck entity.Deck) *UnoGame {
	return &UnoGame{
		Players: players,
		Deck:    deck,
	}
}

func (u *UnoGame) Run() {
	u.Init()
	u.ShuffleDeck()
	u.DrawHands(5)
	u.PreTakeTurns()
	u.TakeTurns()
	u.GameResult()
}

func (u *UnoGame) Init() {
	// TODO: rename player
}

// ShuffleDeck shuffles the deck of cards.
func (u *UnoGame) ShuffleDeck() {
	u.Deck.Shuffle()
}

// DrawHands deals the initial hands to all players.
func (u *UnoGame) DrawHands(numCards int) {
	for i := 0; i < numCards; i++ {
		for _, p := range u.Players {
			p.SetCard(u.Deck.DealCard())
		}
	}
}

// PreTakeTurns run before TakeTurns
func (u *UnoGame) PreTakeTurns() {
	// Start the game by placing a card from the deck on the desk
	u.DeskCard = u.Deck.DealCard()
	fmt.Printf("Starting card on the desk. ")
}

// TakeTurns allows each player to take their turn.
func (u *UnoGame) TakeTurns() {
	for !u.IsGameFinished() {
		player := u.GetCurrentPlayer()

		u.TakeTurnStep(player)

		u.UpdateGameAndMoveToNext()
	}
}

func (u *UnoGame) TakeTurnStep(player entity.IPlayer) {
	fmt.Printf("The desk card is %v\n", u.DeskCard)
	// Check if the player has a valid card to play
	haveValidCards := u.haveValidCards(player.GetHand())
	if haveValidCards {
		// Player has valid cards, let them choose a card to play
		var playerTakeTurn func() entity.Card
		playerTakeTurn = func() entity.Card {
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
func (u *UnoGame) GetCurrentPlayer() entity.IPlayer {
	return u.Players[u.CurrentPlayer]
}

// UpdateGameAndMoveToNext moves the turn to the next player.
func (u *UnoGame) UpdateGameAndMoveToNext() {
	u.CurrentPlayer = (u.CurrentPlayer + 1) % len(u.Players)
}

// IsGameFinished checks if the game is finished.
func (u *UnoGame) IsGameFinished() bool {
	for _, player := range u.Players {
		if len(player.GetHand()) == 0 {
			return true
		}
	}
	return false
}

// GameResult processes the final result of the game.
func (u *UnoGame) GameResult() (winner entity.IPlayer) {
	for _, player := range u.Players {
		fmt.Printf("%s's hand: %v\n", player.GetName(), player.GetHand())
		if len(player.GetHand()) == 0 {
			fmt.Printf("%s has won the game!\n", player.GetName())
			winner = player
		}
	}
	return winner
}

// haveValidCards returns the valid cards that the player can play based on the current desk card.
func (u *UnoGame) haveValidCards(hand []entity.Card) bool {
	for _, card := range hand {
		if u.isValidMove(card) {
			return true
		}
	}
	return false
}

// isValidMove checks if a card is a valid move.
func (u *UnoGame) isValidMove(card entity.Card) bool {
	if card.Color == u.DeskCard.Color || card.Value == u.DeskCard.Value {
		return true
	}
	return false
}

// updateDeskCard updates the desk card and moves the old desk card to the deck.
func (u *UnoGame) updateDeskCard(card entity.Card) {
	u.Deck.Cards = append(u.Deck.Cards, u.DeskCard)
	u.DeskCard = card
}
