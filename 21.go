package main

import (
	"./helpers/numbers/factors"
	"fmt"
)

var DivisorMap = map[int]int{}

func sumItems(n []int) int {
	total := 0

	for _, v := range n {
		total += v
	}

	return total
}

func main() {

	for x :=1; x < 10_000; x++ {
		DivisorMap[x] = sumItems(factors.ProperDivisors(x))
	}

	amicableSum := 0
	for x := 1; x < 10_000; x++ {
		isAmicable := DivisorMap[DivisorMap[x]] == x
		sameNumber := DivisorMap[x] == x

		if isAmicable && !sameNumber {
			amicableSum += x
		}
	}

	fmt.Printf("Answer is: %d", amicableSum)
}
