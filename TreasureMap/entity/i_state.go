package entity

// IState interface
type IState interface {
	OnRoundStart()
	OnTakeDamage(damage Damage) Damage
	OnAttack(IAttackStrategy) IAttackStrategy
}
