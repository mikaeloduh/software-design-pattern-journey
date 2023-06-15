package entity

type IShowdownInput interface {
	InputNum(int, int) int
	InputBool() bool
	InputString() string
}
