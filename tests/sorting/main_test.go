package sorting

import (
	"log"
	"pks_sort/algo"
	"pks_sort/randUtil"
	"testing"
)

//Our random test data.
var random_data [][]int

const maxNumber = 1000

type sortFunc func([]int, bool) []int

var functions = []sortFunc{algo.SelectionSort, algo.BubbleSort, algo.InsertionSort, algo.MergeSort, algo.QuickSort}
var names = []string{"Selection Sort", "Bubble Sort", "Insertion Sort", "Merge Sort", "Quick Sort"}

func setupInput(maxNumber int) {
	random_data = make([][]int, maxNumber) //Create a dynamic array at most 1000 entries.
	for i := 0; i < maxNumber; i += 1 {
		random_data[i] = make([]int, maxNumber)
	}

	for i := 0; i < maxNumber; i += 1 {
		for j := 0; j < maxNumber; j += 1 {
			random_data[i][j] = randUtil.GenerateRandomInt(maxNumber) //Give them some values.
		}
	}
}

func whichAlgo(algo int) string {
	return names[algo]
}

func benchmark_algo(b *testing.B, algo int) {
	maxValue := b.N
	if maxValue > maxNumber {
		b.Logf("Changing benchmark value from %d to %d\n",
			maxValue, maxNumber)
		maxValue = maxNumber
	}
	setupInput(maxValue)
	log.Printf("Testing algorithm %s with size %d", whichAlgo(algo),
		maxValue)
	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		_ = functions[algo](random_data[i], false)
	}
}
func BenchmarkSelectionSort(b *testing.B) {
	benchmark_algo(b, 0)
}

func BenchmarkBubbleSort(b *testing.B) {

	benchmark_algo(b, 1)
}

func BenchmarkInsertionSort(b *testing.B) {

	benchmark_algo(b, 2)
}

func BenchmarkMergeSort(b *testing.B) {
	benchmark_algo(b, 3)
}

func BenchmarkQuickSort(b *testing.B) {

	benchmark_algo(b, 4)
}
