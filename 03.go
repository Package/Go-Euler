package main

import (
	"./helpers/primes"
	"fmt"
	"math"
)

func main() {
	magicNumber := 600851475143
	start := int(math.Sqrt(float64(magicNumber))) + 1

	for x := start; x > 0; x-- {
		if magicNumber % x == 0 && primes.Prime(x) {
			fmt.Printf("Answer is: %d", x)
			return
		}
	}
}