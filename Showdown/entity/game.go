package entity

//type roundResult struct {
//	p   IPlayer
//	c   Card
//	win bool
//}
//type roundResults struct {
//	roundNo  int
//	roundRes roundResult
//}

type RoundResult struct {
	Player IPlayer
	Card   Card
	Win    bool
}

type RoundResults []RoundResult
