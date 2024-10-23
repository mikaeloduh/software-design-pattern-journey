package test

import (
	"geneticalgorithm/ga"
	"math/rand"
	"testing"
	"time"
)

// Product represents a product with production time and demand.
type Product struct {
	Name           string
	ProductionTime int // in hours per unit
	Demand         int // units to produce
}

// Define the products.
var Products = []Product{
	{Name: "A", ProductionTime: 2, Demand: 100},
	{Name: "B", ProductionTime: 4, Demand: 200},
	{Name: "C", ProductionTime: 6, Demand: 300},
}

// Resources available.
const (
	Machines = 2
	Workers  = 4
)

// ScheduleIndividual represents a production schedule.
type ScheduleIndividual struct {
	// For simplicity, we represent the schedule as a sequence of production tasks.
	// Each task specifies which product to produce and how many units.
	Tasks []Task
}

// Task represents a production task.
type Task struct {
	ProductIndex int // index of the product in the Products slice
	Quantity     int // quantity to produce in this task
}

// Fitness calculates the inverse of total production time if demands are met.
func (ind *ScheduleIndividual) Fitness() float64 {
	// Check if all demands are met
	demandMet := make([]int, len(Products))
	for _, task := range ind.Tasks {
		demandMet[task.ProductIndex] += task.Quantity
	}
	for i, product := range Products {
		if demandMet[i] < product.Demand {
			return 0 // Demands not met, fitness is zero
		}
	}

	// Simulate the production to calculate total time
	totalTime := ind.CalculateTotalTime()
	if totalTime == 0 {
		return 0 // Invalid schedule
	}
	// Fitness is the inverse of total time
	return 1.0 / float64(totalTime)
}

// CalculateTotalTime simulates the production schedule and returns total time.
func (ind *ScheduleIndividual) CalculateTotalTime() int {
	type Resource struct {
		AvailableAt int // time when the resource becomes available
	}

	// Initialize resources
	machines := make([]Resource, Machines)
	workers := make([]Resource, Workers)

	currentTime := 0

	for _, task := range ind.Tasks {
		product := Products[task.ProductIndex]
		unitsLeft := task.Quantity

		for unitsLeft > 0 {
			// Find the earliest available machine and worker
			machineAvailableAt := machines[0].AvailableAt
			workerAvailableAt := workers[0].AvailableAt
			machineIndex := 0
			workerIndex := 0

			for i := 1; i < Machines; i++ {
				if machines[i].AvailableAt < machineAvailableAt {
					machineAvailableAt = machines[i].AvailableAt
					machineIndex = i
				}
			}
			for i := 1; i < Workers; i++ {
				if workers[i].AvailableAt < workerAvailableAt {
					workerAvailableAt = workers[i].AvailableAt
					workerIndex = i
				}
			}

			// Start time for this unit is the max of machine and worker availability
			startTime := max(machineAvailableAt, workerAvailableAt)
			// Production time for one unit
			productionTime := product.ProductionTime
			// Update resource availability
			machines[machineIndex].AvailableAt = startTime + productionTime
			workers[workerIndex].AvailableAt = startTime + productionTime

			// Update current time
			if startTime+productionTime > currentTime {
				currentTime = startTime + productionTime
			}

			unitsLeft--
		}
	}

	return currentTime
}

// Crossover performs one-point crossover with another individual.
func (ind *ScheduleIndividual) Crossover(other ga.Individual) (ga.Individual, ga.Individual) {
	otherInd := other.(*ScheduleIndividual)

	// 確定兩個切片的最小長度
	minLen := min(len(ind.Tasks), len(otherInd.Tasks))

	// 如果最小長度為 0，無法進行交配，直接複製父代
	if minLen == 0 {
		return ind.Clone(), other.Clone()
	}

	// 隨機選擇交叉點，範圍從 0 到 minLen
	point := rand.Intn(minLen + 1) // 加 1 以包含 minLen，允許 point 等於 minLen

	// 執行交配，生成子代的任務列表
	child1Tasks := append([]Task{}, ind.Tasks[:point]...)
	child1Tasks = append(child1Tasks, otherInd.Tasks[point:]...)

	child2Tasks := append([]Task{}, otherInd.Tasks[:point]...)
	child2Tasks = append(child2Tasks, ind.Tasks[point:]...)

	child1 := &ScheduleIndividual{Tasks: child1Tasks}
	child2 := &ScheduleIndividual{Tasks: child2Tasks}
	return child1, child2
}

// Mutate randomly alters the schedule.
func (ind *ScheduleIndividual) Mutate() {
	if len(ind.Tasks) == 0 {
		return
	}
	index := rand.Intn(len(ind.Tasks))
	task := &ind.Tasks[index]
	// Randomly change the product or quantity
	if rand.Float64() < 0.5 {
		// Change product
		task.ProductIndex = rand.Intn(len(Products))
	} else {
		// Change quantity
		maxQuantity := Products[task.ProductIndex].Demand
		task.Quantity = rand.Intn(maxQuantity) + 1
	}
}

// Clone creates a deep copy of the individual.
func (ind *ScheduleIndividual) Clone() ga.Individual {
	cloneTasks := make([]Task, len(ind.Tasks))
	copy(cloneTasks, ind.Tasks)
	return &ScheduleIndividual{Tasks: cloneTasks}
}

func TestScheduling(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Initialize population
	populationSize := 50
	initialPopulation := make(ga.Population, populationSize)
	for i := 0; i < populationSize; i++ {
		ind := randomScheduleIndividual()
		initialPopulation[i] = ind
	}

	ga := ga.GeneticAlgorithm{
		Population:    initialPopulation,
		MaxIterations: 100,
		MutationRate:  0.1,
		CrossoverRate: 0.7,
	}

	bestIndividual := ga.Run()
	bestSchedule := bestIndividual.(*ScheduleIndividual)

	// Verify that demands are met
	demandMet := make([]int, len(Products))
	for _, task := range bestSchedule.Tasks {
		demandMet[task.ProductIndex] += task.Quantity
	}
	for i, product := range Products {
		if demandMet[i] < product.Demand {
			t.Errorf("Demand not met for product %s: %d < %d", product.Name, demandMet[i], product.Demand)
		}
	}

	totalTime := bestSchedule.CalculateTotalTime()

	t.Logf("Best schedule total time: %d hours", totalTime)
	t.Log("Production tasks:")
	for _, task := range bestSchedule.Tasks {
		product := Products[task.ProductIndex]
		t.Logf("Produce %d units of Product %s (Production time per unit: %d hours)",
			task.Quantity, product.Name, product.ProductionTime)
	}
}

// Generates a random schedule individual.
func randomScheduleIndividual() *ScheduleIndividual {
	ind := &ScheduleIndividual{}
	totalDemands := make([]int, len(Products))
	for i := 0; i < len(Products); i++ {
		totalDemands[i] = 0
	}

	// Randomly generate tasks until demands are met
	for {
		// Randomly select a product
		productIndex := rand.Intn(len(Products))
		// Calculate remaining demand
		remainingDemand := Products[productIndex].Demand - totalDemands[productIndex]
		if remainingDemand <= 0 {
			continue
		}
		// Randomly decide quantity for this task
		quantity := rand.Intn(remainingDemand) + 1
		ind.Tasks = append(ind.Tasks, Task{ProductIndex: productIndex, Quantity: quantity})
		totalDemands[productIndex] += quantity

		// Check if all demands are met
		demandsMet := true
		for i, product := range Products {
			if totalDemands[i] < product.Demand {
				demandsMet = false
				break
			}
		}
		if demandsMet {
			break
		}
	}

	return ind
}
