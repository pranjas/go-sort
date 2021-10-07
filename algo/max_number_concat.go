package algo

import (
	"log"
	"strconv"
	"unicode/utf8"
)

type numberProcess struct {
	processed bool
	value     string
}

// MaxNumber returns the maximum number which can be made by
// concatenating a bunch of non-negative numbers. The numbers
// themselves can be of any length and in any order.
func MaxNumber(numbers []string) string {
	maxNumber := ""
	count := len(numbers)
	var allNumbers []*numberProcess

	for _, val := range numbers {
		allNumbers = append(allNumbers, &numberProcess{processed: false, value: val})
	}
	//Until we've numbers left
	//We do this process.
	for count > 0 {
		maxNumber += maxNumberInternal(allNumbers)
		count -= 1
	}
	return maxNumber
}

// maxNumberInternal returns the maximum number using which
// we can create a bigger number.
func maxNumberInternal(numbers []*numberProcess) string {
	currentDigit := 1     //start by looking at the first digit.
	currentMax := -1      //since negative numbers aren't possible, this acts as a sentinel
	didProcessed := false //were we able to process at least one number.
	var whichNumber *numberProcess

find_again:
	initialLength := len(numbers)
	for i := 0; i < len(numbers); i++ {
		if numbers[i].processed {
			continue
		}
		didProcessed = true
		numberDigit := getDigitAt(numbers[i].value, currentDigit)
		// if there's no current maximum digit set,
		// set it to this one.
		if currentMax < 0 {
			currentMax = numberDigit
			whichNumber = numbers[i]
			continue
		}
		// if the current maximum is lesser, then we can safely
		// remove the numbers we've seen so far. This block essentially
		// shortens the rest of the slice and starts to look up from
		// the beginning.
		if currentMax < numberDigit {
			currentMax = numberDigit
			whichNumber = numbers[i]
			numbers = numbers[i:]
			i = 0 //0 since it'll be incremented by the for loop.
		} else if currentMax > numberDigit { //if the currentMaximum is bigger just remove this number.
			numbers = removeFromSlice(numbers, i)
			i -= 1
		}
	}
	// If we processed something and there were multiple numbers
	// with the same digits then we've to try and find again.
	// Note that when we do actually come here it implies that there
	// were multiple numbers found with the same currentMax and that we need
	// to lookup again, albeit only on those numbers.
	if didProcessed && len(numbers) > 1 && initialLength != len(numbers) {
		currentDigit += 1
		didProcessed = false
		currentMax = -1
		whichNumber = nil
		goto find_again
	}
	if whichNumber != nil {
		whichNumber.processed = true
		return whichNumber.value
	}
	return ""
}

// removeFromSlice removes from a specific index within slice
// and returns a new slice without that index element in it.
func removeFromSlice(numbers []*numberProcess, index int) []*numberProcess {
	var newSlice []*numberProcess

	if len(numbers) == 0 || index < 0 || index >= len(numbers) {
		return newSlice
	}
	newSlice = append(newSlice, numbers[:index]...)
	newSlice = append(newSlice, numbers[index+1:]...)
	return newSlice
}

func getDigitAt(number string, digitNumber int) int {
	value := ""
	//Though we could've easily iterated using indexing we did it using runes
	//Note that we need to find out the digit at digitNumber and if it's not
	//possible, i.e. the number is shorter in length than digitNumber we just
	//return the last digit of number since that's the only one that matters
	//due to which we're about to do a tie break.
	for i, w := 0, 0; digitNumber > 0 && i < len(number); i += w {
		runeVal, width := utf8.DecodeRuneInString(number[i:])
		value = string(runeVal)
		w = width
		digitNumber -= 1
	}
	intVal, err := strconv.Atoi(value)
	if err != nil { //This should never happen for correct input.
		log.Fatalf("Error finding digit for number %s at index %d",
			number, digitNumber)
		return 0
	}
	return intVal
}
