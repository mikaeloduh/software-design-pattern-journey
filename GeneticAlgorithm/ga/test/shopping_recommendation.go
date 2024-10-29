package test

import (
	"math/rand"

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
