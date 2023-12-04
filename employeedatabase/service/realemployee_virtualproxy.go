package service

type RealEmployeeVirtualProxy struct {
	realEmployee IRealEmployee
	database     IDatabase
}

func NewRealEmployeeVirtualProxy(realEmployee IRealEmployee, database IDatabase) *RealEmployeeVirtualProxy {
	return &RealEmployeeVirtualProxy{
		realEmployee: realEmployee,
		database:     database,
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

func (p *RealEmployeeVirtualProxy) Subordinates() []IEmployee {
	if p.realEmployee.Subordinates() == nil {
		subordinates := make([]IEmployee, 0)
		for _, v := range p.realEmployee.SubordinateIds() {
			subordinate, _ := p.database.GetEmployeeById(v)
			subordinates = append(subordinates, subordinate)
		}

		p.realEmployee.SetSubordinates(subordinates)

		return subordinates
	}

	return p.realEmployee.Subordinates()
}
