package ga

// Individual represents a single solution in the population.
type Individual interface {
	Fitness() float64
	Crossover(other Individual) (Individual, Individual)
	Mutate()
	Clone() Individual
}

// Population is a collection of individuals.
type Population []Individual
