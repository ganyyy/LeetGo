package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}
	// 二分插入, 在合并
	l, r := 0, len(intervals)-1
	for l <= r {
		mid := (l + r) / 2
		if intervals[mid][0] == newInterval[0] {
			l = mid
		}
		if intervals[mid][0] < newInterval[0] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	pre := append([][]int{}, intervals[:l]...)
	intervals = append(append(pre, newInterval), intervals[l:]...)

	fmt.Println(intervals)

	l, r = 0, 1
	for r < len(intervals) {
		left, right := intervals[l], intervals[r]
		if left[1] >= right[0] {
			if left[1] < right[1] {
				left[1] = right[1]
			}
		} else if left[1] < right[0] {
			l++
			intervals[l] = right
		}
		r++
	}
	return intervals[:l+1]
}

func main() {
	fmt.Println(insert([][]int{
		{1, 3},
	}, []int{
		0, 8,
	}))
}
