package entity

type Case struct {
}

type Patient struct {
	Id           string
	Name         string
	Gender       Gender
	Age          int
	Height       float32
	Weight       float32
	PatientCases Case
}
