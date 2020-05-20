package main

import "fmt"

func fib() int {
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

	return total
}

func main() {

	fmt.Printf("Answer is %d", fib())
}
