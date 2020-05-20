package main

import (
	"./helpers/files"
	"fmt"
	"sort"
	"strings"
)

func scoreName(name string) int {
	total := 0

	for _, val := range name {
		total += int(val) - 64 // - 64 as ASCII codes for uppercase letters start from 65
	}

	return total
}

func main() {

	fileData, err := files.ReadAllAsString("./fileinput/22.txt")
	if err != nil {
		panic("Error reading file")
	}

	// Strip off the quotes and split string into array
	data := strings.Split(strings.ReplaceAll(fileData,"\"", ""), ",")

	// Sort the array alphabetically
	sort.Strings(data)

	// Total up each word according to their ASCII values + position in array
	total := 0
	for x := 0; x < len(data); x++ {
		total += scoreName(data[x]) * (x + 1)
	}

	fmt.Printf("Answer is: %d", total)
}
