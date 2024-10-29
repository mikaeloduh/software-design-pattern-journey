package test

import (
	"math/rand"

	ga "geneticalgorithm/gav2"
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

const (
	Budget   = 700
	Capacity = 15
)

// RecommendationGenome represents the genome for the recommendation problem.
type RecommendationGenome struct {
	Quantities []int
}

// Implement Genome interface
func (g *RecommendationGenome) Length() int {
	return len(g.Quantities)
}

func (g *RecommendationGenome) GetGene(index int) interface{} {
	return g.Quantities[index]
}

func (g *RecommendationGenome) SetGene(index int, value interface{}) {
	g.Quantities[index] = value.(int)
}

func (g *RecommendationGenome) SwapGenes(i, j int) {
	g.Quantities[i], g.Quantities[j] = g.Quantities[j], g.Quantities[i]
}

func (g *RecommendationGenome) RandomGene() interface{} {
	// Return a random quantity (0, 1, or 2)
	return rand.Intn(3)
}

func (g *RecommendationGenome) Clone() ga.Genome {
	copied := make([]int, len(g.Quantities))
	copy(copied, g.Quantities)
	return &RecommendationGenome{Quantities: copied}
}

// RecommendationIndividual represents an individual in the recommendation problem.
type RecommendationIndividual struct {
	genome *RecommendationGenome
}

// Implement Individual interface
func (ind *RecommendationIndividual) Genome() ga.Genome {
	return ind.genome
}

func (ind *RecommendationIndividual) SetGenome(genome ga.Genome) {
	ind.genome = genome.(*RecommendationGenome)
}

func (ind *RecommendationIndividual) Fitness() float64 {
	totalPrice := 0
	totalWeight := 0
	totalPreference := 0.0
	for i, qty := range ind.genome.Quantities {
		totalPrice += qty * Merchandises[i].Price
		totalWeight += qty * Merchandises[i].Weight
		totalPreference += float64(qty) * Preferences[Merchandises[i].Category]
	}
	if totalPrice > Budget || totalWeight > Capacity {
		return 0 // Over budget, invalid solution
	}
	return totalPreference
}

func (ind *RecommendationIndividual) Clone() ga.Individual {
	return &RecommendationIndividual{
		genome: ind.genome.Clone().(*RecommendationGenome),
	}
}
