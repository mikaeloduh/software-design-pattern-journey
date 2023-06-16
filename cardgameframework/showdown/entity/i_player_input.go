package entity

type IPlayerInput interface {
	InputNum(int, int) int
	InputBool() bool
	InputString() string
}
