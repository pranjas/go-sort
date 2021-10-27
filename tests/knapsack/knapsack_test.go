package main

import (
	"log"
	"pks_sort/algo"
	"pks_sort/randUtil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type knapsackData struct {
	weights  []int
	values   []int
	capacity int
	result   interface{}
}

var benchMarkWeight []int
var benchMarkValues []int
var benchMarkCapacity int

func init() {
	i := 100
	for i > 0 {
		benchMarkWeight = append(benchMarkWeight, randUtil.GenerateRandomInt(2000)+1)
		benchMarkValues = append(benchMarkValues, randUtil.GenerateRandomInt(2000)+1)
		i -= 1
	}
	benchMarkCapacity = randUtil.GenerateRandomInt(2000) + 1
}

func TestKnapsackNaive(t *testing.T) {

	inputData := [...]knapsackData{
		knapsackData{values: []int{3, 2, 5}, weights: []int{4, 1, 10}, capacity: 6, result: 5},
		knapsackData{values: []int{1, 2, 5, 8}, weights: []int{1, 4, 5, 3}, capacity: 6, result: 9},
		knapsackData{values: []int{9, 1, 7, 3}, weights: []int{4, 1, 3, 2}, capacity: 6, result: 12},
	}

	assert := assert.New(t)
	for i, data := range inputData {
		log.Printf("Checking with data %v\n", data)
		assert.Equal(data.result, algo.KnapsackNaive(data.capacity, data.weights, data.values),
			"Oops Failed test case %d", i)
	}
}

func TestKnapsackFast(t *testing.T) {
	inputData := [...]knapsackData{
		knapsackData{values: []int{3, 2, 5}, weights: []int{4, 1, 10}, capacity: 6, result: 5},
		knapsackData{values: []int{1, 2, 5, 8}, weights: []int{1, 4, 5, 3}, capacity: 6, result: 9},
		knapsackData{values: []int{9, 1, 7, 3}, weights: []int{4, 1, 3, 2}, capacity: 6, result: 12},
	}

	assert := assert.New(t)
	for i, data := range inputData {
		log.Printf("Checking with data %v\n", data)
		algo.ResetKnapsack()
		assert.Equal(data.result, algo.KnapsackFast(data.capacity, data.weights, data.values),
			"Oops Failed test case %d", i)
	}
}

func TestKnapsackFractionFast(t *testing.T) {
	inputData := [...]knapsackData{
		knapsackData{values: []int{3, 2, 5}, weights: []int{4, 1, 10}, capacity: 6, result: 5.5},
		knapsackData{values: []int{1, 2, 5, 8}, weights: []int{1, 4, 5, 3}, capacity: 6, result: 11.0},
		knapsackData{values: []int{9, 1, 7, 3}, weights: []int{4, 1, 3, 2}, capacity: 6, result: 13.75},
	}

	assert := assert.New(t)
	for i, data := range inputData {
		log.Printf("Checking with data %v\n", data)
		assert.Equal(data.result, algo.KnapsackFractionFast(data.capacity, data.weights, data.values),
			"Oops Failed test case %d", i)
	}
}

func BenchmarkKnapsackNaive(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		algo.KnapsackNaive(benchMarkCapacity, benchMarkWeight, benchMarkValues)
	}
}

func BenchmarkKnapsackFast(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		algo.ResetKnapsack()
		algo.KnapsackFast(benchMarkCapacity, benchMarkWeight, benchMarkValues)
	}
}

func BenchmarkKnapsackFractionFast(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		algo.ResetKnapsack()
		algo.KnapsackFractionFast(benchMarkCapacity, benchMarkWeight, benchMarkValues)
	}
}
