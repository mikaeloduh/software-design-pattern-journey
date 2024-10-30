package test

import (
	"math/rand"

	ga "geneticalgorithm/gav2"
)

// Product represents a product with production time and demand.
type Product struct {
	Name           string
	ProductionTime int // in hours per unit
	Demand         int // units to produce
}

// Task represents a production task.
type Task struct {
	ProductIndex int // index of the product in the Products slice
	Quantity     int // quantity to produce in this task
}

// Define the products and available resources.
var (
	Products []Product
	Machines int
	Workers  int
)

// ScheduleGenome represents the genome for the scheduling problem.
type ScheduleGenome struct {
	Tasks []Task
}

// Implement Genome interface
func (g *ScheduleGenome) Length() int {
	return len(g.Tasks)
}

func (g *ScheduleGenome) GetGene(index int) interface{} {
	return g.Tasks[index]
}

func (g *ScheduleGenome) SetGene(index int, value interface{}) {
	g.Tasks[index] = value.(Task)
}

func (g *ScheduleGenome) SwapGenes(i, j int) {
	g.Tasks[i], g.Tasks[j] = g.Tasks[j], g.Tasks[i]
}

func (g *ScheduleGenome) RandomGene() interface{} {
	// Generate a random Task
	productIndex := rand.Intn(len(Products))
	quantity := 1 // Fixed quantity for simplicity

	return Task{
		ProductIndex: productIndex,
		Quantity:     quantity,
	}
}

func (g *ScheduleGenome) Clone() ga.Genome {
	copied := make([]Task, len(g.Tasks))
	copy(copied, g.Tasks)
	return &ScheduleGenome{Tasks: copied}
}

// ScheduleIndividual represents an individual in the scheduling problem.
type ScheduleIndividual struct {
	genome *ScheduleGenome
}

// Implement Individual interface
func (ind *ScheduleIndividual) Genome() ga.Genome {
	return ind.genome
}

func (ind *ScheduleIndividual) SetGenome(genome ga.Genome) {
	ind.genome = genome.(*ScheduleGenome)
}

func (ind *ScheduleIndividual) Fitness() float64 {
	// Calculate total production for each product
	demandMet := make([]int, len(Products))
	for _, task := range ind.genome.Tasks {
		demandMet[task.ProductIndex] += task.Quantity
	}
	// Check if demands are met
	for i, product := range Products {
		if demandMet[i] < product.Demand {
			return 0 // Demands not met, fitness is zero
		}
	}

	// Simulate the production to calculate total time
	totalTime := ind.CalculateTotalTime()
	if totalTime == 0 {
		return 0 // Invalid schedule
	}
	// Fitness is the inverse of total time
	return 1.0 / float64(totalTime)
}

// CalculateTotalTime simulates the production schedule and returns total time.
func (ind *ScheduleIndividual) CalculateTotalTime() int {
	type Resource struct {
		AvailableAt int // time when the resource becomes available
	}

	// Initialize resources
	machines := make([]Resource, Machines)
	workers := make([]Resource, Workers)

	currentTime := 0

	for _, task := range ind.genome.Tasks {
		product := Products[task.ProductIndex]
		unitsLeft := task.Quantity

		for unitsLeft > 0 {
			// Find the earliest available machine and worker
			machineAvailableAt := machines[0].AvailableAt
			workerAvailableAt := workers[0].AvailableAt
			machineIndex := 0
			workerIndex := 0

			for i := 1; i < Machines; i++ {
				if machines[i].AvailableAt < machineAvailableAt {
					machineAvailableAt = machines[i].AvailableAt
					machineIndex = i
				}
			}
			for i := 1; i < Workers; i++ {
				if workers[i].AvailableAt < workerAvailableAt {
					workerAvailableAt = workers[i].AvailableAt
					workerIndex = i
				}
			}

			// Start time for this unit is the max of machine and worker availability
			startTime := max(machineAvailableAt, workerAvailableAt)
			// Production time for one unit
			productionTime := product.ProductionTime
			// Update resource availability
			machines[machineIndex].AvailableAt = startTime + productionTime
			workers[workerIndex].AvailableAt = startTime + productionTime

			// Update current time
			if startTime+productionTime > currentTime {
				currentTime = startTime + productionTime
			}

			unitsLeft--
		}
	}

	return currentTime
}

func (ind *ScheduleIndividual) Clone() ga.Individual {
	return &ScheduleIndividual{
		genome: ind.genome.Clone().(*ScheduleGenome),
	}
}
