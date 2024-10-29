package gav2

// Individual represents a solution in the population.
type Individual interface {
	Genome() Genome
	Fitness() float64
	SetGenome(genome Genome)
	Clone() Individual
}
