package entity

type PatientDatabase struct {
	data map[string]Patient
}

func NewPatientDatabase() *PatientDatabase {
	return &PatientDatabase{}
}
