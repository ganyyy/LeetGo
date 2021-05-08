package main

import "sort"

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	// 倒序了...
	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
	l, r := jobs[0], 0
	for _, v := range jobs {
		r += v
	}
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
				// 给每个工人尝试分配一个任务, 判断是否在预期内
				if workloads[i]+cur <= limit {
					workloads[i] += cur
					if backtrack(idx + 1) {
						return true
					}
					workloads[i] -= cur
				}
				// 如果当前工人未被分配工作，那么下一个工人也必然未被分配工作
				// 或者当前工作恰能使该工人的工作量达到了上限
				// 这两种情况下我们无需尝试继续分配工作
				if workloads[i] == 0 || workloads[i]+cur == limit {
					break
				}
			}
			return false
		}
		return backtrack(0)
	})
}
