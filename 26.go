package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Todo: finish this. Not working yet but logic on the right path.

func BestCycle(n float64) (int, string) {
	asString := strconv.FormatFloat(n, 'f', -1, 64)
	asString = strings.ReplaceAll(asString, "0.", "")
	bestLength := 0
	bestString := ""

	for x := 0; x < len(asString); x++ {

		for y := x + 1; y < len(asString); y++ {
			currString := asString[x:y]
			if len(currString) > bestLength && strings.Count(asString, currString) > 1 {
				bestLength = len(currString)
				bestString = currString
			}
		}
	}

	return bestLength, bestString
}

func main() {

	bestLength := 0
	bestX := 0

	for x := 2; x < 1000; x++ {
		currLength, currString := BestCycle(1.0/float64(x))
		if currLength > bestLength {
			bestX = x
			bestLength = currLength
		}
		fmt.Printf("%d = %d [%s]\n", x, currLength, currString)
	}

	fmt.Printf("Answer is: %d\n", bestX)
}