package factors

import "../primes"

/* Returns all the factors of a number */
func Factors(num int) []int {
	factorSlice := []int {1, num}
	for x := 2; x < int(num / 2) + 1; x++ {
		if num % x == 0 {
			factorSlice = append(factorSlice, x)
		}
	}

	return factorSlice
}

func PrimeFactors(num int) []int {
	factorSlice := Factors(num)
	primeFactorSlice := []int {}

	for x := 0; x < len(factorSlice); x++ {
		if primes.Prime(factorSlice[x]) {
			primeFactorSlice = append(primeFactorSlice, factorSlice[x])
		}
	}

	return primeFactorSlice
}