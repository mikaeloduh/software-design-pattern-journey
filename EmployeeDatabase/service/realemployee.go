package service

type IRealEmployee interface {
	IEmployee
	SetSubordinates(subordinates []IEmployee)
	SubordinateIds() []int
}

type RealEmployee struct {
	id             int
	name           string
	age            int
	subordinates   []IEmployee
	subordinateIds []int
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

func (e *RealEmployee) Subordinates() []IEmployee {
	return e.subordinates
}

func (e *RealEmployee) SetSubordinates(subordinates []IEmployee) {
	e.subordinates = subordinates
}

func (e *RealEmployee) SubordinateIds() []int {
	return e.subordinateIds
}
