package entity

type ShowdownRoundResult struct {
	Player IPlayer
	Card   ShowdownCard
	Win    bool
}

type RoundResults []ShowdownRoundResult
