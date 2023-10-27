package entity

// Prescriber
type Prescriber struct {
	handler IDiagnosticRule
}

func NewPrescriber(handler IDiagnosticRule) *Prescriber {
	return &Prescriber{handler: handler}
}

func (p *Prescriber) Diagnose(patient Patient, symptoms []Symptom) *Prescription {

	return p.handler.Handle(patient, symptoms)
}
