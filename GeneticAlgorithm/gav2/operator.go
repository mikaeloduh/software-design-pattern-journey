package gav2

import (
	"math/rand"
	"sort"
)

// SelectionOperator defines a method for selecting individuals from the population.
type SelectionOperator func(population []Individual) []Individual

// CrossoverOperator defines a crossover operation on genomes.
type CrossoverOperator func(parent1, parent2 Genome) (Genome, Genome)

// MutationOperator defines a mutation operation on a genome.
type MutationOperator func(genome Genome)

// NewTournamentSelection returns a SelectionOperator that performs tournament selection.
func NewTournamentSelection(tournamentSize int) SelectionOperator {
	return func(population []Individual) []Individual {
		selected := make([]Individual, len(population))
		for i := 0; i < len(population); i++ {
			best := population[rand.Intn(len(population))]
			bestFitness := best.Fitness()
			for j := 1; j < tournamentSize; j++ {
				competitor := population[rand.Intn(len(population))]
				if competitor.Fitness() > bestFitness {
					best = competitor
					bestFitness = competitor.Fitness()
				}
			}
			selected[i] = best
		}
		return selected
	}
}

// RankSelection selects individuals using rank-based selection.
func RankSelection(population []Individual) []Individual {
	// Sort the population by fitness in ascending order
	sortedPopulation := make([]Individual, len(population))
	copy(sortedPopulation, population)
	sort.Slice(sortedPopulation, func(i, j int) bool {
		return sortedPopulation[i].Fitness() < sortedPopulation[j].Fitness()
	})

	// Compute cumulative probabilities based on ranks
	totalRank := (len(population) * (len(population) + 1)) / 2
	cumulativeProbabilities := make([]float64, len(population))
	cumulativeSum := 0.0
	for i := 0; i < len(sortedPopulation); i++ {
		rank := i + 1
		probability := float64(rank) / float64(totalRank)
		cumulativeSum += probability
		cumulativeProbabilities[i] = cumulativeSum
	}

	// Select individuals based on cumulative probabilities
	selected := make([]Individual, len(population))
	for i := 0; i < len(population); i++ {
		r := rand.Float64()
		for j, cp := range cumulativeProbabilities {
			if r <= cp {
				selected[i] = sortedPopulation[j]
				break
			}
		}
	}
	return selected
}

// SinglePointCrossover performs single-point crossover on two genomes.
func SinglePointCrossover(parent1, parent2 Genome) (Genome, Genome) {
	length := parent1.Length()
	if length != parent2.Length() || length == 0 {
		// Unable to perform crossover, return clones of the parents
		return parent1.Clone(), parent2.Clone()
	}

	point := rand.Intn(length)

	child1 := parent1.Clone()
	child2 := parent2.Clone()

	for i := point; i < length; i++ {
		gene1 := parent1.GetGene(i)
		gene2 := parent2.GetGene(i)
		child1.SetGene(i, gene2)
		child2.SetGene(i, gene1)
	}

	return child1, child2
}

// RandomReplacementMutation returns a MutationOperator that performs random replacement mutation using the provided gene pool.
func RandomReplacementMutation(genome Genome) {
	length := genome.Length()
	if length == 0 {
		return
	}

	index := rand.Intn(length)
	newGene := genome.RandomGene()
	genome.SetGene(index, newGene)
}

// InversionMutation performs inversion mutation on a genome.
func InversionMutation(genome Genome) {
	length := genome.Length()
	if length < 2 {
		return
	}

	start := rand.Intn(length - 1)
	end := rand.Intn(length-start-1) + start + 1

	for i, j := start, end; i < j; i, j = i+1, j-1 {
		genome.SwapGenes(i, j)
	}
}
