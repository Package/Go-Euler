package main

import (
	"./helpers/stringmanip"
	"fmt"
)

var NumberMap = map[int]string {
	1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine",
	10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen",
	17: "seventeen", 18: "eighteen", 19: "nineteen",
	20: "twenty", 30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty", 90: "ninety" }

func writeWord(n int) string {
	// Edge case
	if n == 1000 { return "one thousand" }

	output := ""

	// Handling 100-999
	if n >= 100 {
		hundredCategory := NumberMap[(n / 100)]
		output += hundredCategory + " hundred"
		n %= 100
		if n > 0 { output += " and " } // separate in between e.g. three hundred and
	}

	// Handling 20-99
	if n >= 20 && n < 100 {
		tenCategory := NumberMap[(n / 10) * 10]
		output += tenCategory
		n %= 10
		if n > 0 { output += "-" } // hyphen in between e.g. forty-two
	}

	// Handling < 20
	if n > 0 {
		output += NumberMap[n]
	}

	return output
}

func main() {

	totalChars := 0
	for x := 1; x <= 1000; x++ {
		totalChars += len(stringmanip.StripSpacesAndHyphens(writeWord(x)))
	}

	fmt.Printf("Answer is: %d\n", totalChars)
}
