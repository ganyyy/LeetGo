package main

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) (ans [][]int) {
	// 单调递增趋势, 双指针定位
	// 这个也可以转换成二分?
	for x, y := 1, 1000; x <= 1000 && y > 0; x++ {
		for y > 0 && customFunction(x, y) > z {
			y--
		}
		if y > 0 && customFunction(x, y) == z {
			ans = append(ans, []int{x, y})
		}
	}
	return
}
