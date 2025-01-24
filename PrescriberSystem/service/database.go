package service

import "prescribersystem/entity"

type PatientDatabase struct {
	data map[string]*entity.Patient
}

func NewPatientDatabase() *PatientDatabase {
	return &PatientDatabase{data: make(map[string]*entity.Patient)}
}

func (d *PatientDatabase) CreatePatient(p *entity.Patient) {
	d.data[p.Id] = p
}

func (d *PatientDatabase) FindPatientById(id string) *entity.Patient {
	return d.data[id]
}

func (d *PatientDatabase) AddPatientCase(p *entity.Patient, c entity.Case) {
	d.data[p.Id].AddCase(c)
}

func (d *PatientDatabase) FindPatientByName(name string) *entity.Patient {
	for _, v := range d.data {
		if v.Name == name {
			return v
		}
	}
	return nil
}
