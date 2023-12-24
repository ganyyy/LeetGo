package main

import "math"

func sumOfArithmeticSequence(a, d, n int) int {
	return n * (2*a + (n-1)*d) / 2
}

func minimumPerimeter0(neededApples int64) int64 {
	var pre int
	calcAppleCount := func(pre, length int) int {
		// 当前周长对应的数量+之前保留的数量
		var ret int
		// 四个端点
		// 四个中点
		// 若干个8*中间点
		// 等差数列的求和公式?
		// 2*length
		ret += 2*4*length + 4*length + sumOfArithmeticSequence(8*(length+1), 8, length-1)
		return ret + pre
	}
	var length = 0
	for {
		pre = calcAppleCount(pre, length)
		if pre >= int(neededApples) {
			return int64(length * 8)
		}
		length++
	}
	return -1
}

func minimumPerimeter(neededApples int64) int64 {
	/*
	   A A A B B B B
	   A A A B B B B
	   A A A B B B B
	   A A A 0 C C C
	   D D D D C C C
	   D D D D C C C
	   D D D D C C C
	*/
	n := int64(math.Cbrt(float64(neededApples) / 4))
	if 2*n*(n+1)*(2*n+1) < neededApples {
		n++
	}
	return 8 * n
}
