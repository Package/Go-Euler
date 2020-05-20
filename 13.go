package main

import (
	"./helpers/files"
	"fmt"
	"math/big"
)


func parseBigInt(bigNumber string) *big.Int {
	n := big.NewInt(0)
	n, ok := n.SetString(bigNumber, 10)
	if !ok {
		panic("Couldn't convert to string")
	}

	return n
}

func main() {

	data, err := files.ReadAsStringArray("./fileinput/13.txt")
	if err != nil {
		panic(err)
	}

	total := big.NewInt(0)
	for _, val := range data {
		total = total.Add(total, parseBigInt(val))
	}

	fmt.Printf("Answer is: %s\n", total.String()[0:10])
}
