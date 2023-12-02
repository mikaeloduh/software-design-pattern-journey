package service

type RealEmployeeVirtualProxy struct {
	realEmployee *RealEmployee
	realDatabase *RealDatabase
}

func NewRealEmployeeProxy(realEmployee *RealEmployee, realDatabase *RealDatabase) *RealEmployeeVirtualProxy {
	return &RealEmployeeVirtualProxy{
		realEmployee: realEmployee,
		realDatabase: realDatabase,
	}
}

func (p *RealEmployeeVirtualProxy) Id() int {
	return p.realEmployee.Id()
}

func (p *RealEmployeeVirtualProxy) Name() string {
	return p.realEmployee.Name()
}

func (p *RealEmployeeVirtualProxy) Age() int {
	return p.realEmployee.Age()
}

func (p *RealEmployeeVirtualProxy) Subordinates() []Employee {
	if p.realEmployee.Subordinates() == nil {
		subordinates := make([]Employee, 0)
		for _, v := range p.realEmployee.SubordinateIds() {
			subordinate, _ := p.realDatabase.GetEmployeeById(v)
			subordinates = append(subordinates, subordinate)
		}

		p.realEmployee.SetSubordinates(subordinates)

		return subordinates
	}

	return p.realEmployee.Subordinates()
}
