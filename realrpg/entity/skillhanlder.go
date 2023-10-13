package entity

import "reflect"

// ISkillHandler is the CoR interface
type ISkillHandler interface {
	Do(target IUnit)
}

// CoR handlers for OnePunch
type Hp500Handler struct {
	next ISkillHandler
}

func (h Hp500Handler) Do(target IUnit) {
	if target.GetHp() >= 500 {
		target.TakeDamage(300)
	} else if h.next != nil {
		h.next.Do(target)
	}
}

type PetrochemicalStateHandler struct {
	next ISkillHandler
}

func (h PetrochemicalStateHandler) Do(target IUnit) {
	if reflect.TypeOf(target.GetState()) == reflect.TypeOf(&PetrochemicalState{}) {
		target.TakeDamage(80)
	} else if h.next != nil {
		h.next.Do(target)
	}
}

type PoisonedStateHandler struct {
	next ISkillHandler
}

func (h PoisonedStateHandler) Do(target IUnit) {
	if reflect.TypeOf(target.GetState()) == reflect.TypeOf(&PoisonedState{}) {
		target.TakeDamage(80)
	} else if h.next != nil {
		h.next.Do(target)
	}
}

type CheerUpStateHandler struct {
	next ISkillHandler
}

func (h CheerUpStateHandler) Do(target IUnit) {
	if reflect.TypeOf(target.GetState()) == reflect.TypeOf(&CheerUpState{}) {
		target.TakeDamage(100)
		target.SetState(NewNormalState(target))
	} else if h.next != nil {
		h.next.Do(target)
	}
}

type NormalStateHandler struct {
	next ISkillHandler
}

func (h NormalStateHandler) Do(target IUnit) {
	if reflect.TypeOf(target.GetState()) == reflect.TypeOf(&NormalState{}) {
		target.TakeDamage(100)
	} else if h.next != nil {
		h.next.Do(target)
	}
}
