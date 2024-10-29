package gav2

// Genome represents a genome that can be manipulated by genetic operators.
type Genome interface {
	Length() int
	GetGene(index int) interface{}
	SetGene(index int, value interface{})
	SwapGenes(i, j int)
	RandomGene() interface{}
	Clone() Genome
}
