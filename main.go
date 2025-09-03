package main

import (
	"fmt"
)

func minMax(vals []int) (int, int) {
	var initialized bool
	var min int
	var max int

	for _, val := range vals {
		if !initialized {
			min = val
			max = val
			initialized = !initialized
			continue
		}
		if min > val {
			min = val
		}
		if max < val {
			max = val
		}

	}
	return min, max
}

func main() {
	vals := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, 100}
	min, max := minMax(vals)
	fmt.Printf("%d %d\n", min, max)
}
