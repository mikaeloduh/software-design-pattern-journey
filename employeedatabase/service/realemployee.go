package service

type RealEmployee struct {
	id             int
	name           string
	age            int
	subordinateIds []int
}

func NewRealEmployee(id int, name string, age int, subordinateIds []int) *RealEmployee {
	return &RealEmployee{id: id, name: name, age: age, subordinateIds: subordinateIds}
}

func (e *RealEmployee) Id() int {
	return e.id
}

func (e *RealEmployee) SetId(id int) {
	e.id = id
}

func (e *RealEmployee) Name() string {
	return e.name
}

func (e *RealEmployee) SetName(name string) {
	e.name = name
}

func (e *RealEmployee) Age() int {
	return e.age
}

func (e *RealEmployee) SetAge(age int) {
	e.age = age
}

func (e *RealEmployee) SetSubordinateIds(subordinateIds []int) {
	e.subordinateIds = subordinateIds
}

func (e *RealEmployee) GetSubordinates() Employee {
	//TODO implement me

	panic("implement me")
}
