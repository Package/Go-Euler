package main

import "fmt"

func generateCollatz(n int) int {
	if n == 1 {
		return 1
	}

	switch n % 2 {
	case 0: // Even
		return n / 2
	default: // Odd
		return (3 * n) + 1
	}
}

func collatzSeq(num int) int {
	sequence := []int {num}
	currentVal := num

	for {
		nextVal := generateCollatz(currentVal)
		sequence = append(sequence, nextVal)

		// Generate collatz values until we get to the end case of the item being 1
		if nextVal == 1 {
			return len(sequence)
		}

		currentVal = nextVal
	}
}

func main() {

	bestVal := 1
	bestLength := 1

	for x := 1; x < 1_000_000; x++ {
		valuesInSequence := collatzSeq(x)
		if valuesInSequence > bestLength {
			bestLength = valuesInSequence
			bestVal = x
		}
	}

	fmt.Printf("Sequence starting with %d contains %d elements\n", bestVal, bestLength)
}
