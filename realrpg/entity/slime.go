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
	Writer    io.Writer
}

func NewSlime(writer io.Writer) *Slime {
	s := &Slime{
		MaxHP:     100,
		CurrentHP: 100,
		MaxMP:     0,
		CurrentMP: 0,
		STR:       50,
		Writer:    writer,
	}
	s.SetState(NewNormalState(s))
	s.AddSkill(NewBasicAttack(s))

	return s
}

func (u *Slime) AddSkill(skill ISkill) {
	u.Skills = append(u.Skills, skill)
}

func (u *Slime) OnRoundStart() {
	u.TakeAction()
}

func (u *Slime) TakeAction() {
	u.State.OnRoundStart()

	fmt.Fprintf(u.Writer, "Slime is taking action")
}

func (u *Slime) SetState(s IState) {
	u.State = s
}

func (u *Slime) GetHp() int {
	//TODO implement me
	panic("implement me")
}

func (u *Slime) SetHp(i int) {
	//TODO implement me
	panic("implement me")
}

func (u *Slime) GetMp() int {
	//TODO implement me
	panic("implement me")
}

func (u *Slime) GetSTR() int {
	//TODO implement me
	panic("implement me")
}

func (u *Slime) GetState() IState {
	//TODO implement me
	panic("implement me")
}

func (u *Slime) TakeDamage(damage int) {
	//TODO implement me
	panic("implement me")
}

func (u *Slime) ConsumeMp(mp int) {
	//TODO implement me
	panic("implement me")
}

func (u *Slime) GetTroop() *Troop {
	return u.troop
}

func (u *Slime) SetTroop(t *Troop) {
	u.troop = t
}
