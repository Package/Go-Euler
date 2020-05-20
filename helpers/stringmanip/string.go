package stringmanip

import (
	"math"
	"strconv"
)

func Palindrome(s string) bool {
	halfWay := int(math.Ceil(float64(len(s) / 2)))

	//fmt.Printf("%s is %d characters long. Half way is %d characters long\n", s, len(s), halfWay)

	for x := 0; x <= halfWay; x++ {
		c1 := s[x]
		c2 := s[len(s) - (x+1)]
		//fmt.Printf("Comparing %c and %c\n", c1, c2)
		if c1 != c2 {
			return false
		}
	}

	return true
}

func PalindromeInt(num int) bool {
	// Convert to a string before falling back to other palindrome func
	converted := strconv.Itoa(num)

	return Palindrome(converted)
}