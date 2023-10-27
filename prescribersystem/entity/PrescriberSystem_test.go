package entity

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestPrescriberSystem(t *testing.T) {
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

func TestPrescriberSystem_Run(t *testing.T) {
	db := NewPatientDatabase()
	sys := NewPrescriberSystem(db)

	sys.Up()

	var wg sync.WaitGroup

	t.Run("test new PatientDatabase", func(t *testing.T) {

		DemandPrescriptionAndPrintJSON := func() {
			defer wg.Done()
			p := sys.SchedulePrescriber(Demand{
				ID:       1,
				Patient:  *NewPatient("a0000001", "p1", Male, 87, 159, 100),
				Symptoms: []Symptom{Snore},
			})
			c := Case{
				CaseTime:     time.Now(),
				Symptoms:     nil,
				Prescription: *p,
			}
			sys.SavePatientCaseToDB(c)
			sys.SavePatientCaseToJSON(c)
		}

		wg.Add(1)
		go DemandPrescriptionAndPrintJSON()

	})

	t.Run("tet", func(t *testing.T) {

		DemandPrescriptionAndPrintCSV := func() {
			defer wg.Done()
			p := sys.SchedulePrescriber(Demand{
				ID:       2,
				Patient:  *NewPatient("a0000001", "p1", Male, 87, 159, 100),
				Symptoms: []Symptom{Snore},
			})
			c := Case{
				CaseTime:     time.Now(),
				Symptoms:     nil,
				Prescription: *p,
			}
			sys.SavePatientCaseToDB(c)
			sys.SavePatientCaseToJSON(c)
		}

		wg.Add(1)
		go DemandPrescriptionAndPrintCSV()

	})

	wg.Wait()

	sys.Down()
}
