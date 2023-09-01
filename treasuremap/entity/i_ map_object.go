package entity

type IMapObject interface {
	OnRoundStart()
	AfterRoundStart()
	OnRoundEnd()
	isRoundEnd() bool
	SetPosition(p *Position)
	SetState(s IState)
	SetSpeed(speed int)
}
