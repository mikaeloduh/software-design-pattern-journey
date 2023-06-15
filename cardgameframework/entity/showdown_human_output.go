package entity

import "fmt"

type ShowdownHumanOutput struct{}

func (p ShowdownHumanOutput) AskShowCardOutput(name string) {
	fmt.Printf("%s, please select a card to show ", name)
}

func (p ShowdownHumanOutput) TakeTurnStartOutput(name string) {
	fmt.Printf("\n* Now is %s 's turn.\n", name)
}

func (p ShowdownHumanOutput) PrintCardsOutput(cards []ShowdownCard) {
	for i, c := range cards {
		if i%5 == 0 && i != 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%2d : [%4s ]  ", i, c.String())
	}
	fmt.Print("\n")
}

func (p ShowdownHumanOutput) GameOverOutput(winner IShowdownPlayer, players []IShowdownPlayer) {
	fmt.Printf("\n============== Game Over ==============\nThe Winner is P%d: %s\n", winner.Id(), winner.Name())
	for _, player := range players {
		fmt.Printf("%-8s: %d point\n", player.Name(), player.Point())
	}
}

func (p ShowdownHumanOutput) RoundResultOutput(i int, rrs RoundResults) {
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

func (p ShowdownHumanOutput) RoundStartOutput(i int) {
	fmt.Printf("\n============== Round %d ==============\n", i)
}

func (p ShowdownHumanOutput) RenameOutput(name string) {
	fmt.Printf("%s, please enter your name: ", name)
}
