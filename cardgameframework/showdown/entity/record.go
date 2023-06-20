package entity

type Result struct {
	Player IPlayer
	Card   Card
	Win    bool
}

type RoundResult []Result

type Record []RoundResult
