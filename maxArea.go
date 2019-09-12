package main

import (
	"fmt"
)

//常规思路
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max := 0
	minLen := 0
	for i := 0; i < len(height); i++ {
		nowI := height[i]
		for j := 0; j < i; j++ {
			nowJ := height[j]
			if nowI > nowJ {
				minLen = nowJ
			} else {
				minLen = nowI
			}
			area := minLen * (i - j)
			if area > max {
				max = area
			}
		}
	}
	return max
}

// dp思想解决问题
func maxArea2(height []int) int {

	max := 0
	for i, j := 0, len(height)-1; i < j; {
		if d := j - i; height[i] < height[j] {
			if a := d * height[i]; a > max {
				max = a
			}
			i++
		} else {
			if a := d * height[j]; a > max {
				max = a
			}
			j--
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea2([]int{1, 2, 4, 3}))
}
