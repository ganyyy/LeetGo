package main

import "sort"

func numFactoredBinaryTrees(arr []int) (ans int) {
	sort.Ints(arr)
	n := len(arr)
	idx := make(map[int]int, n)
	for i, x := range arr {
		idx[x] = i
	}
	f := make([]int, n)
	for i, val := range arr {
		f[i] = 1
		for j, x := range arr[:i] {
			if x*x > val {
				break
			}
			if x*x == val {
				// 重复的数
				f[i] += f[j] * f[j]
				break
			}
			if val%x == 0 {
				// 两个因数
				if k, ok := idx[val/x]; ok {
					// 左右子树互换也算一种
					f[i] += f[j] * f[k] * 2
				}
			}
		}
		ans += f[i]
	}
	return ans % 1000000007
}
