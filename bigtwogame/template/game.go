package template

type IGame interface {
	Run()
}

type PlayingGame[T ICard, F IPlayer[T]] interface {
	Init()
	PreTakeTurns()
	TakeTurnStep(player F)
	GetCurrentPlayer() F
	UpdateGameAndMoveToNext()
	IsGameFinished() bool
	GameResult() F
}

type GameFramework[T ICard, F IPlayer[T]] struct {
	Deck        *Deck[T]
	Players     []F
	NumCard     int
	PlayingGame PlayingGame[T, F]
}

func (f *GameFramework[T, F]) Run() {
	f.Init()
	f.ShuffleDeck()
	f.DrawHands(f.NumCard)
	f.PreTakeTurns()
	f.TakeTurns()
	f.GameResult()
}

func (f *GameFramework[T, F]) Init() {
	for _, p := range f.Players {
		p.Rename()
	}
	f.PlayingGame.Init()
}

func (f *GameFramework[T, F]) ShuffleDeck() {
	f.Deck.Shuffle()
}

func (f *GameFramework[T, F]) DrawHands(numCards int) {
	for i := 0; i < numCards; i++ {
		for _, p := range f.Players {
			p.SetCard(f.Deck.DealCard())
		}
	}
}

func (f *GameFramework[T, F]) PreTakeTurns() {
	f.PlayingGame.PreTakeTurns()
}

func (f *GameFramework[T, F]) TakeTurns() {
	for !f.PlayingGame.IsGameFinished() {
		player := f.PlayingGame.GetCurrentPlayer()

		f.PlayingGame.TakeTurnStep(player)

		f.PlayingGame.UpdateGameAndMoveToNext()
	}
}

func (f *GameFramework[T, F]) GameResult() IPlayer[T] {
	return f.PlayingGame.GameResult()
}
