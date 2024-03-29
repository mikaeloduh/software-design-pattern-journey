package entity

import (
	"fmt"
	"io"
)

type Slime struct {
	MaxHP     int
	CurrentHP int
	MaxMP     int
	CurrentMP int
	STR       int
	State     IState
	Skills    []ISkill
	skillIdx  int
	troop     *Troop
	writer    io.Writer
	observers []IObserver
	ai        AI
}

func NewSlime(ai AI, writer io.Writer) *Slime {
	s := &Slime{
		MaxHP:     100,
		CurrentHP: 100,
		MaxMP:     0,
		CurrentMP: 0,
		STR:       50,
		writer:    writer,
		ai:        ai,
	}
	s.SetState(NewNormalState(s))
	s.AddSkill(NewBasicAttack(s))

	return s
}

func (u *Slime) AddSkill(skill ISkill) {
	u.Skills = append(u.Skills, skill)
}

func (u *Slime) OnRoundStart() {
	u.State.OnRoundStart()

	fmt.Fprintf(u.writer, "Slime is taking action")
}

func (u *Slime) TakeTurn(targets []IUnit) {
	basicAttack := u.Skills[0]
	// TODO: AI Select the targets
	basicAttack.BeforeDo(targets...)
	basicAttack.Do(nil)
	u.ai.IncrSeed()
}

func (u *Slime) TakeDamage(damage int) {
	result := u.CurrentHP - damage
	if result < 0 {
		result = 0
	} else if result > u.MaxHP {
		result = u.MaxHP
	}
	u.CurrentHP = result

	if u.CurrentHP <= 0 {
		u.Notify()
	}
}

func (u *Slime) ConsumeMp(_ int) {
	panic("invalid operation")
}

func (u *Slime) Register(skill IObserver) {
	u.observers = append(u.observers, skill)
}

func (u *Slime) UnRegister(skill IObserver) {
	var temp []IObserver
	for _, o := range u.observers {
		if o != skill {
			temp = append(temp, o)
		}
	}
	u.observers = temp
}

func (u *Slime) Notify() {
	for _, o := range u.observers {
		o.Update(u)
	}
}

func (u *Slime) GetHp() int {
	return u.CurrentHP
}

func (u *Slime) SetHp(hp int) {
	u.CurrentHP = hp
}

func (u *Slime) GetMp() int {
	return u.CurrentMP
}

func (u *Slime) GetSTR() int {
	return u.STR
}

func (u *Slime) SetState(s IState) {
	u.State = s
}

func (u *Slime) GetState() IState {
	return u.State
}

func (u *Slime) GetTroop() *Troop {
	return u.troop
}

func (u *Slime) SetTroop(t *Troop) {
	u.troop = t
}
