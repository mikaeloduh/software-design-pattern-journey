package entity

import (
	"fmt"
	"realrpg/utils"
)

type Hero struct {
	Name      string
	MaxHP     int
	CurrentHP int
	MaxMP     int
	CurrentMP int
	STR       int
	State     IState
	Skills    []ISkill
	skillIdx  int
	troop     *Troop
	observers []IObserver
}

func NewHero(name string) *Hero {
	h := &Hero{
		Name:      name,
		MaxHP:     1000,
		CurrentHP: 1000,
		MaxMP:     900,
		CurrentMP: 900,
		STR:       50,
	}
	h.SetState(NewNormalState(h))
	h.AddSkill(NewBasicAttack(h))

	return h
}

func (u *Hero) AddSkill(skill ISkill) {
	u.Skills = append(u.Skills, skill)
}

func (u *Hero) OnRoundStart() {
	u.State.OnRoundStart()
}

func (u *Hero) TakeTurn(candidateTargets []IUnit) {
	// Select skill
	// TODO: prompt select skill message
	if err := u.selectSkill(0); err != nil {
		return
	}
	// Select targets
	// TODO: prompt select target message
	_ = u.selectTarget(nil)
	// Consume CurrentMP and take action
	_ = u.doSkill(nil)
}

func (u *Hero) selectSkill(i int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("index out of range")
		}
	}()

	if !u.Skills[i].IsMpEnough() {
		err = fmt.Errorf("not enough CurrentMP")
	}

	u.skillIdx = i

	return err
}

func (u *Hero) getSelectedSkill() ISkill {
	return u.Skills[u.skillIdx]
}

func (u *Hero) selectTarget(targets ...IUnit) error {
	return u.getSelectedSkill().BeforeDo(targets...)
}

func (u *Hero) doSkill(targets ...IUnit) error {
	return u.getSelectedSkill().Do(targets...)
}

func (u *Hero) TakeDamage(damage int) {
	u.CurrentHP = utils.MinMaxLimit(0, u.MaxHP, u.CurrentHP-damage)

	if u.CurrentHP <= 0 {
		u.SetState(NewDeadState(u))
	}
}

func (u *Hero) ConsumeMp(mp int) {
	u.CurrentMP = utils.MinMaxLimit(0, u.MaxMP, u.CurrentMP-mp)
}

func (u *Hero) Register(skill IObserver) {
	u.observers = append(u.observers, skill)
}

func (u *Hero) UnRegister(skill IObserver) {
	var temp []IObserver
	for _, o := range u.observers {
		if o != skill {
			temp = append(temp, o)
		}
	}
	u.observers = temp
}

func (u *Hero) Notify() {
	for _, o := range u.observers {
		o.Update(u)
	}
}

func (u *Hero) GetHp() int {
	return u.CurrentHP
}

func (u *Hero) SetHp(hp int) {
	u.CurrentHP = hp
}

func (u *Hero) GetMp() int {
	return u.CurrentMP
}

func (u *Hero) GetSTR() int {
	return u.STR
}

func (u *Hero) GetState() IState {
	return u.State
}

func (u *Hero) SetState(s IState) {
	u.State = s
}

func (u *Hero) GetTroop() *Troop {
	return u.troop
}

func (u *Hero) SetTroop(t *Troop) {
	u.troop = t
}
