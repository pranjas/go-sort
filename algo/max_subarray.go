package algo

import (
	"fmt"

	"github.com/pkg/errors"
)

// MaxSubArrayNaive is a naive algorithm to find out the subarray
// that has the maximum value. It does so by considering each element
// as a starting point and then attempts to find the biggest array
// starting from that element.
func MaxSubArrayNaive(arr []int, printOutput bool) (int, error) {
	i, j, count := 0, 0, 0
	if arr == nil || len(arr) == 0 {
		return 0, errors.Errorf("The passed in array slice can't be nil or empty")
	}

	//Since we know that we've at least 1 number in our array
	//assume that's the largest subarray, i.e. arr[0] itself is
	//the largest subarray.

	overallMax := arr[0]

	//Consider each element as the start
	//of the largest subarray. Now all
	//we need to do is keep computing
	//the sum unless we decrease the sum
	//value.

	for i = 0; i < len(arr); i++ {
		currentMax := arr[i]
		for j = i + 1; j < len(arr); j++ {
			//If we're able to increase the value
			//then we can continue to figure out
			//maximum sub-array when the starting element
			//is arr[i]
			if currentMax+arr[j] > currentMax {
				currentMax += arr[j]
			} else {
				break
			}
		}
		if printOutput {
			count += 1
			fmt.Printf("%d. Found subarray %v with sum %d\n",
				count, arr[i:j], currentMax)
		}
		//At the end of this loop we'll have computed either
		//a larger or smaller sum so just check it with the
		//curent maximum that we've.
		if overallMax < currentMax {
			overallMax = currentMax
		}
	}
	return overallMax, nil
}

// MaxSubArray provides the maximum value of the subarray using a faster
// algorithm than the naive one. Instead of checking every element and it's
// subarray this algorithm exploits the fact that sum can only increase with
// positive numbers. Note that this is a greedy algorithm by exploiting the
// mathematical property that sum increases with positive numbers.
// consider these two examples
//
// [-1, 2, 0] and [0, 2, -1]. A negative number might appear either before
// or after a positive value,
//
// in the first case though the sum increases but
// it does so because of a positive number after -1, thus we need to "consider"
// the subarray beginning at 2.
//
// In the second case, the sum will decrease which again marks the boundary for
// for a possible subarray.
func MaxSubArray(arr []int, printOutput bool) (int, error) {
	i, j, count := 0, 0, 0
	if arr == nil || len(arr) == 0 {
		return 0, errors.Errorf("The passed in array slice can't be nil or empty")
	}
	overallMax := arr[0]
	/*
	 *
	 */
	for i = 0; i < len(arr); i = j + 1 {
		currentMax := arr[i]
		for j = i + 1; j < len(arr); j++ {
			if currentMax+arr[j] > currentMax {
				currentMax += arr[j]
				//A positive number always
				//increases the sum, therefore
				//if the previous number was negative
				//then we must stop looking further
				//at this point.
				if arr[j-1] < 0 && arr[j] > 0 {
					j -= 1
					break
				}
			} else { //We ran into a number which decreased the sum
				j -= 1 //we rejected this number
				break
			}
		}
		if printOutput {
			count += 1
			if j == len(arr) {
				fmt.Printf("%d. Found subarray %v with sum  %d\n", count, arr[i:], currentMax)
			} else {
				fmt.Printf("%d. Found subarray %v with sum %d\n", count, arr[i:j+1], currentMax)
			}
		}
		if currentMax > overallMax {
			overallMax = currentMax
		}
	}
	return overallMax, nil
}
