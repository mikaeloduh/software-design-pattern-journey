package service

import (
	"fmt"
	"os"
)

type RealDatabaseProtectionProxy struct {
	nextDatabaseProxy IDatabase
}

func NewRealDatabaseProtectionProxy(nextDatabaseProxy IDatabase) *RealDatabaseProtectionProxy {
	return &RealDatabaseProtectionProxy{
		nextDatabaseProxy: nextDatabaseProxy,
	}
}

func (d *RealDatabaseProtectionProxy) GetEmployeeById(id int) (IEmployee, error) {
	if checkPassword() == false {
		return nil, fmt.Errorf("invalid passowrd")
	}

	employee, _ := d.nextDatabaseProxy.GetEmployeeById(id)

	return employee, nil
}

func checkPassword() bool {
	return os.Getenv("PASSWORD") == "1qaz2wsx"
}
