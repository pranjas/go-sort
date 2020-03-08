package main

import (
	"fmt"
	"math/rand" /*For random numbers and dynamic test arrays*/
	"time"      /*For seeding random generator*/
)

func SelectionSort(input []int, printComparision bool) []int {
	var result []int = make([]int, len(input)) //Need to create an array first.
	copy(result, input)                        //Make a copy. Can work with input as well.
	for select_at, _ := range result {
		//For each loop run, find the smallest and put it in
		//the select_at position. We do this by swapping it with
		//the larger element found.
		for compare_at := select_at + 1; compare_at < len(result); compare_at += 1 {
			if printComparision {
				fmt.Printf("Comparing result[%d]=%d with result[%d]=%d\n",
					select_at, result[select_at], compare_at, result[compare_at])
			}
			//If current position we want to fill has a higher value
			//then swap it with the lower value found.
			//If we do this with an auxilary array the we must not
			//use the values we've seen so far. A swap thus makes
			//this easier to keep track of, as we'll eventually get
			//to the bigger value.
			if result[select_at] > result[compare_at] {
				result[select_at], result[compare_at] =
					result[compare_at], result[select_at]
			}
		}
	}
	return result

}
func main() {
	test_data := []int{1, 9, 0, -2, 4, 6, 8, 100, 30, 3}
	sorted_array := SelectionSort(test_data, false)
	//Our random test data.
	var random_data []int
	random_source := rand.NewSource(time.Now().UnixNano())
	random_generator := rand.New(random_source)
	random_data = make([]int, random_generator.Intn(1000)) //Create a dynamic array at most 1000 entries.
	for i := 0; i < len(random_data); i += 1 {
		random_data[i] = random_generator.Intn(10000) //Give them some values.
	}
	sorted_random_data := SelectionSort(random_data, false)
	fmt.Printf("test_data = %v\n", test_data)                   //Static array
	fmt.Printf("sorted_data = %v\n", sorted_array)              //Static array sorted.
	fmt.Printf("test_data random = %v\n", random_data)          //Dynamic array.
	fmt.Printf("sorted random data = %v\n", sorted_random_data) //Dynamic array sorted.

}
