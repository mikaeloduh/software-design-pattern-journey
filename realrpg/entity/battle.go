package entity

type Battle struct {
	troop1 *Troop
	troop2 *Troop
}

func NewBattle(troop1 *Troop, troop2 *Troop) *Battle {
	b := &Battle{troop1: troop1, troop2: troop2}
	b.Init()

	return b
}

func (b *Battle) Init() {

}

func (b *Battle) StartRound() {
	for _, unit := range *b.troop1 {
		unit.OnRoundStart()
	}
}
