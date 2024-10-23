package ga

import (
	"math/rand"
)

// Individual represents a single solution in the population.
type Individual interface {
	Fitness() float64
	Crossover(other Individual) (Individual, Individual)
	Mutate()
	Clone() Individual
}

// Population is a collection of individuals.
type Population []Individual

// GeneticAlgorithm encapsulates the parameters and operations of the GA.
type GeneticAlgorithm struct {
	Population           Population
	MaxIterations        int
	MutationRate         float64
	CrossoverRate        float64
	TerminationCondition func(iteration int, population Population) bool
}

// Run executes the genetic algorithm and returns the best individual found.
func (ga *GeneticAlgorithm) Run() Individual {
	currentPopulation := ga.Population
	for iteration := 0; iteration < ga.MaxIterations; iteration++ {
		// Selection
		parents := ga.Selection(currentPopulation)
		// Crossover
		offspring := ga.Crossover(parents)
		// Mutation
		ga.Mutation(offspring)
		// Update population
		currentPopulation = offspring
		// Check termination condition
		if ga.TerminationCondition != nil && ga.TerminationCondition(iteration, currentPopulation) {
			break
		}
	}
	// Return the best individual
	return ga.FindBestIndividual(currentPopulation)
}

// Selection performs tournament selection on the population.
func (ga *GeneticAlgorithm) Selection(pop Population) Population {
	var selected Population
	tournamentSize := 3
	for i := 0; i < len(pop); i++ {
		best := pop[rand.Intn(len(pop))]
		bestFitness := best.Fitness()
		for j := 1; j < tournamentSize; j++ {
			competitor := pop[rand.Intn(len(pop))]
			if competitor.Fitness() > bestFitness {
				best = competitor
				bestFitness = competitor.Fitness()
			}
		}
		selected = append(selected, best)
	}
	return selected
}

// Crossover performs crossover on the selected parents to produce offspring.
func (ga *GeneticAlgorithm) Crossover(parents Population) Population {
	var offspring Population
	for i := 0; i < len(parents); i += 2 {
		parent1 := parents[i]
		parent2 := parents[(i+1)%len(parents)]
		if rand.Float64() < ga.CrossoverRate {
			child1, child2 := parent1.Crossover(parent2)
			offspring = append(offspring, child1, child2)
		} else {
			offspring = append(offspring, parent1.Clone(), parent2.Clone())
		}
	}
	return offspring
}

// Mutation applies mutation to the offspring population.
func (ga *GeneticAlgorithm) Mutation(pop Population) {
	for _, individual := range pop {
		if rand.Float64() < ga.MutationRate {
			individual.Mutate()
		}
	}
}

// FindBestIndividual returns the individual with the highest fitness.
func (ga *GeneticAlgorithm) FindBestIndividual(pop Population) Individual {
	best := pop[0]
	bestFitness := best.Fitness()
	for _, individual := range pop[1:] {
		fitness := individual.Fitness()
		if fitness > bestFitness {
			best = individual
			bestFitness = fitness
		}
	}
	return best
}
