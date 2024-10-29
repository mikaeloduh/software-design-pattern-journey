package gav2

import "math/rand"

// GeneticAlgorithm encapsulates the parameters and operations of the GA.
type GeneticAlgorithm struct {
	Population           []Individual
	MaxIterations        int
	MutationRate         float64
	CrossoverRate        float64
	TerminationCondition func(iteration int, population []Individual) bool
	SelectionOperator    SelectionOperator
	CrossoverOperator    CrossoverOperator
	MutationOperator     MutationOperator
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
func (ga *GeneticAlgorithm) Selection(pop []Individual) []Individual {
	return ga.SelectionOperator(pop)
}

// Crossover performs crossover on the selected parents to produce offspring.
func (ga *GeneticAlgorithm) Crossover(parents []Individual) []Individual {
	var offspring []Individual
	for i := 0; i < len(parents); i += 2 {
		parent1 := parents[i]
		parent2 := parents[(i+1)%len(parents)]
		if rand.Float64() < ga.CrossoverRate && ga.CrossoverOperator != nil {
			genome1, genome2 := ga.CrossoverOperator(parent1.Genome(), parent2.Genome())
			child1 := parent1.Clone()
			child2 := parent2.Clone()
			child1.SetGenome(genome1)
			child2.SetGenome(genome2)
			offspring = append(offspring, child1, child2)
		} else {
			offspring = append(offspring, parent1.Clone(), parent2.Clone())
		}
	}
	return offspring
}

// Mutation applies mutation to the offspring population.
func (ga *GeneticAlgorithm) Mutation(pop []Individual) {
	for _, individual := range pop {
		if rand.Float64() < ga.MutationRate && ga.MutationOperator != nil {
			ga.MutationOperator(individual.Genome())
		}
	}
}

// FindBestIndividual returns the individual with the highest fitness.
func (ga *GeneticAlgorithm) FindBestIndividual(pop []Individual) Individual {
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
