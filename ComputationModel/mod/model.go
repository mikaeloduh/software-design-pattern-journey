package mod

type IModel interface {
	Matrix() [][]float64
	LinearTransformation(value []float64) ([]float64, error)
}
