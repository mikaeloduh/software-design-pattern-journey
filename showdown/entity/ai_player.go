package entity

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type AIPlayer struct {
	id           int
	name         string
	HandCards    []Card
	point        int
	usedExchange bool
	count        int
	who          IPlayer
}

func (ai *AIPlayer) inputBool() bool {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) == 1
}

func (ai *AIPlayer) inputNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return min + rand.Intn(max-min+1)
}

func (ai *AIPlayer) YouExchangeMyCard(card Card) (Card, error) {
	if len(ai.HandCards) < 1 {
		fmt.Printf("AI %d us fucking up", ai.id)
		return Card{}, errors.New(fmt.Sprintf("AI %d does not have enough cards to proceed with the exchange.", ai.id))
	}

	toPlay := ai.inputNum(0, len(ai.HandCards)-1)
	myCard := ai.HandCards[toPlay]
	ai.HandCards[toPlay] = card

	return myCard, nil
}

func (ai *AIPlayer) MeExchangeYourCard(player IPlayer) error {
	if len(ai.HandCards) < 1 {
		fmt.Println("yr fucking up")
		return errors.New(fmt.Sprintf("Player %d (AI) does not have enough cards to proceed with the exchange.", ai.id))
	}

	toPlay := ai.inputNum(0, len(ai.HandCards)-1)
	c := ai.HandCards[toPlay]

	ex, err := player.YouExchangeMyCard(c)
	if err != nil {
		return err
	}
	ai.HandCards[toPlay] = ex

	return nil
}

func (ai *AIPlayer) Point() int {
	return ai.point
}

func (ai *AIPlayer) AddPoint() {
	ai.point += 1
}

func (ai *AIPlayer) TakeTurn(players []IPlayer) Card {
	fmt.Printf("Player %d taking turn...\n", ai.id)

	// 1. exchange?
	if !ai.usedExchange {
		wantExchange := ai.inputBool()
		if wantExchange {
			var toExchangeCard func()
			toExchangeCard = func() {
				fmt.Println("before exchange")
				toExchange := ai.inputNum(0, 3)
				ai.who = players[toExchange] // TODO: inputNum
				if err := ai.MeExchangeYourCard(ai.who); err != nil {
					fmt.Println("try another exchange")
					toExchangeCard()
				}
			}
			toExchangeCard()
			ai.usedExchange = true
		}
	} else {
		ai.count--
		if ai.count == 0 {
			fmt.Println("before exchange back")
			_ = ai.MeExchangeYourCard(ai.who)
		}
	}

	// 2. Show card
	toPlay := ai.inputNum(0, len(ai.HandCards)-1)
	showCard := ai.HandCards[toPlay]
	ai.HandCards = append([]Card{}, append(ai.HandCards[0:toPlay], ai.HandCards[toPlay+1:]...)...)

	return showCard
}

func (ai *AIPlayer) Id() int {
	return ai.id
}

func (ai *AIPlayer) Name() string {
	return ai.name
}

func (ai *AIPlayer) GetCard(card Card) {
	ai.HandCards = append(ai.HandCards, card)
}

func NewAIPlayer(id int) *AIPlayer {
	return &AIPlayer{
		id:   id,
		name: "AI has no name",
	}
}
