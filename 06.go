package main

import "fmt"

func main() {
	sum, squareSum := 0, 0

	for x:= 1; x <= 100; x++ {
		sum += x
		squareSum += x * x
	}

	sum *= sum

	fmt.Printf("Answer is: %d\n", sum - squareSum)
}
