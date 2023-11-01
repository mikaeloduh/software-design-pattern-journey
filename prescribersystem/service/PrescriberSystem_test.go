package service

import (
	"github.com/stretchr/testify/assert"
	"prescribersystem/entity"
	"sync"
	"testing"
	"time"
)

func TestPrescriberSystem(t *testing.T) {
	t.Parallel()

	config := Config{
		COVID19:            true,
		Attractive:         true,
		SleepApneaSyndrome: true,
	}

	t.Run("test new PatientDatabase", func(t *testing.T) {
		db := entity.NewPatientDatabase()

		assert.IsType(t, &entity.PatientDatabase{}, db)
	})

	t.Run("test new PrescriberSystem", func(t *testing.T) {
		db := entity.NewPatientDatabase()
		sys := NewPrescriberSystem(db, config)

		assert.IsType(t, &PrescriberSystem{}, sys)
	})
}

func TestPrescriberSystem_Run(t *testing.T) {
	config := Config{
		COVID19:            true,
		Attractive:         true,
		SleepApneaSyndrome: true,
	}
	db := entity.NewPatientDatabase()
	sys := NewPrescriberSystem(db, config)

	sys.Up()

	var wg sync.WaitGroup

	t.Run("test new PatientDatabase and save as JSON", func(t *testing.T) {

		DemandPrescriptionAndPrintJSON := func() {
			defer wg.Done()
			p := sys.SchedulePrescriber(Demand{
				ID:       1,
				Patient:  *entity.NewPatient("a0000001", "p1", entity.Male, 87, 159, 100),
				Symptoms: []entity.Symptom{entity.Snore},
			})
			c := entity.Case{
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

	t.Run("test new PatientDatabase and save as CSV", func(t *testing.T) {

		DemandPrescriptionAndPrintCSV := func() {
			defer wg.Done()
			p := sys.SchedulePrescriber(Demand{
				ID:       2,
				Patient:  *entity.NewPatient("a0000001", "p1", entity.Male, 87, 159, 100),
				Symptoms: []entity.Symptom{entity.Snore},
			})
			c := entity.Case{
				CaseTime:     time.Now(),
				Symptoms:     nil,
				Prescription: *p,
			}
			sys.SavePatientCaseToDB(c)
			sys.SavePatientCaseToCSV(c)
		}

		wg.Add(1)
		go DemandPrescriptionAndPrintCSV()

	})

	wg.Wait()

	sys.Down()
}
