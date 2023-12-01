package service

type Employee interface {
	GetSubordinates() Employee
}
