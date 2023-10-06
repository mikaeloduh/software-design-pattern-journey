package entity

import "fmt"

type IUnit interface {
	AddSkill(skill ISkill)
	selectSkill(i int) (ISkill, error)
	TakeAction()
	SetState()
	GetHp() int
	SetHp(int)
}

type Hero struct {
	Name     string
	HP       int
	MP       int
	STR      int
	Skills   []ISkill
	skillIdx int
}

func NewHero(name string) *Hero {
	str := 50
	return &Hero{
		Name:   name,
		HP:     1000,
		MP:     900,
		STR:    str,
		Skills: []ISkill{&BasicAttack{Damage: str}},
	}
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

func (u *Hero) SetState() {
	//TODO implement me
	panic("implement me")
}

func (u *Hero) GetHp() int {
	return u.HP
}

func (u *Hero) SetHp(hp int) {
	u.HP = hp
}
