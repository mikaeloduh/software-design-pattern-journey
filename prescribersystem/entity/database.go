package entity

type PatientDatabase struct {
	data map[string]Patient
}

func NewPatientDatabase() *PatientDatabase {
	return &PatientDatabase{data: make(map[string]Patient)}
}

func (d *PatientDatabase) CreatePatient(p Patient) {
	d.data[p.Id] = p
}
