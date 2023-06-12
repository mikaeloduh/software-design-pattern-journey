package entity

type IOutput interface {
	RenameOutput(name string)
	RoundStartOutput(i int)
	RoundResultOutput(i int, roundResults RoundResults)
	GameOverOutput(winner IPlayer, players []IPlayer)
	PrintCardsOutput(cards []ShowdownCard)
	TakeTurnStartOutput(name string)
	AskShowCardOutput(name string)
}
