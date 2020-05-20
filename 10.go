package main

import (
	"./helpers/primes"
	"fmt"
)

func main() {

	primeSum := 2 // 2 is an edge case prime

	for x := 3; x < 2_000_000; x+= 2{
		if primes.Prime(x) {
			primeSum += x
		}
	}

	fmt.Printf("Answer is %d", primeSum)
}
