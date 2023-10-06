package entity

type IUnit interface {
	TakeAction()
	SetState()
	GetHp() int
	SetHp(int)
}

type Hero struct {
	Skills []ISkill
	HP     int
}

func NewHero(skills []ISkill) *Hero {
	return &Hero{
		Skills: skills,
		HP:     100,
	}
}

func (u *Hero) TakeAction() {
	//TODO implement me
	panic("implement me")
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
