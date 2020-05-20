package factorial

import (
	"math/big"
)

var bigFactorialMap = map[int]*big.Int{}
var intFactorialMap = map[int64]int64{}

func Factorial(n int64) int64 {
	if n == 1 {
		return 1
	}

	// If key is in map, "ok" will be true and we can return what we've worked out previously
	if cachedVal, ok := intFactorialMap[n]; ok {
		return cachedVal
	}

	// We don't know the factorial of this number yet so recursively work it out then cache the result
	intFactorialMap[n] = n * Factorial(n-1)

	return intFactorialMap[n]
}

func FactorialBigInt(n int) *big.Int {
	if n == 1 {
		return big.NewInt(1)
	}

	// If key is in map, "ok" will be true and we can return what we've worked out previously
	if cachedVal, ok := bigFactorialMap[n]; ok {
		return cachedVal
	}

	// We don't know the factorial of this number yet so recursively work it out then cache the result
	newBigInt := big.NewInt(int64(n))
	newBigInt = newBigInt.Mul(newBigInt, FactorialBigInt(n-1))
	bigFactorialMap[n] = newBigInt

	return bigFactorialMap[n]
}
