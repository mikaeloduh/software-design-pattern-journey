package service

type Employee interface {
	Id() int
	SetId(id int)
	Name() string
	SetName(name string)
	Age() int
	SetAge(age int)
	GetSubordinates() Employee
}
