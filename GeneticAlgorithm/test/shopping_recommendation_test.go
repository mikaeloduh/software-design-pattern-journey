package test

import (
	"math/rand"
	"testing"
	"time"

	"geneticalgorithm/ga"
)

// Merchandise represents a product with price, weight, and category.
type Merchandise struct {
	Price    int
	Weight   int
	Category string
}

// Define available products.
var Merchandises = []Merchandise{
	{Price: 100, Weight: 2, Category: "A"}, // Merchandise 1
	{Price: 200, Weight: 3, Category: "A"}, // Merchandise 2
	{Price: 150, Weight: 5, Category: "B"}, // Merchandise 3
	{Price: 300, Weight: 4, Category: "B"}, // Merchandise 4
	{Price: 180, Weight: 6, Category: "C"}, // Merchandise 5
	{Price: 250, Weight: 7, Category: "C"}, // Merchandise 6
}

// Customer's preferences for each category.
var Preferences = map[string]float64{
	"A": 0.8,
	"B": 0.6,
	"C": 0.2,
}

// Constraints for the recommendation.
const (
	Budget   = 700
	Capacity = 15
)

// RecommendationIndividual represents a recommendation list.
type RecommendationIndividual struct {
	Quantities []int // quantities of each product
}

// Fitness calculates the individual's fitness based on preferences and constraints.
func (ind *RecommendationIndividual) Fitness() float64 {
	totalPrice := 0
	totalWeight := 0
	totalPreference := 0.0
	for i, qty := range ind.Quantities {
		totalPrice += qty * Merchandises[i].Price
		totalWeight += qty * Merchandises[i].Weight
		totalPreference += float64(qty) * Preferences[Merchandises[i].Category]
	}
	if totalPrice > Budget || totalWeight > Capacity {
		return 0
	}
	return totalPreference
}

// Crossover performs one-point crossover with another individual.
func (ind *RecommendationIndividual) Crossover(other ga.Individual) (ga.Individual, ga.Individual) {
	otherInd := other.(*RecommendationIndividual)
	point := rand.Intn(len(ind.Quantities))
	child1 := &RecommendationIndividual{Quantities: make([]int, len(ind.Quantities))}
	child2 := &RecommendationIndividual{Quantities: make([]int, len(ind.Quantities))}
	for i := 0; i < len(ind.Quantities); i++ {
		if i < point {
			child1.Quantities[i] = ind.Quantities[i]
			child2.Quantities[i] = otherInd.Quantities[i]
		} else {
			child1.Quantities[i] = otherInd.Quantities[i]
			child2.Quantities[i] = ind.Quantities[i]
		}
	}
	return child1, child2
}

// Mutate randomly alters the individual's quantities.
func (ind *RecommendationIndividual) Mutate() {
	index := rand.Intn(len(ind.Quantities))
	// Mutate the quantity at the selected index
	// Randomly increase or decrease the quantity
	change := rand.Intn(3) - 1 // -1, 0, or +1
	newQty := ind.Quantities[index] + change
	if newQty < 0 {
		newQty = 0
	}
	ind.Quantities[index] = newQty
}

// Clone creates a deep copy of the individual.
func (ind *RecommendationIndividual) Clone() ga.Individual {
	clone := &RecommendationIndividual{Quantities: make([]int, len(ind.Quantities))}
	copy(clone.Quantities, ind.Quantities)
	return clone
}

func TestRecommendation(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// 初始化種群
	populationSize := 50
	initialPopulation := make(ga.Population, populationSize)
	for i := 0; i < populationSize; i++ {
		ind := &RecommendationIndividual{Quantities: make([]int, len(Merchandises))}
		for j := 0; j < len(Merchandises); j++ {
			ind.Quantities[j] = rand.Intn(3) // 隨機數量 0, 1 或 2
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

	// 確認最佳推薦符合限制條件
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
