package entity

type PrescriberSystem struct {
	db         *PatientDatabase
	prescriber Prescriber
}

func NewPrescriberSystem(db *PatientDatabase) *PrescriberSystem {
	return &PrescriberSystem{db: db}
}
