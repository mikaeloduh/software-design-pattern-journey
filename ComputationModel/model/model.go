package model

type IModel interface {
	Matrix() [][]float64
	LinearTransformation(value []float64) ([]float64, error)
}
