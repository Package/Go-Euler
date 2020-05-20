package primes

import "math"

// Returns whether or not a number is prime
func Prime(num int) bool {
	// Edge cases
	if num == 2 {
		return true
	}

	// All even numbers (bar 2) cannot be prime
	if num < 2 || num % 2 == 0 {
		return false
	}

	// Only need to check upto square root +1
	upperCheck := int(math.Sqrt(float64(num))) + 1

	for x := 3; x < upperCheck ; x+=2 {
		if num % x == 0 {
			return false
		}
	}

	// If we get here then passed all checks so it's prime
	return true
}