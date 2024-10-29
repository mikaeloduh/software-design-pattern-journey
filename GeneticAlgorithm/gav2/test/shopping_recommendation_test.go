package test

import (
	"math/rand"
	"testing"
	"time"

	ga "geneticalgorithm/gav2"
)

func TestRecommendation(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Initialize population
	populationSize := 50
	initialPopulation := make([]ga.Individual, populationSize)
	for i := 0; i < populationSize; i++ {
		genome := &RecommendationGenome{Quantities: make([]int, len(Merchandises))}
		for j := 0; j < len(Merchandises); j++ {
			genome.Quantities[j] = genome.RandomGene().(int) // Random quantities 0, 1, or 2
		}
		ind := &RecommendationIndividual{genome: genome}
		initialPopulation[i] = ind
	}

	ga := ga.GeneticAlgorithm{
		Population:        initialPopulation,
		MaxIterations:     100,
		MutationRate:      0.1,
		CrossoverRate:     0.7,
		SelectionOperator: ga.NewTournamentSelection(3),
		CrossoverOperator: ga.SinglePointCrossover,
		MutationOperator:  ga.RandomReplacementMutation,
	}

	bestIndividual := ga.Run()
	bestRec := bestIndividual.(*RecommendationIndividual)

	// Verify that constraints are not violated
	totalPrice := 0
	totalWeight := 0
	totalPreference := 0.0
	for i, qty := range bestRec.genome.Quantities {
		totalPrice += qty * Merchandises[i].Price
		totalWeight += qty * Merchandises[i].Weight
		totalPreference += float64(qty) * Preferences[Merchandises[i].Category]
	}
	if totalPrice > Budget {
		t.Errorf("Total price exceeds budget: %d > %d", totalPrice, Budget)
	}
	if totalWeight > Capacity {
		t.Errorf("Total weight exceeds capacity: %d > %d", totalWeight, Capacity)
	}

	t.Log("Best recommendation:")
	for i, qty := range bestRec.genome.Quantities {
		if qty > 0 {
			t.Logf("Merchandise %d x %d (Price %d, Weight %dkg, Category %s, Preference %.2f)",
				i+1, qty, Merchandises[i].Price, Merchandises[i].Weight, Merchandises[i].Category, Preferences[Merchandises[i].Category])
		}
	}
	t.Logf("Total Price: %d", totalPrice)
	t.Logf("Total Weight: %dkg", totalWeight)
	t.Logf("Total Preference: %.2f", totalPreference)

}
