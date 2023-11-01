package entity

import (
	"encoding/json"
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

func NewPrescriberSystemFacade(configFile *os.File) *PrescriberSystemFacade {
	p := &PrescriberSystemFacade{
		prescriberSystem: NewPrescriberSystem(
			NewPatientDatabase(),
			fileToConfig(configFile),
		),
	}
	p.prescriberSystem.Up()

	return p
}

func fileToConfig(file *os.File) Config {
	// TODO: to be implement
	return Config{}
}

func (f *PrescriberSystemFacade) setSupportRules(file *os.File) {
}

func (f *PrescriberSystemFacade) ImportDatabaseByJSON(file *os.File) {
	var patients []Patient
	decoder := json.NewDecoder(file)
	decoder.Decode(&patients)
	for _, p := range patients {
		f.prescriberSystem.db.CreatePatient(&p)
	}
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
