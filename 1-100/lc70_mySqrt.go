package main

import "fmt"

func mySqrt(x int) int {
	min, max := 0, x
	for max-min > 1 {
		m := (min + max) / 2
		if m > x/m {
			max = m
		} else {
			min = m
		}
	}
	return min
}

func mySqrt2(x int) int {
	if x <= 1 {
		return x
	}
	const e = 1e-10
	t := float64(x)
	xx := float64(x)
	for r := xx / t; t-xx/t > e; /**核心是这里, 如果有精度需要就是 t - x/t > 最小精度*/ r = xx / t {
		t = (t + r) / 2
	}
	fmt.Println(t)
	return int(t)
}

func main() {
	fmt.Println(mySqrt2(145))
}
