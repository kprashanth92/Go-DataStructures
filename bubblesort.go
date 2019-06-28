package main

import (
	"fmt"
)

var toBeSorted [10]int = [10]int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}

//o(n2)

func bubbleSort(input [10]int) {
	n := len(input) // get the length of array
	swapped := true
	// keep swapping until all elements are swapped
	for swapped {
		// set swapped to false
		swapped = false
		// the  below loop would compare , keep swapping the largest number to the end of array.
		for i := 1; i < n; i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				// even if a single element is swapped then keep continuing.
				swapped = true
			}
		}
	}

	fmt.Println(input)
}

func main() {
	bubbleSort(toBeSorted)
}
