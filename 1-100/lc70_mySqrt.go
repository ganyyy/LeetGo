package main

import "fmt"

func mySqrt(x int) int {
	mi, ma := 0, x
	for ma-mi > 1 {
		m := (mi + ma) / 2
		if m > x/m {
			ma = m
		} else {
			mi = m
		}
	}
	return mi
}

func mySqrt2(x int) int {
	if x <= 1 {
		return x
	}
	const e = 1e-10
	// 先假设x的平方根t是x, 然后不断的逼近.
	// 逼近的公式就是 (t + x/t) / 2, 判断的条件就是 t-x/t > e
	t := float64(x)
	S := float64(x)
	// 牛顿迭代法: f
	// f(x) = x^2 - S
	// f'(x) = 2x
	// x_next = x - f(x)/f'(x) = x - (x^2-S)/2x = (x+S/x)/2
	for t-S/t > e /**核心是这里, 如果有精度需要就是 t - x/t > 最小精度*/ {
		t = (t + S/t) / 2
	}
	fmt.Println(t)
	return int(t)
}

func main() {
	fmt.Println(mySqrt2(145))
}
