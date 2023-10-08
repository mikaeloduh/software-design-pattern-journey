package entity

import "fmt"

type IUnit interface {
	AddSkill(skill ISkill)
	TakeAction()
	SetState(IState)
	GetHp() int
	SetHp(int)
	GetSTR() int
	GetState() IState
}

type Hero struct {
	Name     string
	HP       int
	MP       int
	STR      int
	State    IState
	Skills   []ISkill
	skillIdx int
}

func NewHero(name string) *Hero {
	h := &Hero{
		Name: name,
		HP:   1000,
		MP:   900,
		STR:  50,
	}
	h.SetState(NewNormalState(h))
	h.AddSkill(NewBasicAttack(h))

	return h
}

func (u *Hero) AddSkill(skill ISkill) {
	u.Skills = append(u.Skills, skill)
}

func (u *Hero) TakeAction() {
	// Select skill
	if _, err := u.selectSkill(0); err != nil {
		return
	}
	// Select targets
	u.selectTarget(nil)
	// Consume MP and take action
	u.doSkill()
}

func (u *Hero) selectSkill(i int) (ISkill, error) {
	if u.MP < u.Skills[i].GetMPCost() {
		return nil, fmt.Errorf("not enough MP")
	}

	u.skillIdx = i

	return u.Skills[i], nil
}

func (u *Hero) getSelectedSkill() ISkill {
	return u.Skills[u.skillIdx]
}

func (u *Hero) selectTarget(targets []IUnit) {
	skill := u.getSelectedSkill()
	skill.SelectTarget(targets)
}

func (u *Hero) doSkill() {
	skill := u.getSelectedSkill()
	u.MP -= skill.GetMPCost()

	skill.Do()
}

func (u *Hero) SetState(s IState) {
	u.State = s
}

func (u *Hero) GetHp() int {
	return u.HP
}

func (u *Hero) GetMp() int {
	return u.MP
}

func (u *Hero) SetHp(hp int) {
	u.HP = hp
}

func (u *Hero) GetSTR() int {
	return u.STR
}

func (u *Hero) GetState() IState {
	return u.State
}
