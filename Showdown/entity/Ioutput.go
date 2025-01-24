package entity

type IOutput interface {
	RenameOutput(name string)
	RoundStartOutput(i int)
	RoundResultOutput(i int, roundResults RoundResults)
	GameOverOutput(winner IPlayer, players []IPlayer)
	YouExchangeMyCardOutput(name string)
	MeExchangeYourCardOutput()
	MeExchangeYourCardErrorOutput(err error)
	PrintCardsOutput(cards []Card)
	AskToExchangeCardOutput(name string)
	ToExchangeCardOutput()
	TakeTurnStartOutput(name string)
	ExchangeBackOutput()
	AskShowCardOutput(name string)
}
