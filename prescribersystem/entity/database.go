package entity

type PatientDatabase struct {
	data map[string]*Patient
}

func NewPatientDatabase() *PatientDatabase {
	return &PatientDatabase{data: make(map[string]*Patient)}
}

func (d *PatientDatabase) CreatePatient(p *Patient) {
	d.data[p.Id] = p
}

func (d *PatientDatabase) FindPatientById(id string) *Patient {
	return d.data[id]
}
