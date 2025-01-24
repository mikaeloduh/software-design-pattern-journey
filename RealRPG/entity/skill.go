package entity

type ISkill interface {
	IObserver
	IsMpEnough() bool
	BeforeDo(targets ...IUnit) error
	Do(targets ...IUnit) error
}
