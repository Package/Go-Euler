package factors

import "../primes"

// Returns an array of all the factors of the provided number
func Factors(num int) []int {
	factorSlice := []int {1, num}
	for x := 2; x < int(num / 2) + 1; x++ {
		if num % x == 0 {
			factorSlice = append(factorSlice, x)
		}
	}

	return factorSlice
}

// Returns an array of all the factors of the provided number that are PRIME
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