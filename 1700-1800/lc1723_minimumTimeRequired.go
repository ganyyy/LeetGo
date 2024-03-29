package main

import (
	"math"
	"sort"
)

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	// 倒序了...
	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
	l, r := jobs[0], 0
	for _, v := range jobs {
		r += v
	}
	// 好家伙, 二分查找+递归

	// 二分查找的下限为最长的任务
	// 二分查找的上限是一个人干完所有任务
	return l + sort.Search(r-l, func(limit int) bool {
		limit += l
		workloads := make([]int, k)
		// 针对当前分配的limit, 尝试进行判断是否满足分配的条件
		var backtrack func(int) bool
		backtrack = func(idx int) bool {
			if idx == n {
				return true
			}
			cur := jobs[idx]
			for i := range workloads {
				// 尝试将当前任务分配给一个工人, 查看是否可以执行
				if workloads[i]+cur <= limit {
					workloads[i] += cur
					if backtrack(idx + 1) {
						// 减小limit
						return true
					}
					workloads[i] -= cur
				}
				// 能到这一步, 说明此次分配的工作递归无法实现
				// 如果当前工人未被分配工作，那么下一个工人也必然未被分配工作
				// 或者当前工作恰能使该工人的工作量达到了上限, 说明给当前工人分配该工作无法完成. 因为可以完成的情况下, 已经返回了.
				// 这两种情况下我们无需尝试继续分配工作
				if workloads[i] == 0 || workloads[i]+cur == limit {
					break
				}
			}
			// 增加limit
			return false
		}
		return backtrack(0)
	})
}

func minimumTimeRequired2(jobs []int, k int) int {

	var sum = make([]int, k)
	var ret = math.MaxInt32

	var dfs func(i int, used int, m int)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dfs = func(i int, used int, m int) {
		if m >= ret {
			return
		}
		if i == len(jobs) {
			ret = m
			return
		}
		if used < k {
			sum[used] = jobs[i]
			dfs(i+1, used+1, max(sum[used], m))
			sum[used] -= jobs[i]
		}
		for j := 0; j < used; j++ {
			sum[j] += jobs[i]
			dfs(i+1, used, max(sum[j], m))
			sum[j] -= jobs[i]
		}
	}

	dfs(0, 0, 0)
	return ret
}
