package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	var minus bool
	if dividend < 0 {
		dividend = -dividend
		minus = true
	}
	if divisor < 0 {
		divisor = -divisor
		if minus {
			minus = false
		} else {
			minus = true
		}
	}

	if dividend < divisor {
		return 0
	}

	getValue := func(a, b int) (int, int) {
		start := 0
		tmp := b
		for a >= b {
			start++
			b = tmp << start
		}
		return 1 << (start - 1), a - (b >> 1)
	}

	total := 0
	for {
		count, val := getValue(dividend, divisor)
		total += count
		dividend = val
		if val < divisor {
			break
		}
	}

	if minus {
		return -total
	} else {
		if total > math.MaxInt32 {
			total = math.MaxInt32
		}
		return total
	}
}

func main() {
	res := divide(-2147483648, -1)
	fmt.Println(res)
}
