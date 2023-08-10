package main

import "fmt"

func canJump(nums []int) bool {
	// end是下一次的跳点
	// m是在到达下一次跳点前 跳点和当前位置能跳的最远距离的最大值
	m := 0
	for i := 0; i < len(nums); i++ {
		if i > m {
			return false
		}
		m = max(m, i+nums[i])
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
