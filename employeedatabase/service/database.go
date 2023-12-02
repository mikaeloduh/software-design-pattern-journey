package service

type Database interface {
	GetEmployeeById(id int) (Employee, error)
}
