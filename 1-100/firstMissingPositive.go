package main

import "fmt"

func firstMissingPositive(nums []int) int {
	count := len(nums)
	for i := 0; i < count; i++ {
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
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
