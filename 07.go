package main

import (
	"./helpers/numbers/primes"
	"fmt"
)

func main() {

	primesFound := []int{ 2 }
	for x := 3; len(primesFound) < 10_001; x+= 2 {
		if primes.Prime(x) {
			primesFound = append(primesFound, x)
		}
	}

	// Slices from zero so this is the 10,001st prime
	fmt.Println(primesFound[10_000])
}
