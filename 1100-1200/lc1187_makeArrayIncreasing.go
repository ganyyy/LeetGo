//go:build ignore

package main

import (
	"math"
	"sort"
)

func makeArrayIncreasing(a, b []int) int {
	// 添加哨兵, 简化代码逻辑
	a = append(a, math.MaxInt)
	// 排序+去重
	sort.Ints(b)
	m := 0
	for _, x := range b[1:] {
		if b[m] != x {
			m++
			b[m] = x // 原地去重
		}
	}
	b = b[:m+1]

	// 执行替换
	n := len(a)         // 注意 a 已经添加了一个元素
	f := make([]int, n) // f[i] 表示的是 arr1[:i+1] 中哪些元素没有被替换, 越大说明被替换的元素越少
	// 最终结果就是 n-f[n-1](因为最后一个元素铁定不会替换)

	for i, x := range a {
		// 找到第一个 >= x 对应的位置
		// 那么 对于任意 < k 的索引而言, 都是严格小于 x 的数
		k := sort.SearchInts(b, x)

		res := 0 // 小于 a[i] 的数全部替换, 相当于保留0个数
		if k < i {
			res = math.MinInt // 这个咋理解呢? 此时 k < i 说明全部替换的情况下, 数量不够, 等同于无解
		}
		if i > 0 && a[i-1] < x { // 先看 i-1, 无替换, 保留前一个状态即可. 为啥要单独拉出来算呢?
			res = max(res, f[i-1]) // +1 统一放到最后
		} else {
		}
		// 这里隐含的逻辑, 就是 使用 b[k-1] 替换了 a[i-1]

		// 最多可以替换多少的元素呢?
		for j := i - 2; j > i-k-1 && j >= 0; j-- {
			// 如果是连环替换的话, 前提是得保证  b[x+1] > a[x], 此时就可以使用 b[x] 替换 a[x]
			// 对于 b而言, 迭代的区间是 [max(1, k-i+1), k-1]
			// 对于 a而言, 迭代的区间是 [max(i-k, 0), i-2]
			// 因为本身 b 已经排序且去重, 所以肯定越接近k的b越大
			if b[k-(i-j-1)] > a[j] {
				// a[j+1] 到 a[i-1] 替换成 b[k-(i-j-1)] 到 b[k-1]
				res = max(res, f[j])
			}
		}
		f[i] = res + 1 // 把 +1 移到这里，表示 a[i] 不变
	}
	if f[n-1] < 0 {
		// 对应无解的情况
		return -1
	}
	return n - f[n-1]
}

func main() {
	println(makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 3, 2, 4}))
}
