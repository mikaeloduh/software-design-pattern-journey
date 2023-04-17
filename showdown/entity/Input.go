package entity

type Input interface {
	InputNum(int, int) int
	InputBool() bool
	InputString() string
}
