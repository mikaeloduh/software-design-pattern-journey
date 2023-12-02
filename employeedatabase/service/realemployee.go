package service

type RealEmployee struct {
	id             int
	name           string
	age            int
	subordinateIds []int
	subordinates   []Employee
}

func NewRealEmployee(id int, name string, age int, subordinateIds []int) *RealEmployee {
	return &RealEmployee{id: id, name: name, age: age, subordinateIds: subordinateIds}
}

func (e *RealEmployee) Id() int {
	return e.id
}

func (e *RealEmployee) Name() string {
	return e.name
}

func (e *RealEmployee) Age() int {
	return e.age
}

func (e *RealEmployee) SubordinateIds() []int {
	return e.subordinateIds
}

func (e *RealEmployee) Subordinates() []Employee {
	return e.subordinates
}

func (e *RealEmployee) SetSubordinates(subordinates []Employee) {
	e.subordinates = subordinates
}
