package algo

import (
	"fmt"
)

func SelectionSort(input []int, printComparision bool) []int {
	//var result []int = make([]int, len(input)) //Need to create an array first.
	//copy(result, input)                        //Make a copy. Can work with input as well.
	result := input
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
