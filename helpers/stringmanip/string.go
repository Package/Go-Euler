package stringmanip

import (
	"math"
	"strconv"
	"strings"
)

// Returns whether or not a string is considered a palindrome.
// That is when it reads the same backwards, e.g. "Hannah" or "Bob"
func Palindrome(s string) bool {
	halfWay := int(math.Ceil(float64(len(s) / 2)))

	for x := 0; x <= halfWay; x++ {
		c1 := s[x]
		c2 := s[len(s) - (x+1)]
		if c1 != c2 {
			return false
		}
	}

	return true
}

// Returns whether or not a int is considered a palindrome.
// That is when it reads the same backwards, e.g. "9009" or "000000"
func PalindromeInt(num int) bool {
	// Convert to a string before falling back to other palindrome func
	converted := strconv.Itoa(num)

	return Palindrome(converted)
}

// Sums each individual digit within the provided string and returns the total.
func SumOfDigits(s string) int {
	total := 0

	for x := 0; x < len(s); x++ {
		value, err := strconv.Atoi(string(s[x]))
		if err != nil {
			panic("Couldn't convert to string")
		}
		total += value
	}

	return total
}

// Returns the string with all spaces and hyphens stripped out
func StripSpacesAndHyphens(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "-", "")
}