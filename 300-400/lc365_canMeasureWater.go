package main

func canMeasureWater(x int, y int, z int) bool {
	// x+y >= z
	if x+y < z {
		return false
	}
	if z == 0 {
		return true
	}

	if x > y {
		x, y = y, x
	}
	if x == 0 {
		return y == z
	}
	// 核心是这里: 满足 ax+by = z即可. 判断条件就是
	// 二者的最大公约数是否能被z取余
	for y%x != 0 {
		y, x = x, y%x
	}
	return z%x == 0
}

func main() {

}
