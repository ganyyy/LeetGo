package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 间距最大的一定要去掉
	var res int
	var end = intervals[0][1] // end是最后一个不需要合并的区间的结尾值

	for i := 1; i < len(intervals); i++ {
		if end > intervals[i][0] {
			// 需要合并, 保留最小的结尾
			if end > intervals[i][1] {
				end = intervals[i][1]
			}
			res++
		} else {
			// 不需要合并, 那么下一个合法的就是i
			end = intervals[i][1]
		}
	}

	return res
}
