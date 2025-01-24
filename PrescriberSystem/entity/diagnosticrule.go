package entity

import "prescribersystem/utils"

func ChainDiagnosticRule(rules ...IDiagnosticRule) IDiagnosticRule {
	if len(rules) == 0 {
		return nil
	}

	for i := 0; i < len(rules)-1; i++ {
		rules[i].Pass(rules[i+1])
	}

	return rules[0]
}

// IDiagnosticRule interface
type IDiagnosticRule interface {
	Handle(Patient, []Symptom) *Prescription
	Pass(IDiagnosticRule)
}

// HerbRule
type EmptyRule struct {
	Next IDiagnosticRule
}

func (r *EmptyRule) Pass(next IDiagnosticRule) {
	r.Next = next
}

func (r *EmptyRule) Handle(patient Patient, symptoms []Symptom) *Prescription {
	if r.Next != nil {
		return r.Next.Handle(patient, symptoms)
	} else {
		return nil
	}
}

// HerbRule
type HerbRule struct {
	Next IDiagnosticRule
}

func (r *HerbRule) Pass(next IDiagnosticRule) {
	r.Next = next
}

func (r *HerbRule) Handle(patient Patient, symptoms []Symptom) *Prescription {
	if utils.ContainsElement(symptoms, Cough) && utils.ContainsElement(symptoms, Headache) {
		return NewHerbPrescription()
	} else if r.Next != nil {
		return r.Next.Handle(patient, symptoms)
	} else {
		return nil
	}
}

// InhibitorRule
type InhibitorRule struct {
	Next IDiagnosticRule
}

func (r *InhibitorRule) Pass(next IDiagnosticRule) {
	r.Next = next
}

func (r *InhibitorRule) Handle(patient Patient, symptoms []Symptom) *Prescription {
	if patient.Age == 18 && utils.ContainsElement(symptoms, Sneeze) {
		return NewInhibitorPrescription()
	} else if r.Next != nil {
		return r.Next.Handle(patient, symptoms)
	} else {
		return nil
	}
}

// ShutUpRule
type ShutUpRule struct {
	Next IDiagnosticRule
}

func (r *ShutUpRule) Pass(next IDiagnosticRule) {
	r.Next = next
}

func (r *ShutUpRule) Handle(patient Patient, symptoms []Symptom) *Prescription {
	if utils.ContainsElement(symptoms, Snore) && utils.BMI(patient.Height, patient.Weight) > 26 {
		return NewShutUpPrescription()
	} else if r.Next != nil {
		return r.Next.Handle(patient, symptoms)
	} else {
		return nil
	}
}
