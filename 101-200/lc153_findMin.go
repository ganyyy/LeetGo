package main

import "fmt"

func findMin(nums []int) int {
	// 取中值, 不管如何旋转, 如果 中值比 右值小, 说明拐点在左边, 否则拐点在右边
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}
