package entity

import "reflect"

// CoR handlers for OnePunch
type Hp500Handler struct {
	next ISkillHandler
}

func (h Hp500Handler) Do(target IUnit, unit IUnit) {
	if target.GetHp() >= 500 {
		damage := unit.GetState().OnAttack(300)
		target.TakeDamage(damage)
	} else if h.next != nil {
		h.next.Do(target, unit)
	}
}

type PetrochemicalStateHandler struct {
	next ISkillHandler
}

func (h PetrochemicalStateHandler) Do(target IUnit, unit IUnit) {
	if reflect.TypeOf(target.GetState()) == reflect.TypeOf(&PetrochemicalState{}) {
		damage := unit.GetState().OnAttack(80)
		target.TakeDamage(damage)
	} else if h.next != nil {
		h.next.Do(target, unit)
	}
}

type PoisonedStateHandler struct {
	next ISkillHandler
}

func (h PoisonedStateHandler) Do(target IUnit, unit IUnit) {
	if reflect.TypeOf(target.GetState()) == reflect.TypeOf(&PoisonedState{}) {
		damage := unit.GetState().OnAttack(80)
		target.TakeDamage(damage)
	} else if h.next != nil {
		h.next.Do(target, unit)
	}
}

type CheerUpStateHandler struct {
	next ISkillHandler
}

func (h CheerUpStateHandler) Do(target IUnit, unit IUnit) {
	if reflect.TypeOf(target.GetState()) == reflect.TypeOf(&CheerUpState{}) {
		damage := unit.GetState().OnAttack(100)
		target.TakeDamage(damage)
		target.SetState(NewNormalState(target))
	} else if h.next != nil {
		h.next.Do(target, unit)
	}
}

type NormalStateHandler struct {
	next ISkillHandler
}

func (h NormalStateHandler) Do(target IUnit, unit IUnit) {
	if reflect.TypeOf(target.GetState()) == reflect.TypeOf(&NormalState{}) {
		damage := unit.GetState().OnAttack(100)
		target.TakeDamage(damage)
	} else if h.next != nil {
		h.next.Do(target, unit)
	}
}
