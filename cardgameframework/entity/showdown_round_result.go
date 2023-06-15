package entity

type ShowdownRoundResult struct {
	Player IShowdownPlayer
	Card   ShowdownCard
	Win    bool
}

type RoundResults []ShowdownRoundResult
