package algo

import (
	"fmt"
)

type split struct {
	Input     []int
	Sorted    bool
	Seen      bool
	Lefthalf  *split
	Righthalf *split
}

func isSplitSorted(sp *split) bool {
	return sp.Sorted
}

func MergeSort(input []int, verbose bool) []int {
	//Copy the original input.
	result := make([]int, len(input))
	copy(result, input)

	var splitList []*split

	initialSplit := split{Input: result, Sorted: false, Lefthalf: nil, Righthalf: nil}
	splitList = append(splitList, &initialSplit)
	//Loop until our initialSplit is not sorted.
	for !isSplitSorted(&initialSplit) {
		//Attempt to simulate a stack.
		//This way we don't need to rely on system allocated stack.
		//We're adding new splits at the end thus we pick them
		//from the end.
		firstSplit := splitList[len(splitList)-1]

		//If the split we picked isn't sorted,
		//then let's try to split it.
		if !isSplitSorted(firstSplit) {
			if verbose {
				prefix := ""
				if firstSplit.Seen {
					prefix = "Merging"
					fmt.Printf("%s %v and %v\n", prefix,
						firstSplit.Lefthalf.Input, firstSplit.Righthalf.Input)
				} else {
					prefix = "Splitting"
					fmt.Printf("%s %v\n", prefix, firstSplit.Input)
				}
				fmt.Printf("==============\n")
			}
			//Maybe this split just needs to be merged,
			//if that's a success then nothing more to do
			//then mark this split as sorted.
			//note that we leave this split on our list as
			//it'll be removed on the next iteration so no need
			//for adding special code here.
			if !merge(firstSplit, verbose) {
				//if the merge didn' succeed,
				//continue with the split.
				split1, split2 := splitInHalf(firstSplit, verbose)

				//if this split was success, then append the new splits,
				//at the end of our simulated stack.
				if split2 != nil {
					splitList = append(splitList, split1, split2)
				}
			}
		} else { //this split is done, remove it from the simulated stack.
			splitList = splitList[:len(splitList)-1]
		}
		//This is used purely for print purpose.
		firstSplit.Seen = true
	}
	return result
}

func splitInHalf(splitInput *split, verbose bool) (*split, *split) {
	var split1, split2 split

	if len(splitInput.Input) <= 1 {
		splitInput.Sorted = true
		if verbose {
			fmt.Printf("split %v is of length 1. Marking it as sorted\n", splitInput.Input)
		}
		return splitInput, nil
	}

	splitLength := len(splitInput.Input) / 2
	split1.Input = splitInput.Input[0:splitLength]
	split1.Sorted = false

	split2.Input = splitInput.Input[splitLength:]
	split2.Sorted = false
	splitInput.Lefthalf = &split1
	splitInput.Righthalf = &split2
	if verbose {
		fmt.Printf("Added two more splits %v and %v to split %v\n", split1.Input,
			split2.Input, splitInput.Input)
	}
	return &split1, &split2
}

//Attempt to merge a split
func merge(s1 *split, verbose bool) bool {

	if s1.Lefthalf == nil || s1.Righthalf == nil {
		return false
	}

	if !isSplitSorted(s1.Lefthalf) || !isSplitSorted(s1.Righthalf) {
		return false
	}

	leftIndex := 0
	rightIndex := 0
	auxIndex := 0
	//An aux array is required to hold the contents of both splits,
	//reason being that if we attempt to move elements then we would
	//change the permise that both splits are sorted.
	auxArray := make([]int, len(s1.Input))
	if verbose {
		fmt.Printf("Before merge %v\n", s1.Input)
	}
	for (leftIndex < len(s1.Lefthalf.Input)) && (rightIndex < len(s1.Righthalf.Input)) {
		if s1.Lefthalf.Input[leftIndex] > s1.Righthalf.Input[rightIndex] {
			auxArray[auxIndex] = s1.Righthalf.Input[rightIndex]
			rightIndex += 1
		} else {
			auxArray[auxIndex] = s1.Lefthalf.Input[leftIndex]
			leftIndex += 1
		}
		auxIndex++
	}
	//Copy the right half
	if leftIndex == len(s1.Lefthalf.Input) {
		for rightIndex < len(s1.Righthalf.Input) {
			auxArray[auxIndex] = s1.Righthalf.Input[rightIndex]
			rightIndex++
			auxIndex++
		}
	} else {
		for leftIndex < len(s1.Lefthalf.Input) {
			auxArray[auxIndex] = s1.Lefthalf.Input[leftIndex]
			leftIndex++
			auxIndex++
		}
	}
	copy(s1.Input, auxArray)
	if verbose {
		fmt.Printf("After merge %v\n", s1.Input)
	}
	s1.Sorted = true
	return true
}
