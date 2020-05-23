package main

import (
	"fmt"
	"time"
)

func main() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	currDate := time.Date(1901, 01, 01, 00,00,00, 00, loc)
	lastDate := time.Date(2000, 12, 31, 00, 00, 00, 00, loc)
	sundays := 0

	for currDate.Before(lastDate) {

		// Check if sunday
		if currDate.Weekday() == time.Sunday {
			sundays += 1
		}

		// Add one month
		currDate = currDate.AddDate(0, 1, 0)
	}

	fmt.Printf("Answer is: %d", sundays)
}
