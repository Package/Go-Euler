package main

import (
	"./helpers/stringmanip"
	"fmt"
	"math/big"
)

func main() {

	bigPower := big.NewInt(0)
	bigPower.Exp(big.NewInt(2), big.NewInt(1000), nil)

	fmt.Printf("Answer is: %d\n", stringmanip.SumOfDigits(bigPower.String()))
}
