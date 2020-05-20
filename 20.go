package main

import (
	"./helpers/numbers/factorial"
	"./helpers/stringmanip"
	"fmt"
)

func main() {
	digits := factorial.FactorialBigInt(100).String()
	fmt.Printf("Answer is: %d\n", stringmanip.SumOfDigits(digits))
}
