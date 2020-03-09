package main

import (
	"fmt"
	"sort/algo"
	"sort/randUtil"
)

func main() {
	test_data := []int{1, 9, 0, -2, 4, 6, 8, 100, 30, 3}
	sorted_array := algo.SelectionSort(test_data, false)
	//Our random test data.
	var random_data []int
	random_data = make([]int, randUtil.GenerateRandomInt(1000)) //Create a dynamic array at most 1000 entries.
	for i := 0; i < len(random_data); i += 1 {
		random_data[i] = randUtil.GenerateRandomInt(10000) //Give them some values.
	}
	sorted_random_data := algo.SelectionSort(random_data, false)
	fmt.Printf("test_data = %v\n", test_data)          //Static array
	fmt.Printf("sorted_data = %v\n", sorted_array)     //Static array sorted.
	fmt.Printf("test_data random = %v\n", random_data) //Dynamic array.
	//TODO time the algos

	fmt.Println("Using selectionSort")
	fmt.Println("==================")
	fmt.Printf("sorted random data = %v\n", sorted_random_data) //Dynamic array sorted.

	fmt.Println("Using BubbleSort")
	fmt.Println("==================")
	fmt.Printf("sorted random data =%v\n", algo.BubbleSort(random_data, false))
}
