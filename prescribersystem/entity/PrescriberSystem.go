package entity

import "fmt"

// PrescriberSystem
type PrescriberSystem struct {
	db     *PatientDatabase
	worker *PrescriberWorker
}

func NewPrescriberSystem(db *PatientDatabase) *PrescriberSystem {
	return &PrescriberSystem{
		db:     db,
		worker: NewPrescriberWorker(1),
	}
}

func (s *PrescriberSystem) Run() {
	go s.worker.Start()
}

func (s *PrescriberSystem) Down() {
	s.worker.Stop()
}

func (s *PrescriberSystem) SchedulePrescriber(d Demand) Prescription {
	fmt.Printf("Scheduling damand #%d...\n", d.ID)
	w := s.getWorker()
	w.reqCh <- d

	return <-w.resCh
}

func (s *PrescriberSystem) getWorker() *PrescriberWorker {
	// put scheduling algorithm in here
	return s.worker
}

func (s *PrescriberSystem) SavePrescriptionToDB(p Prescription) {
	fmt.Println("saving to DB")
}

func (s *PrescriberSystem) SavePrescriptionToJSON(p Prescription) {
	fmt.Println("saving to JSON")
}

func (s *PrescriberSystem) SavePrescriptionToCSV(p Prescription) {
	fmt.Println("saving to CSV")
}

// Demand
type Demand struct {
	ID       int
	Patient  Patient
	Symptoms []Symptom
}

// PrescriberWorker
type PrescriberWorker struct {
	ID         int
	prescriber *Prescriber
	reqCh      chan Demand
	resCh      chan Prescription
	doneCh     chan struct{}
}

func NewPrescriberWorker(ID int) *PrescriberWorker {
	return &PrescriberWorker{
		ID:         ID,
		prescriber: NewDefaultPrescriber(),
		reqCh:      make(chan Demand),
		resCh:      make(chan Prescription),
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
			w.resCh <- *p
		case <-w.doneCh:
			fmt.Printf("Worker %d stopped.\n", w.ID)
			return
		}
	}
}

func (w *PrescriberWorker) Stop() {
	close(w.doneCh)
}
