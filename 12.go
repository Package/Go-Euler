package main

import (
	"./helpers/numbers/factors"
	"fmt"
)

// Needs serious optimisation. Runs like a dog but the answer is 76576500

func main() {
	x := 1
	triangleNum := 1

	for {
		facts := factors.Factors(triangleNum)

		if len(facts) > 500 {
			fmt.Printf("Answer is: %d\n", triangleNum)
			break
		}

		x += 1
		triangleNum += x
	}
}
