package service

import (
	"fmt"
	"os"
)

type RealDatabaseProtectionProxy struct {
	realDatabase *RealDatabase
}

func NewRealDatabaseProtectionProxy() *RealDatabaseProtectionProxy {
	return &RealDatabaseProtectionProxy{
		realDatabase: NewRealDatabase("data.txt"),
	}
}

func (d *RealDatabaseProtectionProxy) Init() {
	err := d.realDatabase.Init()
	if err != nil {
		return
	}
}

func (d *RealDatabaseProtectionProxy) GetEmployeeById(id int) (Employee, error) {
	if os.Getenv("PASSWORD") != "1qaz2wsx" {
		return nil, fmt.Errorf("invalid passowrd")
	}

	realEmployee, _ := d.realDatabase.GetEmployeeById(id)

	return NewRealEmployeeProxy(realEmployee, d.realDatabase), nil
}
