package numberutil

import "../factors"

func ArraySum(n []int) int {
	total := 0
	for x := 0; x < len(n); x++ {
		total += n[x]
	}

	return total
}

func PerfectNumber(n int) bool {
	properDivisors := factors.ProperDivisors(n)
	return ArraySum(properDivisors) == n
}
