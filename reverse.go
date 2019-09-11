package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	end := 0
	if x < 0 {
		end = 1
		x = -x
	}
	res := 0

	for x > 0 {
		res = res * 10 + x % 10
		x = x / 10
	}

	if end == 1 {
		if -res < math.MinInt32 {
			return 0
		}
		return -res
	}
	if res > math.MaxInt32 {
		return 0
	}
	return res
}

func main() {
	fmt.Println(reverse(-123))
}
