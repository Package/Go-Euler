package main

import (
	"./helpers/files"
	"fmt"
	"strconv"
)

func sumElements(s [4]string) int {
	total := 1

	for k, v := range s {
		intVal, _ := strconv.Atoi(string(v))
		total *= intVal
		fmt.Printf("Key: %d Value: %d\n", k, intVal)
	}

	fmt.Println("")
	return total
}

// Returns the highest product this grid element can make, in any direction
func processGridElement(x int, y int, data [][]string) int {
	totals := []int {}

	// Down
	if x < len(data) - 3 {
		totals = append(totals, sumElements([4] string { data[x][y], data[x+1][y], data[x+2][y], data[x+3][y] }))
	}

	// Up
	if x > 3 {
		totals = append(totals, sumElements([4] string { data[x][y], data[x-1][y], data[x-2][y], data[x-3][y] }))
	}

	// Right
	if y < (len(data) - 3) {
		totals = append(totals, sumElements([4] string { data[x][y], data[x][y+1], data[x][y+2], data[x][y+3] }))
	}

	// Left
	if y > 3 {
		totals = append(totals, sumElements([4] string { data[x][y], data[x][y-1], data[x][y-2], data[x][y-3] }))
	}

	// Diagonal (right)
	if x < len(data) - 3 && y > 3 {
		totals = append(totals, sumElements([4] string { data[x][y], data[x+1][y-1], data[x+2][y-2], data[x+3][y-3] }))
	}

	// Diagonal (left)
	if x < len(data) - 3 && y < len(data) - 3 {
		totals = append(totals, sumElements([4] string { data[x][y], data[x+1][y+1], data[x+2][y+2], data[x+3][y+3] }))
	}

	return biggestTotal(totals)
}

func biggestTotal(totals []int) int {
	biggest := 0
	for _, x := range totals {
		if x > biggest {
			biggest = x
		}
	}

	return biggest
}

func main() {
	data, err := files.ReadAsMultiArray("./fileinput/11.txt")
	if err != nil {
		panic(err)
	}

	bestScore := 0
	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data); y++ {
			res := processGridElement(x, y, data)
			if res > bestScore {
				bestScore = res
			}
		}
	}

	fmt.Printf("Answer is: %d\n", bestScore)
}
