package main

import "fmt"

func isTriplet(x, y, z int) bool {
	return x*x+y*y == z*z
}

func main() {

	for x := 1; x < 1000; x++ {
		for y := x + 1; y < 1000; y++ {
			for z := y + 1; z < 1000; z++ {
				if x + y + z == 1000 && isTriplet(x, y, z) {
					fmt.Printf("Answer is %d\n", x*y*z)
					return
				}
			}
		}
	}

	panic("Couldn't find it")
}
