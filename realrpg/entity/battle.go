package entity

type Battle struct {
	troop1 Troop
	troop2 Troop
}

func (b *Battle) Init() {

}

func (b *Battle) StartRound() {
	for _, unit := range b.troop1 {
		unit.OnRoundStart()
	}
}
