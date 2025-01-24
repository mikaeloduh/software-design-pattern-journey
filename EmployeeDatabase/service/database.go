package service

type IDatabase interface {
	GetEmployeeById(id int) (IEmployee, error)
}
