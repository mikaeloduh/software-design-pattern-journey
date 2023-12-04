package service

type IEmployee interface {
	Id() int
	Name() string
	Age() int
	Subordinates() []IEmployee
}
