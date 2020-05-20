package main

import "fmt"

func main() {

	x := 1
	y := 2
	total := 2 // Start with 2

	for {
		next := x + y

		if next > 4000000 {
			break
		}

		if next % 2 == 0 {
			total += next
		}

		x = y
		y = next

	}

	fmt.Printf("Answer is %d", total)
}
