package entity

type IState interface {
	OnAttack(damage int) int
	OnRoundStart()
}
