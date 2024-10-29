package test

import (
	"math/rand"
	"testing"
	"time"

	"geneticalgorithm/ga"
)

func TestRecommendation(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// initialize population
	populationSize := 50
	initialPopulation := make(ga.Population, populationSize)
	for i := 0; i < populationSize; i++ {
		ind := &RecommendationIndividual{Quantities: make([]int, len(Merchandises))}
		for j := 0; j < len(Merchandises); j++ {
			ind.Quantities[j] = rand.Intn(3) // random from 0, 1 or 2
		}
		initialPopulation[i] = ind
	}

	ga := ga.GeneticAlgorithm{
		Population:    initialPopulation,
		MaxIterations: 100,
		MutationRate:  0.1,
		CrossoverRate: 0.7,
	}

	bestIndividual := ga.Run()
	bestRec := bestIndividual.(*RecommendationIndividual)

	// Verify that constraints are not violated
	totalPrice := 0
	totalWeight := 0
	totalPreference := 0.0
	for i, qty := range bestRec.Quantities {
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
	for i, qty := range bestRec.Quantities {
		if qty > 0 {
			t.Logf("Merchandise %d x %d (Price %d, Weight %dkg, Category %s, Preference %.2f)",
				i+1, qty, Merchandises[i].Price, Merchandises[i].Weight, Merchandises[i].Category, Preferences[Merchandises[i].Category])
		}
	}
	t.Logf("Total Price: %d", totalPrice)
	t.Logf("Total Weight: %dkg", totalWeight)
	t.Logf("Total Preference: %.2f", totalPreference)
}
