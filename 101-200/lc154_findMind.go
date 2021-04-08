package main

func findMin4(nums []int) int {
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

func findMin3(nums []int) int {
	// 和上道题一样呗, 先去掉重复值在搞
	var l, r = 0, len(nums) - 1
	var mid int

	for l < r {
		for l < r && nums[l] == nums[l+1] {
			l++
		}
		for l < r && nums[r] == nums[r-1] {
			r--
		}

		mid = l + (r-l)>>1
		if nums[mid] > nums[r] {
			// 左边有序
			l = mid + 1
		} else {
			// 右边有序
			r = mid
		}
	}

	return nums[l]
}

func main() {
	findMin([]int{2, 2, 2, 0, 1})
}
