package template

type IGame interface {
	Run()
}

type GameFramework struct {
	Deck    IDeck
	Players []IPlayer
}

func (f *GameFramework) Run() {
	f.Init()
	f.ShuffleDeck()
	f.DrawHands()
	f.PreTakeTurns()
	f.TakeTurns()
	f.GameResult()
}

func (f *GameFramework) Init() {
	for _, p := range f.Players {
		p.Rename()
	}
}

func (f *GameFramework) ShuffleDeck() {
	u.Deck.Shuffle()
}

func (f *GameFramework) DrawHands(numCards int) {
	for i := 0; i < numCards; i++ {
		for _, p := range f.Players {
			p.SetCard(f.Deck.DealCard())
		}
	}
}

func (f *GameFramework) PreTakeTurns() {}

func (f *GameFramework) TakeTurns() {
	for !f.IsGameFinished() {
		player := f.GetCurrentPlayer()

		f.TakeTurnStep(player)

		f.UpdateGameAndMoveToNext()
	}
}

func (f *GameFramework) GameResult() (winner IPlayer) {}

/// Functions used by TakeTurns ///

func (f *GameFramework) TakeTurnStep(player IPlayer) {}

func (f *GameFramework) GetCurrentPlayer() IPlayer {}

func (f *GameFramework) UpdateGameAndMoveToNext() {}

func (f *GameFramework) IsGameFinished() bool {}

// 好了，被發現了，這個程式是不能動的，繼承語法也是假的，Go 就沒有，
// 樣板方法在 Go 十分困難，我不想用 work around 破壞 Go 的慣用風格
// 這題就算我敗了吧 ༼´◓ɷ◔`༽
