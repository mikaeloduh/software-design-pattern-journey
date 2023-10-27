package entity

type Prescription struct {
	Name             string
	PotentialDisease string
	Medicines        string
	Usage            string
}

func NewShutUpPrescription() *Prescription {
	return &Prescription{
		Name:             "Shut-up",
		PotentialDisease: "SleepApneaSyndrome",
		Medicines:        "tape",
		Usage:            "tape up your mouth and shut-up",
	}
}

func NewInhibitorPrescription() *Prescription {
	return &Prescription{
		Name:             "Inhibitor",
		PotentialDisease: "Attractive",
		Medicines:        "Side-band",
		Usage:            "Uglify yourself and its all done",
	}
}

func NewHerbPrescription() *Prescription {
	return &Prescription{
		Name:             "Herb No.1",
		PotentialDisease: "COVID-19",
		Medicines:        "Herb No.1",
		Usage:            "boil with 500ml water for 3 minus",
	}
}
