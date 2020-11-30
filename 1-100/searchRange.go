package main

import "fmt"

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	// 找左边界
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target || target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 存在并且相等, 即为左边界
	if left < len(nums) && nums[left] == target {
		res[0] = left
	} else {
		// 找不到, 直接返回
		return res
	}
	// 找右边界
	right = len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	res[1] = left - 1
	// 找右边界
	return res
}

func searchRangeNew(nums []int, target int) []int {
	var left, right = 0, len(nums)
	if right == 0 {
		return []int{-1, -1}
	}
	// 先看看能不能找得到
	var res = make([]int, 2)
	// 左边
	for left < right {
		var mid = left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	res[0] = left

	// 右边
	left, right = left+1, len(nums)
	for left < right {
		var mid = left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	res[1] = left - 1

	return res
}

func main() {
	var nums = []int{7}
	fmt.Println(searchRange(nums, 8))
}
