package entity

// Prescriber
type Prescriber struct {
	handler IDiagnosticRule
}

func NewPrescriber(handler IDiagnosticRule) *Prescriber {
	return &Prescriber{handler: handler}
}

func NewDefaultPrescriber() *Prescriber {
	handler := ChainDiagnosticRule(&HerbRule{}, &InhibitorRule{}, &ShutUpRule{})

	return &Prescriber{handler: handler}
}

func (p *Prescriber) Diagnose(patient Patient, symptoms []Symptom) *Prescription {
	//time.Sleep(3 * time.Second) // Simulate task processing

	return p.handler.Handle(patient, symptoms)
}
