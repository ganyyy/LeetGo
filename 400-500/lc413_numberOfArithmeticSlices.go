package main

func numberOfArithmeticSlices(nums []int) int {
	// 双指针

	if len(nums) < 3 {
		return 0
	}

	// 做过类似的题目,

	var left, right = 0, 1

	var ret int
	for right < len(nums)-1 {
		if nums[right+1]-nums[right] == nums[right]-nums[right-1] {
			right++
			continue
		}
		if right-left >= 2 {
			ret += (right - left - 1) * (right - left) / 2
		}
		left = right
		right++
	}

	if right-left >= 2 {
		// 加上区间内的等差数列的数量
		ret += (right - left - 1) * (right - left) / 2
	}

	return ret
}

func main() {
	print(numberOfArithmeticSlices([]int{1, 2, 3, 6, 9, 6, 3, 4, 5}))
}
