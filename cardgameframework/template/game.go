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

func (f *GameFramework[T]) GameResult() {
	f.PlayingGame.GameResult()
}

/// Functions used by TakeTurns ///

//func (f *GameFramework[T]) TakeTurnStep(player IPlayer) {}

//func (f *GameFramework[T]) GetCurrentPlayer() IPlayer {}

//func (f *GameFramework[T]) UpdateGameAndMoveToNext() {}

//func (f *GameFramework[T]) IsGameFinished() bool {}

// 好了，被發現了，這個程式是不能動的，繼承語法也是假的，Go 就沒有，
// 樣板方法在 Go 十分困難，我不想用 work around 破壞 Go 的慣用風格
// 這題就算我敗了吧 ༼´◓ɷ◔`༽
