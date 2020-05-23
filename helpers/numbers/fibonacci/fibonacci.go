package fibonacci

var fibMap = map[int]int{1:1, 2:1}

func Fib(n int) int {
	if n <= 1 { return 1 }

	// Already worked out this one before so return it
	mapItem, ok := fibMap[n]
	if ok {
		return mapItem
	}

	nextItem := Fib(n - 1) + Fib(n - 2)
	fibMap[n] = nextItem

	return nextItem
}
