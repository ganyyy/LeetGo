package main

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	//将非0数字(right) 替换到前边(left)
	var left, right int

	for right < len(nums) {
		if nums[right] == 0 {
			right++
			continue
		}
		if nums[left] == 0 {
			nums[left], nums[right] = nums[right], nums[left]
		}
		right++
		left++
	}
}
