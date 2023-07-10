package entity

import "fmt"

type AIPlayerOutput struct{}

func (ai AIPlayerOutput) AskShowCardOutput(name string) {
	fmt.Printf("%s (AI) is selecting card to show...\n", name)
}

func (ai AIPlayerOutput) TakeTurnStartOutput(name string) {
	fmt.Printf("\n* Now is %s (AI) 's turn.\n", name)
}

func (ai AIPlayerOutput) PrintCardsOutput([]ShowDownCard) {
}

func (ai AIPlayerOutput) GameOverOutput(winner IShowdownPlayer[ShowDownCard], players []IShowdownPlayer[ShowDownCard]) {
	fmt.Printf("\n============== Game Over ==============\nThe Winner is P%d: %s\n", winner.Id(), winner.GetName())
	for _, p := range players {
		fmt.Printf("%-8s: %d point\n", p.GetName(), p.Point())
	}
}

func (ai AIPlayerOutput) RoundResultOutput(i int, rrs RoundResult) {
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
		if rr.Win == true {
			fmt.Print(" win      ")
		}
		fmt.Print("          ")
	}
	fmt.Print("\n")
}

func (ai AIPlayerOutput) RenameOutput(string) {
	//TODO implement me
	panic("implement me")
}

func (ai AIPlayerOutput) RoundStartOutput(i int) {
	fmt.Printf("\n============== Round %d ==============\n", i)
}
