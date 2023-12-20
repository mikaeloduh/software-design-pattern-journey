package entity

// ISkillHandler is the CoR interface
type ISkillHandler interface {
	Do(target IUnit, unit IUnit)
}
