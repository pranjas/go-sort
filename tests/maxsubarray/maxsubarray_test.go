package maxsubarray

import (
	"fmt"
	"pks_sort/algo"
	"pks_sort/randUtil"
	"testing"
)

var random_data []int

func init() {
	random_data = setupInput(5000000, false)
}

func setupInput(maxNumber int, onlyPositive bool) []int {
	random_data := make([]int, maxNumber) //Create a dynamic array at most 1000 entries.
	if !onlyPositive {
		for i := 0; i < maxNumber; i += 1 {
			random_data[i] = randUtil.GenerateRandomIntIncludingNegatives(maxNumber) //Give them some values.
		}
	} else {
		for i := 0; i < maxNumber; i += 1 {
			random_data[i] = randUtil.GenerateRandomInt(maxNumber) //Give them some values.
		}
	}
	return random_data
}

func TestMaxSubArrayFast(t *testing.T) {
	test_data := random_data
	fmt.Printf("Finding max subarray in %v\n", test_data)
	if test_value, err := algo.MaxSubArray(test_data, true); err == nil {
		t.Logf("Maximum subarray is of value %d", test_value)
		fmt.Printf("Maximum subarray is of value %d\n", test_value)
	} else {
		t.Errorf("Error: %v\n", err)
	}
}

func TestMaxSubArrayNaive(t *testing.T) {
	test_data := random_data
	fmt.Printf("Finding max subarray in %v\n", test_data)
	if test_value, err := algo.MaxSubArrayNaive(test_data, true); err == nil {
		t.Logf("Maximum subarray is of value %d", test_value)
		fmt.Printf("Maximum subarray is of value %d\n", test_value)
	} else {
		t.Errorf("Error: %v\n", err)
	}
}

func BenchmarkMaxSubArrayFast(b *testing.B) {
	test_data := setupInput(500000, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algo.MaxSubArray(test_data, false)
	}
}

func BenchmarkMaxSubArrayNaive(b *testing.B) {
	test_data := setupInput(500000, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algo.MaxSubArrayNaive(test_data, false)
	}
}
