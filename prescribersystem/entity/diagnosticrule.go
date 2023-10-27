package entity

import "prescribersystem/utils"

type IDiagnosticRule interface {
	Handle(Patient, []Symptom) *Prescription
}

type HerbRule struct {
	Next IDiagnosticRule
}

func (r HerbRule) Handle(patient Patient, symptoms []Symptom) *Prescription {
	if utils.ContainsElement(symptoms, Cough) && utils.ContainsElement(symptoms, Headache) {
		return NewHerbPrescription()
	} else if r.Next != nil {
		return r.Next.Handle(patient, symptoms)
	} else {
		return nil
	}
}

type InhibitorRule struct {
	Next IDiagnosticRule
}

func (r InhibitorRule) Handle(patient Patient, symptoms []Symptom) *Prescription {
	if patient.Age == 18 && utils.ContainsElement(symptoms, Sneeze) {
		return NewInhibitorPrescription()
	} else if r.Next != nil {
		return r.Next.Handle(patient, symptoms)
	} else {
		return nil
	}
}

type ShutUpRule struct {
	Next IDiagnosticRule
}

func (r ShutUpRule) Handle(patient Patient, symptoms []Symptom) *Prescription {
	if utils.ContainsElement(symptoms, Snore) && utils.BMI(patient.Height, patient.Weight) > 26 {
		return NewShutUpPrescription()
	} else if r.Next != nil {
		return r.Next.Handle(patient, symptoms)
	} else {
		return nil
	}
}
