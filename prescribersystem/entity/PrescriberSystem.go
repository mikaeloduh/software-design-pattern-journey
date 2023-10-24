package entity

type Prescription struct {
	Name             string
	PotentialDisease string
	Medicines        string
	Usage            string
}

type PrescriberSystem struct {
	db         *PatientDatabase
	prescriber Prescriber
}

func NewPrescriberSystem(db *PatientDatabase) *PrescriberSystem {
	return &PrescriberSystem{db: db}
}

type Prescriber struct {
}

func (p *Prescriber) Diagnose(patient Patient, symptoms []Symptom) Prescription {

	return Prescription{}
}
