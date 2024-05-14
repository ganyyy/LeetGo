package main

import "slices"

func findMinimumTime(tasks [][]int) (ans int) {
	// 按照结束时间正序排序
	slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
	run := make([]bool, tasks[len(tasks)-1][1]+1)
	for _, t := range tasks {
		start, end, d := t[0], t[1], t[2]
		for _, b := range run[start : end+1] { // 去掉运行中的时间点
			if b {
				// 可以一起运行, 所以单独的时间点就能减去
				d--
			}
		}
		// 补充到后缀里, 顺便也上个标记.
		// 尽可能地往后靠, 这样可以让重叠的区间更多!
		for i := end; d > 0; i-- { // 剩余的 d 填充区间后缀
			if !run[i] {
				run[i] = true // 运行
				d--
				ans++
			}
		}
	}
	return
}
