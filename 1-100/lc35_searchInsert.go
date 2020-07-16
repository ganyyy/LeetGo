package main

import "fmt"

func searchInsert(nums []int, target int) int {
	// 二分查找
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] > target {
		return left
	} else {
		return left + 1
	}
}

func main() {
	a := []int{1, 3, 5, 6}
	for _, v := range []int{5, 2, 7, 0} {
		fmt.Println(searchInsert(a, v))
	}
}
