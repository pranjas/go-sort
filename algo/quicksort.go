package algo

import (
	"fmt"
)

type quickSplit struct {
	Input      []int
	Done       bool
	LeftSplit  *quickSplit
	RightSplit *quickSplit
}

func isSplitDone(split *quickSplit) bool {
	if split.LeftSplit == nil && split.RightSplit == nil {
		return split.Done
	}
	result := true
	if split.LeftSplit != nil {
		result = result && split.LeftSplit.Done
	}
	if split.RightSplit != nil {
		result = result && split.RightSplit.Done
	}
	//If both splits are done then mark this split as done.
	split.Done = result
	return result
}

func partition(split *quickSplit, verbose bool) (*quickSplit, *quickSplit) {
	//If the split contains only
	//one element then we're done.
	if len(split.Input) <= 1 {
		split.Done = true
		return nil, nil
	}
	//We start from index 0, and our pivot
	//is the last element in this split.
	partIndex := 0
	pivot := split.Input[len(split.Input)-1]

	for i := 0; i < len(split.Input)-1; i += 1 {
		if split.Input[i] <= pivot {
			split.Input[partIndex], split.Input[i] =
				split.Input[i], split.Input[partIndex]
			partIndex++
		}
	}

	if verbose {
		fmt.Printf("Original split %v\n", split.Input)
	}
	//Swap the pivot with the index partIndex.
	split.Input[partIndex], split.Input[len(split.Input)-1] =
		split.Input[len(split.Input)-1], split.Input[partIndex]

	//Increment partIndex to include pivot in left half in the slice.
	splitLeft := quickSplit{Input: split.Input[0:partIndex], Done: false}
	splitRight := quickSplit{Input: split.Input[partIndex:], Done: false}
	if verbose {
		fmt.Printf("Splitted into left part = %v and right part =%v\n",
			splitLeft.Input, splitRight.Input)
	}
	//We've two more splits, assign them to the original
	//split.
	split.LeftSplit = &splitLeft
	split.RightSplit = &splitRight

	return &splitLeft, &splitRight
}

func QuickSort(input []int, verbose bool) []int {
	result := input
	initialSplit := quickSplit{Input: result, Done: false}

	//Simulate a stack
	var stack []*quickSplit
	var split *quickSplit
	stack = append(stack, &initialSplit) //Add the whole array as the first split.

	for !isSplitDone(&initialSplit) {
		split = stack[len(stack)-1]
		//If the split was resolved (done)
		//remove this split from stack.
		if isSplitDone(split) {
			stack = stack[:len(stack)-1]
			continue
		}
		p1, p2 := partition(split, verbose)

		//Add the partitions, if any back to the
		//stack. These are the new partitions from the
		//original partition.
		if p1 != nil {
			stack = append(stack, p1)
		}
		if p2 != nil {
			stack = append(stack, p2)
		}
	}
	return result
}
