package entity

import "fmt"

type ShowdownAIOutput struct{}

func (ai ShowdownAIOutput) AskShowCardOutput(name string) {
	fmt.Printf("%s (AI) is selecting card to show...\n", name)
}

func (ai ShowdownAIOutput) TakeTurnStartOutput(name string) {
	fmt.Printf("\n* Now is %s (AI) 's turn.\n", name)
}

func (ai ShowdownAIOutput) PrintCardsOutput([]ShowdownCard) {
}

func (ai ShowdownAIOutput) GameOverOutput(winner IShowdownPlayer, players []IShowdownPlayer) {
	fmt.Printf("\n============== Game Over ==============\nThe Winner is P%d: %s\n", winner.Id(), winner.Name())
	for _, p := range players {
		fmt.Printf("%-8s: %d point\n", p.Name(), p.Point())
	}
}

func (ai ShowdownAIOutput) RoundResultOutput(i int, rrs RoundResults) {
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

func (ai ShowdownAIOutput) RenameOutput(string) {
	//TODO implement me
	panic("implement me")
}

func (ai ShowdownAIOutput) RoundStartOutput(i int) {
	fmt.Printf("\n============== Round %d ==============\n", i)
}
