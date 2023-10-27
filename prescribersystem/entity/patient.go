package entity

import "time"

type Case struct {
	CaseTime time.Time
}

type Patient struct {
	Id     string
	Name   string
	Gender Gender
	Age    int
	Height float32
	Weight float32
	Cases  map[time.Time]Case
}

func NewPatient(id string, name string, gender Gender, age int, height float32, weight float32) *Patient {
	return &Patient{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Height: height,
		Weight: weight,
		Cases:  make(map[time.Time]Case),
	}
}

func (p *Patient) AddCase(c Case) {
	p.Cases[c.CaseTime] = c
}
