package algo

import (
	"fmt"
	"sort"
)

//KnapsackNaive attempts to recursively calculate the maximum
//possible value which can be created from the bag.
//The basic idea is if we've placed something in the bag then
//it doesn't matter if the bag is made bigger, since that item
//will still be there and we pick the biggest item always if
//we can fit it in the bag.

var callStack int
var DoPrint string

//This is for printing the call stack in a manner
//that it's readable and mapped to each call of the
//recursive function.
func init() {
	callStack = 0
}

func printSpaces() {
	if DoPrint == "" {
		return
	}
	for i := 0; i < callStack; i++ {
		fmt.Printf("-")
	}
}

//We need a way for us to bind value and weight so we can
//manipulate them together instead of working in indices always.
//This is ofcourse an internal structure
type item struct {
	weight int
	value  int
	seen   bool
}

type itemsWithFraction struct {
	item
	valPerWeight float64
}

//knapsackNaiveDriver is the main routine which actually
//attempts to find the maximum knapsack value. The problem
//with this method is the repeated computations.
//
//When we pick up one item and attempt to find the maximum
//value of the remaining knapsack with the remaining items.
//
//So it's possible that for the same size we might've to
//recompute the same thing for the remaining items multiple
//times.
//
//One way is to save these values, i.e. based on the size
//and the items for which we want to calculate the optimal
//knapsack value. Thus it would be a tuple such as
// size , item 1, item 2, item 3 ... item n, knapsack-value.
//
//For one knapsack size there may be multiple values which might
//be possible, thus if we consider a map based on sizes, then each
//value of such a map is itself a map containing the values for each
//possible combination of items. Thus
//map[size] = a map of values for each combination of items.
//
//If we consider the items to be bits then the map values of each
//combination will just be number corresponding to that bit combination.
//However since we can have large amount of items, say over a 1000, it's
//better to have String representation of these bits for saving the
//optimal values.
func knapsackNaiveDriver(capacity int, items []item) int {
	maxValue := 0
	if capacity <= 0 {
		return 0
	}
	callStack++
	printSpaces()
	if DoPrint != "" {
		fmt.Printf(">Calculating maximum for size %d\n", capacity)
	}

	for i := 0; i < len(items); i++ {
		//Don't look at the item already seen so far.
		if items[i].seen {
			continue
		}
		if items[i].weight <= capacity {
			items[i].seen = true
			printSpaces()
			if DoPrint != "" {
				fmt.Printf(">Selected %d\n", items[i].value)
			}
			callStack++
			lesserBag := knapsackNaiveDriver(capacity-items[i].weight, items)
			if lesserBag+items[i].value > maxValue {
				maxValue = lesserBag + items[i].value
				printSpaces()
				if DoPrint != "" {
					fmt.Printf("Picking item %d with value %d, currentMax = %d\n",
						i, items[i].value, maxValue)
				}
			}
			items[i].seen = false
			callStack--
		}
	}
	callStack--
	return maxValue

}

func KnapsackNaive(capacity int, weights, values []int) int {
	var items []item
	for i := 0; i < len(weights); i++ {
		it := item{weight: weights[i], value: values[i], seen: false}
		items = append(items, it)
	}
	return knapsackNaiveDriver(capacity, items)
}

//This is so that we can use sort on the slice.
type itemsWithFractionSlice []itemsWithFraction

func (x itemsWithFractionSlice) Len() int {
	return len(x)
}

func (x itemsWithFractionSlice) Less(i, j int) bool {
	return x[i].valPerWeight < x[j].valPerWeight
}

func (x itemsWithFractionSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

var knapsackFastMap map[int]map[int]int

func ResetKnapsack() {
	knapsackFastMap = nil
	hits = 0
	miss = 0
	callDepth = 0
	maxCallDepth = 0
}

//KnapsackFast is a faster version of the Naive algorithm. The
//problem with Naive algorithm is that it'll try out all possible
//outcomes and will also recalculate things that might've been
//calculated before. The fast algorithm makes one slight change
//to this by
//
//1. Saving results from previous computations for a given input.
//
//2. We use a map since all values won't be required and we can
//save quite a bit of space.
//The greedy part says that once we make a decision which is optimal
//i.e find out the best solution from a given set of values we never
//have to change it.
//This is a combination of both greedy and Dynamic Programming
//and is guaranteed to find "one of the" solutions
//containing the maximum possible knapsack value.

var hits int
var miss int
var callDepth int
var maxCallDepth int

func decreaseAndSetCallDepth() {
	if maxCallDepth < callDepth {
		maxCallDepth = callDepth
	}
	callDepth -= 1
}
func knapsackFastDriver(capacity int, items []item) int {
	maxValue := 0
	callDepth += 1
	defer decreaseAndSetCallDepth()

	if len(items) == 0 || capacity == 0 {
		hits += 1
		return maxValue
	}
	if knapsackFastMap == nil {
		knapsackFastMap = make(map[int]map[int]int)
	}

	if _, ok := knapsackFastMap[capacity]; ok {
		if _, ok = knapsackFastMap[capacity][len(items)]; ok {
			hits += 1
			return knapsackFastMap[capacity][len(items)]
		}
	}
	miss += 1

	//If we're able to pick this item in knapsack.
	max1 := 0
	if items[0].weight <= capacity {
		lesserKnapsackValue := knapsackFastDriver(capacity-items[0].weight, items[1:])
		max1 = items[0].value + lesserKnapsackValue
	}
	//If we're not able to pick this item in knapsack
	max2 := knapsackFastDriver(capacity, items[1:])

	//Choose the bigger of when we pick and when we don't
	if max1 > max2 {
		maxValue = max1
	} else {
		maxValue = max2
	}
	if _, ok := knapsackFastMap[capacity]; !ok {
		knapsackFastMap[capacity] = make(map[int]int)
	}
	knapsackFastMap[capacity][len(items)] = maxValue
	return maxValue
}

func KnapsackFast(capacity int, weights, values []int) int {
	var items []item
	for i := 0; i < len(weights); i++ {
		it := item{weight: weights[i], value: values[i], seen: false}
		items = append(items, it)
	}
	result := knapsackFastDriver(capacity, items)
	if DoPrint != "" {
		fmt.Printf("Hits = %d, miss = %d\n", hits, miss)
		fmt.Printf("Max Call Depth = %d\n", maxCallDepth)
	}
	return result
}

//KnapsackFractionFast is similar to the KnapsackFast however the difference
//is that we can take partial amount of an item. Thus the only part that
//changes is how much we take and we sort based on the value / weight since
//we'll take the largest value/weight item.
//NOTE: that in case of whole item this value/weight turns into just value since
//so in effect the Whole case is a special case of Fractional Knapsack.
func KnapsackFractionFast(capacity int, weights, values []int) float64 {
	var items itemsWithFractionSlice
	maxValue := 0.0

	for i := 0; i < len(weights); i++ {
		it := itemsWithFraction{item: item{weight: weights[i], value: values[i], seen: false},
			valPerWeight: float64(values[i]) / float64(weights[i])}
		items = append(items, it)
	}
	sort.Sort(items)
	for i := len(items) - 1; i >= 0 && capacity > 0; i -= 1 {
		//If capacity is lesser than the amount of item
		//available just take what we can.
		if items[i].weight > capacity {
			maxValue += items[i].valPerWeight * float64(capacity)
			capacity = 0
		} else { //otherwise take the whole item.
			maxValue += float64(items[i].value)
			capacity -= items[i].weight
		}
	}
	return maxValue
}
