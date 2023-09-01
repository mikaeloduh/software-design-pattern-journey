package entity

// IState interface
type IState interface {
	OnRoundStart()
	OnTakeDamage(damage int) int
	OnAttack(attack AttackMap) AttackMap
}
