//go:build ignore

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

const (
	MAX = 0x7FFFFFFF
	MIN = -0x80000000
)

func abs(val int) uint {
	if val < 0 {
		return -uint(val)
	}
	return uint(val)
}

func divide2(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == MIN && divisor == -1 {
		return MAX
	}
	// 判断结果的正负
	neg := (dividend ^ divisor) < 0
	res := 0
	// 计算时都转换为正值
	t := abs(dividend)
	d := abs(divisor)
	for i := 31; i >= 0; i-- {
		// 从高位开始, 依次计算当前值能满足d的最大值
		p := uint(i)
		if t>>p >= d {
			res += 1 << p
			t -= d << p
		}
	}
	if neg {
		return -res
	}
	return res
}

func main() {
	res := divide(-2147483648, -1)
	fmt.Println(res)
}
