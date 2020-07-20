package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if t := numbers[left] + numbers[right]; t == target {
			return []int{left + 1, right + 1}
		} else if t > target {
			right--
		} else {
			left++
		}
	}
	return nil
}

func binarySearch(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		cur := numbers[i]
		if cur <= target {
			if target-cur >= cur {
				idx := findIdx(numbers[i+1:], target-cur)
				if idx != -1 {
					return []int{i + 1, idx + i + 2}
				}
			} else {
				break
			}
		} else {
			break
		}
	}
	return nil
}

// 二分查找, 找不到返回 -1, 找到返回对应的索引
func findIdx(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		t := nums[mid]
		if t == target {
			return mid
		}
		if t < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

func main() {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
