package entity

type Troop []IUnit

func NewTroop(units ...IUnit) *Troop {
	troop := Troop(units)
	for _, unit := range units {
		unit.SetTroop(&troop)
	}

	return &troop
}

func (t *Troop) AddUnit(unit IUnit) {
	unit.SetTroop(t)
	*t = append(*t, unit)
}

func (t *Troop) Len() int {
	return len(*t)
}
