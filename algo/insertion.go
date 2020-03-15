package algo

import (
	"fmt"
)

func InsertionSort(input []int, verbose bool) []int {
	//Don't modify the original input.
	//result := make([]int, len(input))
	//copy(result, input)
	result := input
	//A single element is always sorted.
	//Start looking from the next element, if any.
	for i := 1; i < len(result); i += 1 {
		//Find where this element goes in the
		//sorted half. The sorted half at any time
		//is from index 0 to index i(inclusive).
		valueToInsert := result[i]
		sorted_index := i - 1
		for ; sorted_index >= 0; sorted_index -= 1 {
			//If the valueToInsert is lesser than the last
			//index in the sorted half then we must move
			//that value right by one.
			if verbose {
				fmt.Printf("Comparing result[%d] with %d\n", sorted_index, valueToInsert)
			}
			if result[sorted_index] > valueToInsert {
				//Move the value in the sorted half to the right
				//by one position.
				result[sorted_index+1] = result[sorted_index]
			} else {
				//We found the position where this values goes.
				break
			}
		}
		//Just put the value there and now our sorted half has
		//increased by one element more.
		//Move to the next element in the outer loop now
		//and repeat.
		result[sorted_index+1] = valueToInsert
	}
	return result
}
