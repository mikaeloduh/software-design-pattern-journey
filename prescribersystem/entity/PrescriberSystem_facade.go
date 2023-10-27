package entity

import (
	"os"
	"time"
)

type Option string

const (
	JSON Option = "JSON"
	CSV  Option = "CSV"
)

type PrescriberSystemFacade struct {
	prescriberSystem *PrescriberSystem
}

func NewPrescriberSystemFacade() *PrescriberSystemFacade {
	p := &PrescriberSystemFacade{
		prescriberSystem: NewPrescriberSystem(NewPatientDatabase()),
	}
	p.prescriberSystem.Up()

	return p
}

func (f *PrescriberSystemFacade) ImportDatabaseByJSON(file os.File) {
	//f.prescriberSystem.db.CreatePatient(p)
}

func (f *PrescriberSystemFacade) Diagnose(name string, symptoms []Symptom, option Option) {
	patient := f.prescriberSystem.db.FindPatientByName(name)
	if patient == nil {
		println("error: patient not found")
	}

	prescription := f.prescriberSystem.SchedulePrescriber(Demand{
		ID:       0,
		Patient:  *patient,
		Symptoms: symptoms,
	})
	if prescription == nil {
		println("error: diagnosis failed, you need a new doctor")
	}

	patientsCase := Case{
		CaseTime:     time.Now(),
		Symptoms:     symptoms,
		Prescription: *prescription,
	}

	f.prescriberSystem.SavePatientCaseToDB(patientsCase)

	if option == JSON {
		f.prescriberSystem.SavePatientCaseToJSON(patientsCase)
	} else if option == CSV {
		f.prescriberSystem.SavePatientCaseToCSV(patientsCase)
	}
}
