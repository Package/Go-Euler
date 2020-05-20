package main

import "fmt"

func divisibleByRange(number, minRange, maxRange int) bool {
	for x := minRange; x <= maxRange; x++ {
		if number % x != 0 {
			return false
		}
	}

	return true
}

func main() {

	x := 1
	for {
		if divisibleByRange(x, 1, 20) {
			fmt.Printf("Answer is %d", x)
			return
		}

		x += 1
	}
}
