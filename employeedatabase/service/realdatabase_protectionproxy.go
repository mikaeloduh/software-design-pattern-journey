package service

import (
	"fmt"
	"os"
)

type RealDatabaseProtectionProxy struct {
	nextDatabaseProxy IDatabase
}

func NewRealDatabaseProtectionProxy() *RealDatabaseProtectionProxy {
	return &RealDatabaseProtectionProxy{
		nextDatabaseProxy: NewRealDatabaseVirtualProxy(),
	}
}

func (d *RealDatabaseProtectionProxy) GetEmployeeById(id int) (IEmployee, error) {
	if os.Getenv("PASSWORD") != "1qaz2wsx" {
		return nil, fmt.Errorf("invalid passowrd")
	}

	employee, _ := d.nextDatabaseProxy.GetEmployeeById(id)

	return employee, nil
}