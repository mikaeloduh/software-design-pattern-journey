package entity

import "fmt"

// PrescriberSystem
type PrescriberSystem struct {
	db     *PatientDatabase
	worker *PrescriberWorker
	config Config
}

func NewPrescriberSystem(db *PatientDatabase, config Config) *PrescriberSystem {
	return &PrescriberSystem{
		db:     db,
		worker: NewPrescriberWorker(1, config),
		config: config,
	}
}

func (s *PrescriberSystem) Up() {
	go s.worker.Start()
}

func (s *PrescriberSystem) Down() {
	s.worker.Stop()
}

func (s *PrescriberSystem) SchedulePrescriber(d Demand) *Prescription {
	fmt.Printf("Scheduling damand #%d...\n", d.ID)
	w := s.getWorker()
	w.reqCh <- d

	return <-w.resCh
}

func (s *PrescriberSystem) getWorker() *PrescriberWorker {
	// put scheduling algorithm in here
	return s.worker
}

func (s *PrescriberSystem) SavePatientCaseToDB(c Case) {
	fmt.Println("saving to DB")
}

func (s *PrescriberSystem) SavePatientCaseToJSON(c Case) {
	fmt.Println("saving to JSON")
}

func (s *PrescriberSystem) SavePatientCaseToCSV(c Case) {
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
	resCh      chan *Prescription
	doneCh     chan struct{}
}

func NewPrescriberWorker(ID int, config Config) *PrescriberWorker {
	var rules []IDiagnosticRule
	if config.IsEnabled(COVID19) {
		rules = append(rules, &HerbRule{})
	}
	if config.IsEnabled(Attractive) {
		rules = append(rules, &InhibitorRule{})
	}
	if config.IsEnabled(SleepApneaSyndrome) {
		rules = append(rules, &ShutUpRule{})
	}

	handler := ChainDiagnosticRule(rules...)

	return &PrescriberWorker{
		ID:         ID,
		prescriber: NewPrescriber(handler),
		reqCh:      make(chan Demand),
		resCh:      make(chan *Prescription),
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
