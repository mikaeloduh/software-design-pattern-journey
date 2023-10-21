package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrescriber(t *testing.T) {
	t.Parallel()

	t.Run("test new PatientDatabase", func(t *testing.T) {
		db := NewPatientDatabase()

		assert.IsType(t, &PatientDatabase{}, db)
	})

	t.Run("test new PrescriberSystem", func(t *testing.T) {
		db := NewPatientDatabase()
		sys := NewPrescriberSystem(db)

		assert.IsType(t, &PrescriberSystem{}, sys)
	})

}
