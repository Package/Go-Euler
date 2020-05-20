package main

import "fmt"

func main() {

	total := 0
	for x := 0; x < 1000; x++ {
		if x % 3 == 0 || x % 5 == 0 {
			total += x
		}
	}

	fmt.Printf("Answer is: %d", total)
}
