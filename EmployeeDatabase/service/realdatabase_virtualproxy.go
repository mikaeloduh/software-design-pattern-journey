package service

type RealDatabaseVirtualProxy struct {
	realDatabase *RealDatabase
}

func NewRealDatabaseVirtualProxy() *RealDatabaseVirtualProxy {
	return &RealDatabaseVirtualProxy{}
}

func (p *RealDatabaseVirtualProxy) GetEmployeeById(id int) (IEmployee, error) {
	if p.realDatabase == nil {
		p.realDatabase = NewRealDatabase("data.txt")
		_ = p.realDatabase.Init()
	}
	employee, _ := p.realDatabase.GetEmployeeById(id)

	return NewRealEmployeeVirtualProxy(employee.(IRealEmployee), p.realDatabase), nil
}
