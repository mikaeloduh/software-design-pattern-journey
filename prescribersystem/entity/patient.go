package entity

import "time"

type Case struct {
	CaseTime time.Time
}

type Patient struct {
	Id           string
	Name         string
	Gender       Gender
	Age          int
	Height       float32
	Weight       float32
	PatientCases map[time.Time]Case
}

func (p *Patient) AddCase(c Case) {
	p.PatientCases[c.CaseTime] = c
}
