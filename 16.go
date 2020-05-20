package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func sumOfDigits(num *big.Int) int {
	total := 0

	for _, val := range num.String() {
		asInt, err := strconv.Atoi(string(val))
		if err != nil {
			panic("Could not convert")
		}

		total += asInt
	}

	return total
}

func main() {

	bigPower := big.NewInt(0)
	bigPower.Exp(big.NewInt(2), big.NewInt(1000), nil)

	fmt.Printf("Answer is: %d\n", sumOfDigits(bigPower))
}
