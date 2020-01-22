package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	ln := len(nums)
	if ln == 0 {
		return -1
	}
	if ln == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// 二分找
	left, right, index := 0, ln-1, -1
	for left < right {
		if nums[left] == target {
			index = left
			break
		}
		if nums[right] == target {
			index = right
			break
		}
		// 看中值
		mid := (left + right) / 2
		if nums[mid] == target {
			index = mid
			break
		}

		if nums[mid] < nums[right] {
			// 右边升序
			if nums[mid] < target && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 左边升序
			if nums[mid] > target && target > nums[left] {
				right = mid - 1
			} else {
				left = mid
			}
		}
	}
	return index
}

func main() {
	var nums = []int{2, 4, 5, 6, 7, 0, 1}
	fmt.Println(search(nums, 7))
}
