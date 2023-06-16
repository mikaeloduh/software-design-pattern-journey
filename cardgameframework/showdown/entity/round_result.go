package entity

type RoundResult struct {
	Player IPlayer
	Card   Card
	Win    bool
}

type RoundResults []RoundResult
