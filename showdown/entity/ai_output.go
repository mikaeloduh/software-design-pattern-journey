package entity

import "fmt"

type AIOutput struct{}

func (ai AIOutput) MeExchangeYourCardOutput() {
	fmt.Printf("Player (AI) is sele	cting card...\n")
}

func (ai AIOutput) MeExchangeYourCardErrorOutput(_ error) {
	//TODO implement me
	panic("implement me")
}

func (ai AIOutput) AskShowCardOutput(name string) {
	fmt.Printf("%s (AI) is selecting card to show...\n", name)
}

func (ai AIOutput) ExchangeBackOutput() {
	fmt.Println("Exchange back")
}

func (ai AIOutput) AskToExchangeCardOutput(name string) {
	fmt.Printf("%s (AI) wants to exchange card \n", name)
}

func (ai AIOutput) ToExchangeCardOutput() {
	fmt.Print("(AI) wants to exchange card\n")
}

func (ai AIOutput) TakeTurnStartOutput(name string) {
	fmt.Printf("\n* Now is %s (AI) 's turn.\n", name)
}

func (ai AIOutput) PrintCardsOutput([]Card) {
}

func (ai AIOutput) YouExchangeMyCardOutput(_ string) {
	fmt.Printf("Player (AI) is selecting card...\n")
}

func (ai AIOutput) GameOverOutput(winner IPlayer, players []IPlayer) {
	fmt.Printf("\n============== Game Over ==============\nThe Winners is P%d: %s\n", winner.Id(), winner.Name())
	for _, p := range players {
		fmt.Printf("%-8s: %d point\n", p.Name(), p.Point())
	}
}

func (ai AIOutput) RoundResultOutput(i int, rrs RoundResults) {
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
		if rr.Win == true {
			fmt.Print(" win      ")
		}
		fmt.Print("          ")
	}
	fmt.Print("\n")
}

func (ai AIOutput) RenameOutput(string) {
	//TODO implement me
	panic("implement me")
}

func (ai AIOutput) RoundStartOutput(i int) {
	fmt.Printf("\n============== Round %d ==============\n", i)
}
