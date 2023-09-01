package entity

type IMapObject interface {
	OnRoundStart()
	AfterRoundStart()
	OnRoundEnd()
	isRoundEnd() bool
	SetPosition(p *Position)
	GetPosition() *Position
	SetState(s IState)
	SetSpeed(speed int)
	TakeDamage(damage int) (resultHp int)
	Heal(hp int)
	GetHp() int
	GetMaxHp() int
}
