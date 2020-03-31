package main

func findMin(nums []int) int {
	// 因为可能会出现重复字符
	// 如果出现了, 就意味着中值可能和右边相等
	// 直接排除即可
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] == nums[right] {
			right--
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	findMin([]int{2, 2, 2, 0, 1})
}
