package entity

import "fmt"

type HumanPlayerOutput struct{}

func (p HumanPlayerOutput) AskShowCardOutput(name string) {
	fmt.Printf("%s, please select a card to show ", name)
}

func (p HumanPlayerOutput) TakeTurnStartOutput(name string) {
	fmt.Printf("\n* Now is %s 's turn.\n", name)
}

func (p HumanPlayerOutput) PrintCardsOutput(cards []ShowDownCard) {
	for i, c := range cards {
		if i%5 == 0 && i != 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%2d : [%4s ]  ", i, c.String())
	}
	fmt.Print("\n")
}

func (p HumanPlayerOutput) GameOverOutput(winner IShowdownPlayer[ShowDownCard], players []IShowdownPlayer[ShowDownCard]) {
	fmt.Printf("\n============== Game Over ==============\nThe Winner is P%d: %s\n", winner.Id(), winner.GetName())
	for _, player := range players {
		fmt.Printf("%-8s: %d point\n", player.GetName(), player.Point())
	}
}

func (p HumanPlayerOutput) RoundResultOutput(i int, rrs RoundResult) {
	fmt.Printf("\n* Round %d end\n", i)
	for _, rr := range rrs {
		fmt.Printf("[%4s ]   ", rr.Card.String())
	}
	fmt.Print("\n")
	for _, rr := range rrs {
		fmt.Printf("%-8s  ", rr.Player.GetName())
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

func (p HumanPlayerOutput) RoundStartOutput(i int) {
	fmt.Printf("\n============== Round %d ==============\n", i)
}

func (p HumanPlayerOutput) RenameOutput(name string) {
	fmt.Printf("%s, please enter your Name: ", name)
}
