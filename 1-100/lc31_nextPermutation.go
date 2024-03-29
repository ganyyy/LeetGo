package main

import "fmt"

func nextPermutation1(nums []int) {
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
		reverse31(0, nums)
	} else {
		// 找到最后一个大于 nums[i-1]的值 nums[j]
		// 互换nums[i-1]和nums[j] 并将 nums[i:]逆序一下
		// 得到的序列就是比当前序列更大的字典序列
		val, j := nums[i-1], i
		for j < len(nums) && val < nums[j] {
			j++
		}
		nums[i-1], nums[j-1] = nums[j-1], val
		reverse31(i, nums)
	}
}

func reverse31(start int, nums []int) {
	for j, k := start, len(nums)-1; j < k; {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
}

func nextPermutation(nums []int) {

	// 1,5,4,3,2 -> 2,1,3,4,5
	// 从后向前 找到非逆序位置 [1], 5,4,3,2
	// 从[5,4,3,2]中找到首个大于1的位置 5,4,3,[2]
	// 交换[1],[2]  -> 2,5,4,3,1
	// 逆序后半部分 -> 2,1,3,4,5

	if len(nums) <= 1 {
		return
	}
	// 检测是否当前排列是最大的
	var nextIdx = -1
	// 从后向前找, 不满足降序的数并记录索引
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			nextIdx = i
			break
		}
	}

	if nextIdx != -1 {
		// 找到第一个小于 nums[nextIdx]. 这里使用二分查找
		var i = nextIdx + findIdx(nums[nextIdx+1:], nums[nextIdx])
		// 互换最后一个比 nextIdx大的数
		nums[nextIdx], nums[i] = nums[i], nums[nextIdx]
		// 后半段逆序处理
		reverseAll(nums[nextIdx+1:])
	} else {
		reverseAll(nums)
	}
}

// 整体逆序
func reverseAll(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

// 二分查找, 后半段是逆序的
// 我们要找的是 比给定值要大的 最小值.
// 所以 如果nums中存在和 n相等的值, 需要继续向左边查找
func findIdx(nums []int, n int) int {
	var l, r = 0, len(nums)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if nums[m] <= n {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func nextPermutation3(nums []int) {
	// 1,5,4,3,2 -> 2,1,3,4,5
	// 找到非逆序位置 [1], 5,4,3,2
	var lastPos int
	for lastPos = len(nums) - 2; lastPos >= 0 && nums[lastPos] >= nums[lastPos+1]; lastPos-- {
	}
	if lastPos == -1 {
		reverseAll(nums)
		return
	}
	// 从[5,4,3,2]中找到首个大于1的位置 5,4,3,[2]
	var val = nums[lastPos]
	var pos int
	for pos = len(nums) - 1; pos >= 0 && nums[pos] <= val; pos-- {
	}
	// 交换[1],[2]  -> 2,5,4,3,1
	nums[lastPos], nums[pos] = nums[pos], nums[lastPos]
	// 逆序后半部分 -> 2,1,3,4,5
	reverseAll(nums[lastPos+1:])
	return
}

func main() {
	var nums = []int{1, 5, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
