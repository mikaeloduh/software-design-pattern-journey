package service

import (
	"fmt"
	"prescribersystem/entity"
)

// Demand
type Demand struct {
	ID       int
	Patient  entity.Patient
	Symptoms []entity.Symptom
}

// PrescriberSystem
type PrescriberSystem struct {
	db     *PatientDatabase
	worker *PrescriberWorker
	config Config
}

func NewPrescriberSystem(db *PatientDatabase, config Config) *PrescriberSystem {
	var rules []entity.IDiagnosticRule
	if config.IsEnabled(COVID19) {
		rules = append(rules, &entity.HerbRule{})
	}
	if config.IsEnabled(Attractive) {
		rules = append(rules, &entity.InhibitorRule{})
	}
	if config.IsEnabled(SleepApneaSyndrome) {
		rules = append(rules, &entity.ShutUpRule{})
	}

	return &PrescriberSystem{
		db:     db,
		worker: NewPrescriberWorker(1, entity.ChainDiagnosticRule(rules...)),
		config: config,
	}
}

func (s *PrescriberSystem) Up() {
	go s.worker.Start()
}

func (s *PrescriberSystem) Down() {
	s.worker.Stop()
}

func (s *PrescriberSystem) SchedulePrescriber(d Demand) *entity.Prescription {
	fmt.Printf("Scheduling damand #%d...\n", d.ID)
	w := s.getWorker()
	w.reqCh <- d

	return <-w.resCh
}

func (s *PrescriberSystem) getWorker() *PrescriberWorker {
	// put scheduling algorithm in here
	return s.worker
}

func (s *PrescriberSystem) SavePatientCaseToDB(c entity.Case) {
	fmt.Println("saving to DB")
}

func (s *PrescriberSystem) SavePatientCaseToJSON(c entity.Case) {
	fmt.Println("saving to JSON")
}

func (s *PrescriberSystem) SavePatientCaseToCSV(c entity.Case) {
	fmt.Println("saving to CSV")
}

// PrescriberWorker
type PrescriberWorker struct {
	ID         int
	prescriber *entity.Prescriber
	reqCh      chan Demand
	resCh      chan *entity.Prescription
	doneCh     chan struct{}
}

func NewPrescriberWorker(ID int, rule entity.IDiagnosticRule) *PrescriberWorker {
	return &PrescriberWorker{
		ID:         ID,
		prescriber: entity.NewPrescriber(rule),
		reqCh:      make(chan Demand),
		resCh:      make(chan *entity.Prescription),
		doneCh:     make(chan struct{}),
	}
}

func (w *PrescriberWorker) Start() {
	fmt.Printf("Worker %d is ready.\n", w.ID)
	for {
		select {
		case demand := <-w.reqCh:
			fmt.Printf("Worker %d is processing Demand %d.\n", w.ID, demand.ID)
			p := w.prescriber.Diagnose(demand.Patient, demand.Symptoms)
			fmt.Printf("Worker %d finished Demand %d.\n", w.ID, demand.ID)
			w.resCh <- p
		case <-w.doneCh:
			fmt.Printf("Worker %d stopped.\n", w.ID)
			return
		}
	}
}

func (w *PrescriberWorker) Stop() {
	close(w.doneCh)
}
