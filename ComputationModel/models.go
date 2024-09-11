package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type IModels interface {
	CreateModel(name string) (IModel, error)
}

var once sync.Once
var instance IModel

type Models struct {
}

func NewModels() IModels {
	return &Models{}
}

func (m Models) CreateModel(name string) (IModel, error) {
	once.Do(func() {
		instance, _ = m.newModel(name)
	})
	return instance, nil
}

func (m Models) newModel(name string) (IModel, error) {
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

	return &baseModel{matrix: matrix}, nil
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
