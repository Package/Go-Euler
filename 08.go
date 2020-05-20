package main

import (
	"./helpers/files"
	"fmt"
	"strconv"
)

func scoreDigits(s string) int {

	digitSum := 1
	for _, v := range s {

		// Strange converting due to mapping over string returning ASCII code values for the char
		valAsInt, _ := strconv.Atoi(string(v))
		digitSum *= valAsInt
	}

	return digitSum
}

func findBestAdjacentDigits(s string, numDigits int) int {
	bestScore := 0

	// We might say we want the best 4 adjacent digits but as arrays are zero based we take one off.
	realNumDigits := numDigits - 1

	for x := realNumDigits; x < len(s); x ++ {
		subStr := s[x - realNumDigits : x + 1]
		digitSum := scoreDigits(subStr)

		//fmt.Printf("%s is worth %d\n", subStr, digitSum)
		if digitSum > bestScore {
			bestScore = digitSum
		}
	}

	return bestScore
}

func main() {

	data, err := files.ReadAllAsString("./fileinput/08.txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf("Answer is: %d", findBestAdjacentDigits(data, 13))
}
