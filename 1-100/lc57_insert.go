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

// 失败了, 无法处理将数据插入到中间的情况
func insert2(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}
	l, r := 0, 1
	isInsert := false
	for r < len(intervals) {
		left, right := intervals[l], intervals[r]
		if !isInsert {
			isInsert = combine(left, newInterval)
		}
		if left[1] >= right[0] {
			if left[1] < right[1] {
				left[1] = right[1]
			}
		} else {
			l++
			intervals[l] = right
		}
		r++
	}
	// 最后还是没插入,
	if !isInsert {
		// 合并interval[i]和newInterval
		left := intervals[l]
		if !combine(left, newInterval) {
			// 新增, 判断是加前边还是后边
			if left[0] > newInterval[1] {
				intervals = append(append([][]int{}, newInterval), intervals...)
			} else {
				intervals = append(intervals, newInterval)
			}
			l++
		}
	}
	return intervals[:l+1]
}

func combine(a, b []int) bool {
	if (a[0] <= b[0] && b[0] <= a[1]) || (b[0] <= a[0] && a[0] <= b[1]) {
		if a[0] > b[0] {
			a[0] = b[0]
		}
		if a[1] < b[1] {
			a[1] = b[1]
		}
		return true
	}
	return false
}

func main() {
	fmt.Println(insert2([][]int{
		{6, 8},
	}, []int{
		1, 5,
	}))
}
