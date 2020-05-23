package main

import (
	"fmt"
	"math/big"
)

func main() {

	x := big.NewInt(1)
	y := big.NewInt(1)
	index := 3

	for {
		nextFib := x.Add(x, y)
		length := len(nextFib.String())
		if length >= 1_000 {
			fmt.Printf("Index: %d\n", index)
			return
		}

		index += 1
		x = y
		y = nextFib
	}
}
