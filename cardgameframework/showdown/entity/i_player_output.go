package entity

type IPlayerOutput interface {
	RenameOutput(name string)
	RoundStartOutput(i int)
	RoundResultOutput(i int, roundResults RoundResult)
	GameOverOutput(winner IPlayer, players []IPlayer)
	PrintCardsOutput(cards []Card)
	TakeTurnStartOutput(name string)
	AskShowCardOutput(name string)
}
