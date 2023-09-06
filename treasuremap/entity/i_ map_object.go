package entity

type IMapObject interface {
	SetPosition(p *Position)
	GetPosition() *Position
}

type IStatefulMapObject interface {
	IMapObject
	OnRoundStart()
	AfterRoundStart()
	OnRoundEnd()
	isRoundEnd() bool
	SetState(s IState)
	SetSpeed(speed int)
	TakeDamage(damage Damage) (resultHp Hp)
	Heal(hp Hp)
	GetHp() Hp
	GetMaxHp() Hp
}
