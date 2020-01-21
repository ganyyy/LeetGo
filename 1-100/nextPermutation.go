package main

import "fmt"

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	// 反向找到第一个不满足降序的数,
	i := len(nums) - 1
	for ; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}
	if i == 0 {
		// 表示nums是最大序列, 那就直接整体逆序
		reverse(0, nums)
	} else {
		// 找到第一个大于 nums[i-1]的值 nums[j]
		// 互换nums[i-1]和nums[j] 并将 nums[i:]逆序一下
		// 得到的序列就是比当前序列更大的字典序列
		val, j := nums[i-1], i
		for j < len(nums) && val < nums[j] {
			j++
		}
		nums[i-1], nums[j-1] = nums[j-1], val
		reverse(i, nums)
	}
}

func reverse(start int, nums []int) {
	for j, k := start, len(nums)-1; j < k; {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
}

func main() {
	var nums = []int{1, 1, 5}
	nextPermutation(nums)
	fmt.Println(nums)
}
