package entity

type IShowdownOutput interface {
	RenameOutput(name string)
	RoundStartOutput(i int)
	RoundResultOutput(i int, roundResults RoundResults)
	GameOverOutput(winner IShowdownPlayer, players []IShowdownPlayer)
	PrintCardsOutput(cards []ShowdownCard)
	TakeTurnStartOutput(name string)
	AskShowCardOutput(name string)
}
