package service

type RealEmployee struct {
	Id             int
	Name           string
	Age            int
	SubordinateIds []int
}

func (e *RealEmployee) GetSubordinates() Employee {
	//TODO implement me
	panic("implement me")
}
