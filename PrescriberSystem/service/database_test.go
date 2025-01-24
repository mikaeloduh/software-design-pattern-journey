package service

import (
	"github.com/stretchr/testify/assert"
	"prescribersystem/entity"
	"testing"
)

func TestPatientDatabase(t *testing.T) {
	t.Parallel()

	t.Run("test add Patient into database and should be found in it", func(t *testing.T) {
		db := NewPatientDatabase()
		p := entity.NewPatient("a0000001", "p1", entity.Male, 87, 159, 100)

		db.CreatePatient(p)

		assert.Equal(t, db.FindPatientById("a0000001"), p)
	})
}
