package service

type Employee interface {
	Id() int
	Name() string
	Age() int
	Subordinates() []Employee
}
