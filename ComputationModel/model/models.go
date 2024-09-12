package model

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// IModels
type IModels interface {
	CreateModel(name string) (IModel, error)
}

// models
type models struct {
}

var once sync.Once
var modelsInstance *models

func NewModels() IModels {
	if modelsInstance == nil {
		once.Do(func() {
			modelsInstance = &models{}
		})
	}
	return modelsInstance
}

var modelInstances = make(map[string]IModel)
var lock = &sync.Mutex{}

func (m *models) CreateModel(name string) (IModel, error) {
	// the first nil-check is to prevent expensive lock operations
	if modelInstances[name] == nil {
		lock.Lock()
		defer lock.Unlock()
		// the second nil-check ensure that if more than one goroutine bypasses the first check,
		// only one goroutine can create the singleton instance.
		if modelInstances[name] == nil {
			modelInstances[name], _ = m.newModel(name)
		}
	}

	return modelInstances[name], nil
}

func (m *models) newModel(name string) (IModel, error) {
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
