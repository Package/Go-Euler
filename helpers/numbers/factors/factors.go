package factors

import "../primes"

// Returns an array of all the factors of the provided number
func Factors(num int) []int {
	// Append here is merging the two slices together
	return append([]int {1, num}, getFactors(num)...)
}

// Returns an array of all the proper divisors of the provided number
// This is all the factors of the number, NOT including the number itself
func ProperDivisors(num int)[] int {
	return append([]int {1}, getFactors(num)...)
}

// Internally used to get all the factors of a number (NOT including 1 or the number itself)
func getFactors(num int) []int {
	factorSlice := []int {}

	for x := 2; x < num/2+ 1; x++ {
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