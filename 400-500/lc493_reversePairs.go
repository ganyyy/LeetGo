package main

import "fmt"

func reversePairs(nums []int) int {

	// 归并排序
	var tmp = make([]int, len(nums))

	// 原理是相对有序
	var merge func(nums []int) int

	merge = func(nums []int) int {
		if len(nums) <= 1 {
			return 0
		}

		var mid = len(nums) >> 1
		var left, right = nums[:mid], nums[mid:]
		var cnt = merge(left) + merge(right)
		// 统计每一个i满足大于 j*2 的个数

		// 此时不管左还是右两个数组都是有序的
		// 双指针, 依次遍历每一个左边的数, 在从right中查找到比当前 lv/2 还要小的第一个数
		// 注意的是: tr不需要回退, 因为 tr 也是 有序的. 所以right中小于tr的任意数的2倍依旧小于当前的 lv
		// 每次增加的就是tr当前的索引.
		var tr int
		for _, lv := range left {
			for tr < len(right) && lv > 2*right[tr] {
				tr++
			}
			cnt += tr
		}
		// 合并
		var i, j, k int
		for i < mid && j < len(nums)-mid {
			if left[i] < right[j] {
				tmp[k] = left[i]
				i++
			} else {
				tmp[k] = right[j]
				j++
			}
			k++
		}

		if i != mid {
			copy(tmp[k:], left[i:])
		}
		if j != len(nums)-mid {
			copy(tmp[k:], right[j:])
		}

		copy(nums, tmp)

		return cnt
	}
	return merge(nums)
}

func errorReverse(nums []int) int {
	var res int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > 2*nums[j] {
				res++
			}
		}
	}
	return res
}

func main() {
	var t = []int{1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(reversePairs(t))
	fmt.Println(errorReverse(t))
}
