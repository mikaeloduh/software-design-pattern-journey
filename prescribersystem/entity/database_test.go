package entity

import "testing"

func TestPatient(t *testing.T) {
	t.Parallel()

	t.Run("test new Patient", func(t *testing.T) {

	})

	t.Run("test add Patient into database", func(t *testing.T) {
		db := NewPatientDatabase()
		p := Patient{Id: "a000000001"}

		db.CreatePatient(p)
	})
}
