package main

import (
	"fmt"
	"sort"
)

func triangleNumberBad(nums []int) int {
	var cnt int
	sort.Ints(nums)

	// 这太暴力了吧
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			// 这一步可以优化成二分查找, 找到距离max最短的位置
			var max = nums[i] + nums[j]
			for k := j + 1; k < len(nums); k++ {
				if max > nums[k] {
					cnt++
				} else {
					break
				}
			}
		}
	}
	return cnt
}

func triangleNumber(nums []int) int {
	var cnt int
	sort.Ints(nums)

	for i := len(nums) - 1; i > 1; i-- {
		// 双指针的正确用法
		var l, r = 0, i - 1
		// 从右向左计算, 方便进行夹逼
		for l < r {
			// 满足两边之和大于第三边
			if nums[l]+nums[r] > nums[i] {
				cnt += r - l
				r--
			} else {
				l++
			}
		}
	}
	return cnt
}

func main() {
	var src = []int{1, 2, 3, 3, 4, 5, 6}
	var idx = sort.Search(len(src), func(i int) bool {
		return src[i] > 3
	})
	fmt.Println(idx)
}
