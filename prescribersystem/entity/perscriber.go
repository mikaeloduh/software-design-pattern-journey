package entity

import "time"

// Prescriber
type Prescriber struct {
	rule IDiagnosticRule
}

func NewPrescriber(rule IDiagnosticRule) *Prescriber {
	return &Prescriber{rule: rule}
}

func NewDefaultPrescriber() *Prescriber {
	handler := ChainDiagnosticRule(&HerbRule{}, &InhibitorRule{}, &ShutUpRule{})

	return &Prescriber{rule: handler}
}

func (p *Prescriber) Diagnose(patient Patient, symptoms []Symptom) *Prescription {
	time.Sleep(3 * time.Second) // Simulate task processing

	return p.rule.Handle(patient, symptoms)
}
