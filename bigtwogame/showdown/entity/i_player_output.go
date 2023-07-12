package entity

type IPlayerOutput interface {
	RenameOutput(name string)
	RoundStartOutput(i int)
	RoundResultOutput(i int, roundResults RoundResult)
	GameOverOutput(winner IShowdownPlayer, players []IShowdownPlayer)
	PrintCardsOutput(cards []ShowDownCard)
	TakeTurnStartOutput(name string)
	AskShowCardOutput(name string)
}
