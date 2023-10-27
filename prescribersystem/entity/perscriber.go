package entity

import "prescribersystem/utils"

type IDiagnosticRule interface {
	handle(Patient, []Symptom)
}

// Prescriber
type Prescriber struct {
}

func (p *Prescriber) Diagnose(patient Patient, symptoms []Symptom) *Prescription {

	if utils.ContainsElement(symptoms, Cough) && utils.ContainsElement(symptoms, Headache) {
		return NewHerbPrescription()
	} else if patient.Age == 18 && utils.ContainsElement(symptoms, Sneeze) {
		return NewInhibitorPrescription()
	} else if utils.ContainsElement(symptoms, Snore) && utils.BMI(patient.Height, patient.Weight) > 26 {
		return NewShutUpPrescription()
	}

	return nil
}
