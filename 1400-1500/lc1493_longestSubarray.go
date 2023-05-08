//go:build ignore

package main

func longestSubarray(nums []int) int {
	return longestOnes(nums, 1) - 1 // 最后高低得删除一个数
}

func longestOnes(nums []int, k int) int {
	var ret int
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var left, right int
	var cnt int
	for right < len(nums) {
		k -= nums[right] ^ 1
		if k < 0 {
			var last = left
			for left < len(nums) && nums[left] != 0 {
				left++
			}
			// 此时 left == 0, 得跳过去
			left++
			k++
			cnt -= (left - last)
		}
		right++
		cnt++
		ret = max(ret, cnt)
	}

	return ret
}
