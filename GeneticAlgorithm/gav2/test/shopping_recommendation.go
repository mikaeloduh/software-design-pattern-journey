package test

import (
	"math/rand"

	ga "geneticalgorithm/gav2"
)

// Merchandise represents a product with price, weight, and category.
type Merchandise struct {
	Price    float64
	Weight   int
	Category string
}

var (
	Merchandises []Merchandise
	Preferences  map[string]float64
	Budget       float64
	Capacity     int
)

// RecommendationGenome represents the genome for the recommendation problem.
type RecommendationGenome struct {
	RecommendationList []Merchandise
}

// Implement Genome interface
func (g *RecommendationGenome) Length() int {
	return len(g.RecommendationList)
}

func (g *RecommendationGenome) GetGene(index int) interface{} {
	return g.RecommendationList[index]
}

func (g *RecommendationGenome) SetGene(index int, value interface{}) {
	g.RecommendationList[index] = value.(Merchandise)
}

func (g *RecommendationGenome) SwapGenes(i, j int) {
	g.RecommendationList[i], g.RecommendationList[j] = g.RecommendationList[j], g.RecommendationList[i]
}

func (g *RecommendationGenome) RandomGene() interface{} {
	// Return a random merchandise from the list of available merchandises
	return Merchandises[rand.Intn(len(Merchandises))]
}

func (g *RecommendationGenome) Clone() ga.Genome {
	copiedList := make([]Merchandise, len(g.RecommendationList))
	copy(copiedList, g.RecommendationList)
	return &RecommendationGenome{RecommendationList: copiedList}
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
	totalUtility := 0.0
	totalPrice := 0.0
	totalWeight := 0

	for _, merchandise := range ind.genome.RecommendationList {
		preference, exists := Preferences[merchandise.Category]
		if !exists {
			preference = 0.0 // If category not in preferences, assume 0 preference
		}
		totalUtility += preference
		totalPrice += merchandise.Price
		totalWeight += merchandise.Weight
	}

	// Check budget and capacity constraints
	if totalPrice > Budget || totalWeight > Capacity {
		return 0 // Invalid solution
	}

	return totalUtility
}

func (ind *RecommendationIndividual) Clone() ga.Individual {
	return &RecommendationIndividual{
		genome: ind.genome.Clone().(*RecommendationGenome),
	}
}
