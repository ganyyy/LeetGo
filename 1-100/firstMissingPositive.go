package main

import "fmt"

func firstMissingPositive(nums []int) int {
	count := len(nums)
	for i := 0; i < count; i++ {
		// 定位重置, 给每个正整数找到其应有的位置
		// 这里有两点需要注意一下:
		// 负数不需要处理, 超过了数组上限的数不需要处理
		for v := nums[i]; v <= count && v > 0 && v != nums[v-1]; v = nums[i] {
			nums[i], nums[v-1] = nums[v-1], nums[i]
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return count + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 7, 8, 9, 11, 12}))
}
