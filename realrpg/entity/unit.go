package entity

type IUnit interface {
	IObservable
	AddSkill(skill ISkill)
	OnRoundStart()
	TakeTurn(targets []IUnit)
	TakeDamage(damage int)
	ConsumeMp(mp int)
	GetHp() int
	SetHp(int)
	GetMp() int
	GetSTR() int
	GetState() IState
	SetState(IState)
	GetTroop() *Troop
	SetTroop(*Troop)
}
