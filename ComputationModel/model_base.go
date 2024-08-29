package main

import "errors"

type BaseModel struct {
	matrix [][]float64
}

func (o BaseModel) LinearTransformation(value []float64) ([]float64, error) {
	if !o.valid(value) {
		return nil, errors.New("invalid size")
	}

	var out = make([]float64, len(o.matrix))

	for j := 0; j < len(o.matrix); j++ {
		for i := 0; i < len(out); i++ {
			out[i] = out[i] + (value[i] * o.matrix[i][j])
		}
	}

	return out, nil
}

func (o BaseModel) Matrix() [][]float64 {
	return o.matrix
}

func (o BaseModel) valid(value []float64) bool {
	return len(o.matrix) == len(value)
}
