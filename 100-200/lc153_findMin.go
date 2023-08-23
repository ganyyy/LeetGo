package main

import "fmt"

func findMin(nums []int) int {
	// 取中值, 不管如何旋转, 如果 中值比 右值小, 说明拐点在左边, 否则拐点在右边
	// Mark
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

func findMin2(nums []int) int {
	// 极端情况的判断
	if len(nums) == 1 {
		return nums[0]
	}

	// 最小值需要满足什么条件?
	//  mid-1 > mid < mid+1

	var l, r = 0, len(nums) - 1
	var mid int

	// 运气好的好, 这个效率也会好不少
	for l <= r {
		// 有序了, 返回最小值就行
		if nums[l] < nums[r] {
			return nums[l]
		}
		// 先大致判断一下拐点的位置
		mid = l + (r-l)>>1
		// 当前只有两个数了, 返回小的那个
		if mid == l {
			if nums[l] < nums[r] {
				return nums[l]
			} else {
				return nums[r]
			}
		}
		// 三个数, 并且中间满足小于两边, 说明这是一个拐点, 返回当前值即可
		if nums[mid-1] > nums[mid] && nums[mid+1] > nums[mid] {
			return nums[mid]
		}
		// 判断向左还是向右
		if nums[mid] > nums[l] {
			// 左半部分有序, 拐点在右边
			l = mid + 1
		} else {
			// 右半部分有序, 拐点在左边
			r = mid - 1
		}
	}
	return 0
}

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}
