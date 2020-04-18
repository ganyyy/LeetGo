package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	res, sum := nums[0], 0
	for _, v := range nums {
		// 首先要保证和是大于零的, 不然就是当前值
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		// 取每次的和和结果作比较, 大的留下啦
		res = max(sum, res)
	}
	return res
}

func maxSubArrayDP(nums []int) int {
	ln := len(nums)
	sum, m := nums[0], nums[0]
	for i := 1; i < ln; i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		m = max(m, sum)
	}
	return m
}

func maxSubArray2(nums []int) int {

	res, max := 0, math.MinInt32
	for i := 0; i < len(nums); i++ {
		if v := res + nums[i]; v > 0 {
			res = v
		} else {
			res = 0
		}
		if res > max {
			max = res
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
