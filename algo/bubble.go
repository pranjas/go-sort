package algo

import (
	"fmt"
)

//Bubble sort returns a new array of integers
//which is sorted.
//Due to the two loops and no break conditions, the
//complexity is O(N^2)
func BubbleSort(input []int, verbose bool) []int {
	//result := make([]int, len(input))
	//copy(result, input)
	result := input
	//The largest value will bubble out to the end.
	//Instead of selection we just compare and swap
	//a pair of values together.
	//The following loop must always run over the whole
	//length of the array to ensure we see all values.
	for i := 0; i < len(result); i += 1 {
		//At the end of each iteration of the below loop
		//the last element holds the largest value. This
		//is because after the swaps occur in the below loop
		//the largest value bubbles to the last index. Now
		//we only need to cover the remaining elements, thus
		//the -i. The -1 is because we compare two adjacent
		//elements so we can never reach len(result) as there's
		//nothing to compare at len(result)...
		for j := 0; j < len(result)-i-1; j += 1 {
			if verbose {
				fmt.Printf("Comparing result[%d]=%d with result[%d]=%d",
					j, result[j], j+1, result[j+1])
			}
			//To sort in desceding order just reverse
			//the condition below.
			if result[j] > result[j+1] {
				result[j], result[j+1] =
					result[j+1], result[j]
			}
		}
	}
	return result
}
