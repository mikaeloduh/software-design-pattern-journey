package entity

type IUnit interface {
	AddSkill(skill ISkill)
	SelectSkill(i int) ISkill
	TakeAction()
	SetState()
	GetHp() int
	SetHp(int)
}

type Hero struct {
	Name   string
	HP     int
	STR    int
	Skills []ISkill
}

func NewHero(name string) *Hero {
	str := 50
	return &Hero{
		Name:   name,
		HP:     1000,
		STR:    str,
		Skills: []ISkill{&BasicAttack{Damage: str}},
	}
}

func (u *Hero) AddSkill(skill ISkill) {
	u.Skills = append(u.Skills, skill)
}

func (u *Hero) SelectSkill(i int) ISkill {
	return u.Skills[i]
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
