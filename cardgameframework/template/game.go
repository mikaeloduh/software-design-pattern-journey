package template

type IGame interface {
	Run()
}

type PlayingGame[T ICard] interface {
	TakeTurnStep(player IPlayer[T])
	GetCurrentPlayer() IPlayer[T]
	UpdateGameAndMoveToNext()
	IsGameFinished() bool
	GameResult() IPlayer[T]
}

type GameFramework[T ICard] struct {
	Deck    Deck[T]
	Players []IPlayer[T]

	PlayingGame PlayingGame[T]
}

func (f *GameFramework[T]) Run() {
	f.Init()
	f.ShuffleDeck()
	f.DrawHands(5)
	f.PreTakeTurns()
	f.TakeTurns()
	f.GameResult()
}

func (f *GameFramework[T]) Init() {
	for _, p := range f.Players {
		p.Rename()
	}
}

func (f *GameFramework[T]) ShuffleDeck() {
	f.Deck.Shuffle()
}

func (f *GameFramework[T]) DrawHands(numCards int) {
	for i := 0; i < numCards; i++ {
		for _, p := range f.Players {
			p.SetCard(f.Deck.DealCard())
		}
	}
}

func (f *GameFramework[T]) PreTakeTurns() {}

func (f *GameFramework[T]) TakeTurns() {
	for !f.PlayingGame.IsGameFinished() {
		player := f.PlayingGame.GetCurrentPlayer()

		f.PlayingGame.TakeTurnStep(player)

		f.PlayingGame.UpdateGameAndMoveToNext()
	}
}

func (f *GameFramework[T]) GameResult() IPlayer[T] {
	return f.PlayingGame.GameResult()
}
