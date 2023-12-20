package entity

import "fmt"

type NonHero struct {
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
	ai        AI
}

func NewNonHero(ai AI) *NonHero {
	h := &NonHero{
		Name:      "I-am-a-robot",
		MaxHP:     1000,
		CurrentHP: 1000,
		MaxMP:     900,
		CurrentMP: 900,
		STR:       50,
		ai:        ai,
	}
	h.SetState(NewNormalState(h))
	h.AddSkill(NewBasicAttack(h))

	return h
}

func (u *NonHero) AddSkill(skill ISkill) {
	u.Skills = append(u.Skills, skill)
}

func (u *NonHero) OnRoundStart() {
	u.State.OnRoundStart()
}

func (u *NonHero) TakeTurn(candidateTargets []IUnit) {
	// TODO: AI Select a skill
	if err := u.selectSkill(u.ai.RandAction(len(u.Skills))); err != nil {
		return
	}
	u.ai.IncrSeed()
	// TODO: AI Select the targets
	_ = u.selectTarget(nil)
	u.ai.IncrSeed()
	// Consume CurrentMP and take action
	_ = u.doSkill(nil)
}

// Privates
func (u *NonHero) selectSkill(i int) (err error) {
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

func (u *NonHero) getSelectedSkill() ISkill {
	return u.Skills[u.skillIdx]
}

func (u *NonHero) selectTarget(targets ...IUnit) error {
	return u.getSelectedSkill().BeforeDo(targets...)
}

func (u *NonHero) doSkill(targets ...IUnit) error {
	return u.getSelectedSkill().Do(targets...)
}

func (u *NonHero) TakeDamage(damage int) {
	result := u.CurrentHP - damage
	if result < 0 {
		result = 0
	} else if result > u.MaxHP {
		result = u.MaxHP
	}
	u.CurrentHP = result

	if u.CurrentHP <= 0 {
		u.SetState(NewDeadState(u))
	}
}

func (u *NonHero) ConsumeMp(mp int) {
	result := u.CurrentMP - mp
	if result < 0 {
		result = 0
	} else if result > u.MaxMP {
		result = u.MaxMP
	}
	u.CurrentMP = result
}

func (u *NonHero) Register(skill IObserver) {
	u.observers = append(u.observers, skill)
}

func (u *NonHero) UnRegister(skill IObserver) {
	var temp []IObserver
	for _, o := range u.observers {
		if o != skill {
			temp = append(temp, o)
		}
	}
	u.observers = temp
}

func (u *NonHero) Notify() {
	for _, o := range u.observers {
		o.Update(u)
	}
}

func (u *NonHero) GetHp() int {
	return u.CurrentHP
}

func (u *NonHero) SetHp(hp int) {
	u.CurrentHP = hp
}

func (u *NonHero) GetMp() int {
	return u.CurrentMP
}

func (u *NonHero) GetSTR() int {
	return u.STR
}

func (u *NonHero) GetState() IState {
	return u.State
}

func (u *NonHero) SetState(s IState) {
	u.State = s
}

func (u *NonHero) GetTroop() *Troop {
	return u.troop
}

func (u *NonHero) SetTroop(t *Troop) {
	u.troop = t
}
