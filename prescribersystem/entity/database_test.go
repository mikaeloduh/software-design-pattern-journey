package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPatientDatabase(t *testing.T) {
	t.Parallel()

	t.Run("test add Patient into database and should be found in it", func(t *testing.T) {
		db := NewPatientDatabase()
		p := NewPatient("a0000001", "p1", Male, 87, 159, 100)

		db.CreatePatient(p)

		assert.Equal(t, db.FindPatientById("a0000001"), p)
	})
}
