package entity

type IInput interface {
	InputNum(int, int) int
	InputBool() bool
	InputString() string
}
