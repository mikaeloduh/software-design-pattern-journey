package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// interface
type IModel interface {
	Matrix() [][]float64
	LinearTransformation(value []float64) []float64
}

type IModels interface {
	CreateModel(name string) (IModel, error)
}

// class
type Model struct {
	matrix [][]float64
}

func (o Model) LinearTransformation(value []float64) []float64 {
	var out = make([]float64, len(o.matrix))

	for j := 0; j < len(o.matrix); j++ {
		for i := 0; i < len(out); i++ {
			out[i] = out[i] + (value[i] * o.matrix[i][j])
		}
	}

	return out
}

func (o Model) Matrix() [][]float64 {
	return o.matrix
}

type Models struct {
}

func NewModels() IModels {
	return &Models{}
}

func (m Models) CreateModel(name string) (IModel, error) {
	file, err := os.Open(fmt.Sprintf("./data/%s.mat", name))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix = make([][]float64, 1000)
	var index int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix[index], _ = splitStringToFloatArray(scanner.Text())

		index += 1
	}

	return Model{matrix: matrix}, nil
}

func splitStringToFloatArray(line string) ([]float64, error) {
	strArray := strings.Split(line, " ")
	floatArray := make([]float64, len(strArray))

	for i, str := range strArray {
		f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
		if err != nil {
			return nil, err
		}
		floatArray[i] = f
	}

	return floatArray, nil
}

func main() {
	fmt.Print("Hello World!")
}
