package main

import "math"

func judgePoint24(nums []int) bool {
	return judge4(float64(nums[0]), float64(nums[1]), float64(nums[2]), float64(nums[3]))
}

func judge4(a, b, c, d float64) bool {
	// a b
	return judge3(a+b, c, d) ||
		judge3(a*b, c, d) ||
		judge3(a-b, c, d) ||
		judge3(b-a, c, d) ||
		judge3(a/b, c, d) ||
		judge3(b/a, c, d) ||
		// b 3c
		judge3(c+b, a, d) ||
		judge3(c*b, a, d) ||
		judge3(c-b, a, d) ||
		judge3(b-c, a, d) ||
		judge3(c/b, a, d) ||
		judge3(b/c, a, d) ||
		// c 3d
		judge3(c+d, a, b) ||
		judge3(c*d, a, b) ||
		judge3(c-d, a, b) ||
		judge3(d-c, a, b) ||
		judge3(c/d, a, b) ||
		judge3(d/c, a, b) ||
		// b 3d
		judge3(b+d, a, c) ||
		judge3(b*d, a, c) ||
		judge3(b-d, a, c) ||
		judge3(d-b, a, c) ||
		judge3(b/d, a, c) ||
		judge3(d/b, a, c) ||
		// a 3c
		judge3(a+c, b, d) ||
		judge3(a*c, b, d) ||
		judge3(a-c, b, d) ||
		judge3(c-a, b, d) ||
		judge3(a/c, b, d) ||
		judge3(c/a, b, d) ||
		// a 3d
		judge3(a+d, b, c) ||
		judge3(a*d, b, c) ||
		judge3(a-d, b, c) ||
		judge3(d-a, b, c) ||
		judge3(a/d, b, c) ||
		judge3(d/a, b, c)
}

func judge3(a, b, c float64) bool {
	return judge2(a+b, c) ||
		judge2(a*b, c) ||
		judge2(a-b, c) ||
		judge2(b-a, c) ||
		judge2(a/b, c) ||
		judge2(b/a, c) ||
		//2
		judge2(c+b, a) ||
		judge2(c*b, a) ||
		judge2(c-b, a) ||
		judge2(b-c, a) ||
		judge2(c/b, a) ||
		judge2(b/c, a) ||
		//2
		judge2(c+a, b) ||
		judge2(c*a, b) ||
		judge2(c-a, b) ||
		judge2(a-c, b) ||
		judge2(c/a, b) ||
		judge2(a/c, b)
}

func judge2(a, b float64) bool {
	return math.Abs(a+b-24) < 0.001 ||
		math.Abs(a-b-24) < 0.001 ||
		math.Abs(b-a-24) < 0.001 ||
		math.Abs(a*b-24) < 0.001 ||
		math.Abs(a/b-24) < 0.001 ||
		math.Abs(b/a-24) < 0.001
}
