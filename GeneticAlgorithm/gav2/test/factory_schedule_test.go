package test

import (
	"math/rand"
	"testing"

	ga "geneticalgorithm/gav2"
)

func TestScheduling(t *testing.T) {
	rand.Seed(42) // For reproducibility

	// Initialize population
	populationSize := 50
	initialPopulation := make([]ga.Individual, populationSize)
	for i := 0; i < populationSize; i++ {
		genome := randomScheduleGenome()
		ind := &ScheduleIndividual{genome: genome}
		initialPopulation[i] = ind
	}

	ga := ga.GeneticAlgorithm{
		Population:        initialPopulation,
		MaxIterations:     100,
		MutationRate:      0.1,
		CrossoverRate:     0.7,
		SelectionOperator: ga.RankSelection,
		CrossoverOperator: ga.SinglePointCrossover,
		MutationOperator:  ga.InversionMutation,
	}

	bestIndividual := ga.Run()
	bestSchedule := bestIndividual.(*ScheduleIndividual)

	// Verify that demands are met
	demandMet := make([]int, len(Products))
	for _, task := range bestSchedule.genome.Tasks {
		demandMet[task.ProductIndex] += task.Quantity
	}
	for i, product := range Products {
		if demandMet[i] < product.Demand {
			t.Errorf("Demand not met for product %s: %d < %d", product.Name, demandMet[i], product.Demand)
		}
	}

	totalTime := bestSchedule.CalculateTotalTime()

	t.Logf("Best schedule total time: %d hours", totalTime)
	t.Log("Production tasks:")
	for _, task := range bestSchedule.genome.Tasks {
		product := Products[task.ProductIndex]
		t.Logf("Produce %d units of Product %s (Production time per unit: %d hours)",
			task.Quantity, product.Name, product.ProductionTime)
	}
}

// Generates a random schedule genome.
func randomScheduleGenome() *ScheduleGenome {
	genome := &ScheduleGenome{Tasks: []Task{}}
	demandMet := make([]int, len(Products))

	// Keep generating tasks until demands are met
	for {
		task := genome.RandomGene().(Task)
		productIndex := task.ProductIndex
		demandMet[productIndex] += task.Quantity

		// Add the task to the genome
		genome.Tasks = append(genome.Tasks, task)

		// Check if all demands are met
		allDemandsMet := true
		for i, product := range Products {
			if demandMet[i] < product.Demand {
				allDemandsMet = false
				break
			}
		}

		if allDemandsMet {
			break
		}
	}

	// Shuffle the tasks
	rand.Shuffle(len(genome.Tasks), func(i, j int) {
		genome.Tasks[i], genome.Tasks[j] = genome.Tasks[j], genome.Tasks[i]
	})

	return genome
}
