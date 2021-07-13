package main

import (
	"math"
	"sort"
)

func minAbsoluteSumDiffError(nums1 []int, nums2 []int) int {
	// 第一遍找到差值最大的位数

	// 第二遍找到距离差值最大的数最接近的那个数

	// 这个解法是错误的.

	var ret int
	const MOD = 1e9 + 7
	var maxAbs int
	var maxVal int
	for i := 0; i < len(nums1); i++ {
		var cur = abs(nums1[i] - nums2[i])
		if cur > maxAbs {
			maxAbs = cur
			maxVal = nums2[i]
		}
		if ret = ret + cur; ret > MOD {
			ret %= MOD
		}
	}
	// 相等, 直接返回
	if maxAbs == 0 {
		return 0
	}

	// 查找最接近的那个数, 替换后重新计算结果
	var minAbs = math.MaxInt32
	for _, v := range nums1 {
		if v == maxVal {
			continue
		}
		minAbs = min(minAbs, abs(maxVal-v))
	}

	ret += minAbs - maxAbs
	if ret > MOD {
		return ret % MOD
	}
	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minAbsoluteSumDiff(nums1, nums2 []int) int {
	rec := append(sort.IntSlice(nil), nums1...)
	rec.Sort()
	sum, m, n := 0, 0, len(nums1)
	for i, v := range nums2 {
		// 计算总和
		diff := abs(nums1[i] - v)
		sum += diff

		// 从 nums1.sort 中查找距离v最近的值
		j := rec.Search(v)
		// 这个值可能大于等于 v, 也可能小于 v
		// 需要都检查一下? 这个检查逻辑没看懂啥意思
		if j < n {
			m = max(m, diff-(rec[j]-v))
		}
		if j > 0 {
			m = max(m, diff-(v-rec[j-1]))
		}
	}
	return (sum - m) % (1e9 + 7)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a = []int{1, 2, 33, 4, 123, 5123, 512}

	sort.Ints(a)

	var sa = sort.IntSlice(a)

	println(sa.Search(33))
}
