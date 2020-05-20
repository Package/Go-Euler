package main

import (
	"./helpers/stringmanip"
	"fmt"
)

func main() {
	biggestFound := 0

	for x := 999; x > 0; x-- {
		for y := 999; y > 0; y-- {
			product := x * y
			if product > biggestFound && stringmanip.PalindromeInt(product) {
				biggestFound = product
			}
		}
	}

	fmt.Printf("Answer is: %d", biggestFound)
}
