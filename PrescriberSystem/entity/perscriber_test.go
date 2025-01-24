package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrescriber(t *testing.T) {
	t.Parallel()

	t.Run("test Prescriber Diagnose", func(t *testing.T) {
		r := FakeNewPrescriber()
		p := r.Diagnose(Patient{}, []Symptom{})

		assert.IsType(t, &Prescription{}, p)
	})

	t.Run("test Prescriber Diagnose COVID-19", func(t *testing.T) {
		r := FakeNewPrescriber()

		p := r.Diagnose(Patient{}, []Symptom{Headache, Cough})

		assert.Equal(t, NewHerbPrescription(), p)
	})

	t.Run("test Prescriber Diagnose Attractive", func(t *testing.T) {
		r := FakeNewPrescriber()

		p := r.Diagnose(Patient{Age: 18}, []Symptom{Sneeze})

		assert.Equal(t, NewInhibitorPrescription(), p)
	})

	t.Run("test Prescriber Diagnose SleepApneaSyndrome", func(t *testing.T) {
		r := FakeNewPrescriber()

		p := r.Diagnose(Patient{Weight: 100, Height: 150}, []Symptom{Snore})

		assert.Equal(t, NewShutUpPrescription(), p)
	})

}
func FakeNewPrescriber() *Prescriber {
	return &Prescriber{rule: ChainDiagnosticRule(&HerbRule{}, &InhibitorRule{}, &ShutUpRule{})}
}
