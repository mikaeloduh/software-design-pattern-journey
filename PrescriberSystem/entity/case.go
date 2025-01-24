package entity

import "time"

type Case struct {
	CaseTime     time.Time
	Symptoms     []Symptom
	Prescription Prescription
}
