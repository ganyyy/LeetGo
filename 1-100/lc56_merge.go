package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	l, r := 0, 1
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
	fmt.Println(merge([][]int{
		{1, 4},
		{2, 3},
	}))
}
