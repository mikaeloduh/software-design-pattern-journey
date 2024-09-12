package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var instances = make(map[string]IModel)

type IModels interface {
	CreateModel(name string) (IModel, error)
}

type Models struct {
}

func NewModels() IModels {
	return &Models{}
}

var lock = &sync.Mutex{}

func (m Models) CreateModel(name string) (IModel, error) {
	if instances[name] == nil {
		lock.Lock()
		defer lock.Unlock()

		if instances[name] == nil {
			instances[name], _ = m.newModel(name)
		}
	}

	return instances[name], nil
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
