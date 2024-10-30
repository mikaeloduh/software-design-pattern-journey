package test

import (
	"math/rand"
	"testing"
	"time"

	ga "geneticalgorithm/gav2"
)

func TestRecommendation(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Define available Merchandises.
	Merchandises = []Merchandise{
		{Price: 100, Weight: 2, Category: "A"}, // Merchandise 1
		{Price: 200, Weight: 3, Category: "A"}, // Merchandise 2
		{Price: 150, Weight: 5, Category: "B"}, // Merchandise 3
		{Price: 300, Weight: 4, Category: "B"}, // Merchandise 4
		{Price: 180, Weight: 6, Category: "C"}, // Merchandise 5
		{Price: 250, Weight: 7, Category: "C"}, // Merchandise 6
	}
	// Define customer preference budget and capacity constrain
	Preferences = map[string]float64{
		"A": 0.8,
		"B": 0.6,
		"C": 0.2,
	}
	Budget = 700.0
	Capacity = 15

	// Initialize population
	populationSize := 50
	initialPopulation := make([]ga.Individual, populationSize)
	for i := 0; i < populationSize; i++ {
		genome := randomRecommendationGenome()
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
	totalPrice := 0.0
	totalWeight := 0
	for _, merchandise := range bestRec.genome.RecommendationList {
		totalPrice += merchandise.Price
		totalWeight += merchandise.Weight
	}
	if totalPrice > Budget {
		t.Errorf("Budget exceeded: $%.2f > $%.2f", totalPrice, Budget)
	}
	if totalWeight > Capacity {
		t.Errorf("Capacity exceeded: %d > %d", totalWeight, Capacity)
	}

	totalUtility := bestRec.Fitness()
	t.Logf("Best total utility: %.2f", totalUtility)
	t.Logf("Total price: $%.2f", totalPrice)
	t.Logf("Total weight: %d", totalWeight)
	t.Log("Recommended merchandises:")
	for _, merchandise := range bestRec.genome.RecommendationList {
		t.Logf("- Category %s: Price $%.2f, Weight %d", merchandise.Category, merchandise.Price, merchandise.Weight)
	}
}

func randomRecommendationGenome() *RecommendationGenome {
	genome := &RecommendationGenome{RecommendationList: []Merchandise{}}
	totalPrice := 0.0
	totalWeight := 0

	// Randomly add merchandises until reaching the budget or capacity
	for {
		merchandise := Merchandises[rand.Intn(len(Merchandises))]
		if totalPrice+merchandise.Price > Budget || totalWeight+merchandise.Weight > Capacity {
			break // Cannot add more without exceeding constraints
		}
		genome.RecommendationList = append(genome.RecommendationList, merchandise)
		totalPrice += merchandise.Price
		totalWeight += merchandise.Weight
	}

	return genome
}
