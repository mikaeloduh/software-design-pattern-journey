package entity

import "fmt"

type HumanOutput struct{}

func (p HumanOutput) AskShowCardOutput(name string) {
	fmt.Printf("%s, please select a card to show ", name)
}

func (p HumanOutput) ExchangeBackOutput() {
	fmt.Println("Exchange back")
}

func (p HumanOutput) AskToExchangeCardOutput(name string) {
	fmt.Printf("%s, do you want to exchange hand card? ", name)
}

func (p HumanOutput) ToExchangeCardOutput() {
	fmt.Printf("Which player do you want to exchange cards with? ")
}

func (p HumanOutput) TakeTurnStartOutput(name string) {
	fmt.Printf("\n* Now is %s 's turn.\n", name)
}

func (p HumanOutput) PrintCardsOutput(cards []Card) {
	for i, c := range cards {
		if i%5 == 0 && i != 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%2d : [%4s ]  ", i, c.String())
	}
	fmt.Print("\n")
}

func (p HumanOutput) MeExchangeYourCardErrorOutput(err error) {
	fmt.Printf("Error: %v", err)
}

func (p HumanOutput) MeExchangeYourCardOutput() {
	fmt.Printf("Please select your card to exchange: ")
}

func (p HumanOutput) YouExchangeMyCardOutput(name string) {
	fmt.Printf("%s, please select your card to exchange back: ", name)
}

func (p HumanOutput) GameOverOutput(winner IPlayer, players []IPlayer) {
	fmt.Printf("\n============== Game Over ==============\nThe Winner is P%d: %s\n", winner.Id(), winner.Name())
	for _, player := range players {
		fmt.Printf("%-8s: %d point\n", player.Name(), player.Point())
	}
}

func (p HumanOutput) RoundResultOutput(i int, rrs RoundResults) {
	fmt.Printf("\n* Round %d end\n", i)
	for _, rr := range rrs {
		fmt.Printf("[%4s ]   ", rr.Card.String())
	}
	fmt.Print("\n")
	for _, rr := range rrs {
		fmt.Printf("%-8s  ", rr.Player.Name())
	}
	fmt.Print("\n")
	for _, rr := range rrs {
		if rr.Win {
			fmt.Print(" win      ")
		}
		fmt.Print("          ")
	}
	fmt.Print("\n")
}

func (p HumanOutput) RoundStartOutput(i int) {
	fmt.Printf("\n============== Round %d ==============\n", i)
}

func (p HumanOutput) RenameOutput(name string) {
	fmt.Printf("%s, please enter your name: ", name)
}
