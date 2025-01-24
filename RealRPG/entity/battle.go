package entity

import (
	"fmt"
	"reflect"
)

type Battle struct {
	troops []*Troop
	hero   IUnit
}

func NewBattle(troops ...*Troop) *Battle {
	b := &Battle{}
	for _, t := range troops {
		b.troops = append(b.troops, t)
	}

	b.Init()

	return b
}

func (b *Battle) Init() {
	for _, troop := range b.troops {
		for _, unit := range *troop {
			unit.TakeDamage(-99999)
			unit.ConsumeMp(-99999)
			unit.SetState(NewNormalState(unit))
		}
	}
}

func (b *Battle) Start() {
	b.OnRoundStart()

	for !b.IsGameFinished() {
		player := b.GetCurrentPlayer()

		b.TakeTurnStep(player)

		b.UpdateGameAndMoveToNext()
	}
}

func (b *Battle) OnRoundStart() {
	for _, troop := range b.troops {
		for _, unit := range *troop {
			if u, ok := unit.(IUnit); ok == true {
				u.OnRoundStart()
			}
		}
	}
}

func (b *Battle) IsGameFinished() bool {
	for _, troop := range b.troops {
		if b.isAnnihilated(troop) || b.isHeroDie(troop) {
			return true
		}
	}

	return false
}

func (b *Battle) GetCurrentPlayer() IUnit {
	t := b.getCurrentTroop()
	u := b.getCurrentUnit(t)

	return u
}

func (b *Battle) TakeTurnStep(unit IUnit) {
	var candidateTargets []IUnit
	for _, troop := range b.troops {
		for _, u := range *troop {
			if b.isAliveUnit(u) {
				candidateTargets = append(candidateTargets, u)
			}
		}
	}
	unit.TakeTurn(candidateTargets)
}

func (b *Battle) UpdateGameAndMoveToNext() {
	// if hit new round
	b.OnRoundStart()
}

func (b *Battle) GameResult() {
	var winner int
	for i, t := range b.troops {
		if !b.isAnnihilated(t) && !b.isHeroDie(t) {
			winner = i
		}
	}

	fmt.Printf("Winnter is %v troop", winner)
}

// privates
func (b *Battle) isAnnihilated(troop *Troop) bool {
	for _, u := range *troop {
		if b.isAliveUnit(u) {
			return false
		}
	}
	return true
}

func (b *Battle) getCurrentTroop() *Troop {
	troop := b.troops[0]
	b.troops = append(b.troops[1:], troop)

	return troop
}

func (b *Battle) getCurrentUnit(t *Troop) IUnit {
	unit := (*t)[0]
	*t = append((*t)[1:], unit)
	if !b.isActiveUnit(unit) {
		unit = b.getCurrentUnit(t)
	}

	return unit
}

func (b *Battle) isActiveUnit(unit IUnit) bool {
	if reflect.TypeOf(unit.GetState()) == reflect.TypeOf(&DeadState{}) {
		return false
	}
	if reflect.TypeOf(unit.GetState()) == reflect.TypeOf(&PetrochemicalState{}) {
		return false
	}

	return true
}

func (b *Battle) isAliveUnit(unit IUnit) bool {
	if reflect.TypeOf(unit.GetState()) == reflect.TypeOf(&DeadState{}) {
		return false
	}

	return true
}

func (b *Battle) isHeroDie(troop *Troop) bool {
	for _, u := range *troop {
		if _, ok := u.(*Hero); ok == true {
			return !b.isAliveUnit(u)
		}
	}
	return false
}
