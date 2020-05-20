package main

import "fmt"

const GridSize int = 21

var GridMap = [GridSize][GridSize]int{}

// Default all values in the map to -1
func initializeMap() {
	for x := 0; x < GridSize; x++ {
		for y := 0; y < GridSize; y++ {
			GridMap[x][y] = -1
		}
	}
}

func numberOfPaths(x, y int) int {
	if outOfBounds(x, y) {
		return 0
	}

	// At the end
	if x == GridSize - 1 && y == GridSize - 1 {
		return 1
	}

	if GridMap[x][y] != -1 {
		return GridMap[x][y]
	}

	nextPosition := numberOfPaths(x + 1, y) + numberOfPaths(x, y + 1)
	GridMap[x][y] = nextPosition

	return nextPosition
}

func outOfBounds(x, y int) bool {
	return x < 0 || x >= GridSize || y < 0 || y >= GridSize
}

func main() {
	initializeMap()
	fmt.Printf("Answer is: %d", numberOfPaths(0, 0))
}
