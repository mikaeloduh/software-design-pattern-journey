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

func (d *PatientDatabase) AddPatientCase(p *Patient, c Case) {
	d.data[p.Id].AddCase(c)
}

func (d *PatientDatabase) FindPatientByName(name string) *Patient {
	for _, v := range d.data {
		if v.Name == name {
			return v
		}
	}
	return nil
}
