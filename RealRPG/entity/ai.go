package entity

type AI interface {
	IncrSeed()
	RandAction(num int) int
}
