package service

import (
	"fmt"
	"math/rand"
	"time"
)

// Player defines the methods required for a player in the UNO game.
type Player interface {
	SetCard(card Card)
	TakeTurn() Card
	GetName() string
	GetHand() []Card
}

// HumanPlayer represents a human player in the UNO game.
type HumanPlayer struct {
	Name string
	Hand []Card
}

// SetCard adds a card to the player's hand.
func (p *HumanPlayer) SetCard(card Card) {
	p.Hand = append(p.Hand, card)
}

// TakeTurn allows the player to choose a card to play.
func (p *HumanPlayer) TakeTurn() Card {
	fmt.Printf("\n%s's turn. Your hand: %v\n", p.GetName(), p.GetHand())
	var cardIndex int
	for {
		fmt.Print("Enter the index of the card you want to play: ")
		_, err := fmt.Scan(&cardIndex)
		if err == nil && cardIndex >= 0 && cardIndex < len(p.Hand) {
			break
		}
		fmt.Println("Invalid input. Try again.")
	}
	card := p.Hand[cardIndex]
	p.Hand = append(p.Hand[:cardIndex], p.Hand[cardIndex+1:]...)
	return card
}

// GetName returns the player's name.
func (p *HumanPlayer) GetName() string {
	return p.Name
}

// GetHand returns the player's hand.
func (p *HumanPlayer) GetHand() []Card {
	return p.Hand
}

// ComputerPlayer represents a computer player in the UNO game.
type ComputerPlayer struct {
	Name string
	Hand []Card
}

// SetCard adds a card to the player's hand.
func (p *ComputerPlayer) SetCard(card Card) {
	p.Hand = append(p.Hand, card)
}

// TakeTurn randomly selects a card to play.
func (p *ComputerPlayer) TakeTurn() Card {
	cardIndex := rand.Intn(len(p.Hand))
	card := p.Hand[cardIndex]
	p.Hand = append(p.Hand[:cardIndex], p.Hand[cardIndex+1:]...)
	return card
}

// GetName returns the player's name.
func (p *ComputerPlayer) GetName() string {
	return p.Name
}

// GetHand returns the player's hand.
func (p *ComputerPlayer) GetHand() []Card {
	return p.Hand
}

// Card represents an UNO card.
type Card struct {
	Color string
	Value string
}

// Deck represents the UNO deck.
type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	colors := []string{"Red", "Blue", "Green", "Yellow"}
	values := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	deck := Deck{}
	for _, color := range colors {
		for _, value := range values {
			card := Card{Color: color, Value: value}
			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}

// Shuffle randomly shuffles the deck of cards.
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// DealCard deals a card from the deck.
func (d *Deck) DealCard() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

// UnoGame represents the UNO game.
type UnoGame struct {
	Players  []Player
	Deck     Deck
	DeskCard Card
}

// NewUnoGame creates a new instance of the UnoGame.
func NewUnoGame(players []Player, deck Deck) *UnoGame {
	return &UnoGame{
		Players: players,
		Deck:    deck,
	}
}

// ShuffleDeck shuffles the deck of cards.
func (u *UnoGame) ShuffleDeck() {
	u.Deck.Shuffle()
}

// DealHands deals the initial hands to all players.
func (u *UnoGame) DealHands(numCards int) {
	for i := 0; i < numCards; i++ {
		for _, player := range u.Players {
			card := u.Deck.DealCard()
			player.SetCard(card)
		}
	}
}

// TakeTurns allows each player to take their turn.
func (u *UnoGame) TakeTurns() {
	// Start the game by placing a card from the deck on the desk
	u.DeskCard = u.Deck.DealCard()
	fmt.Printf("Starting card on the desk. ")

	for {
		for _, player := range u.Players {
			fmt.Printf("The desk card is %v\n", u.DeskCard)
			// Check if the player has a valid card to play
			haveValidCards := u.haveValidCards(player.GetHand())
			if haveValidCards {
				// Player has valid cards, let them choose a card to play
				var playerTakeTurn func() Card
				playerTakeTurn = func() Card {
					card := player.TakeTurn()
					if !u.isValidMove(card) {
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

			// Check if the player has emptied their hand and won the game
			if len(player.GetHand()) == 0 {
				fmt.Printf("Game over!\n\n")
				return
			}
		}
	}
}

// GameResult processes the final result of the game.
func (u *UnoGame) GameResult() (winner Player) {
	for _, player := range u.Players {
		fmt.Printf("%s's hand: %v\n", player.GetName(), player.GetHand())
		if len(player.GetHand()) == 0 {
			fmt.Printf("\n%s has won the game!\n", player.GetName())
			winner = player
		}
	}
	return winner
}

// haveValidCards returns the valid cards that the player can play based on the current desk card.
func (u *UnoGame) haveValidCards(hand []Card) bool {
	for _, card := range hand {
		if u.isValidMove(card) {
			return true
		}
	}

	return false
}

// isValidMove checks if a card is a valid move.
func (u *UnoGame) isValidMove(card Card) bool {
	if card.Color == u.DeskCard.Color || card.Value == u.DeskCard.Value {
		return true
	}
	return false
}

// updateDeskCard updates the desk card and moves the old desk card to the deck.
func (u *UnoGame) updateDeskCard(card Card) {
	u.Deck.Cards = append(u.Deck.Cards, u.DeskCard)
	u.DeskCard = card
}
