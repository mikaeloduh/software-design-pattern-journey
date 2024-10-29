package test

import (
	"math/rand"
	"testing"
	"time"

	"geneticalgorithm/ga"
)

func TestScheduling(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Initialize population
	populationSize := 50
	initialPopulation := make(ga.Population, populationSize)
	for i := 0; i < populationSize; i++ {
		ind := randomScheduleIndividual()
		initialPopulation[i] = ind
	}

	ga := ga.GeneticAlgorithm{
		Population:    initialPopulation,
		MaxIterations: 100,
		MutationRate:  0.1,
		CrossoverRate: 0.7,
	}

	bestIndividual := ga.Run()
	bestSchedule := bestIndividual.(*ScheduleIndividual)

	// Verify that demands are met
	demandMet := make([]int, len(Products))
	for _, task := range bestSchedule.Tasks {
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
	for _, task := range bestSchedule.Tasks {
		product := Products[task.ProductIndex]
		t.Logf("Produce %d units of Product %s (Production time per unit: %d hours)",
			task.Quantity, product.Name, product.ProductionTime)
	}
}

// Generates a random schedule individual.
func randomScheduleIndividual() *ScheduleIndividual {
	ind := &ScheduleIndividual{}
	totalDemands := make([]int, len(Products))
	for i := 0; i < len(Products); i++ {
		totalDemands[i] = 0
	}

	// Randomly generate tasks until demands are met
	for {
		// Randomly select a product
		productIndex := rand.Intn(len(Products))
		// Calculate remaining demand
		remainingDemand := Products[productIndex].Demand - totalDemands[productIndex]
		if remainingDemand <= 0 {
			continue
		}
		// Randomly decide quantity for this task
		quantity := rand.Intn(remainingDemand) + 1
		ind.Tasks = append(ind.Tasks, Task{ProductIndex: productIndex, Quantity: quantity})
		totalDemands[productIndex] += quantity

		// Check if all demands are met
		demandsMet := true
		for i, product := range Products {
			if totalDemands[i] < product.Demand {
				demandsMet = false
				break
			}
		}
		if demandsMet {
			break
		}
	}

	return ind
}
