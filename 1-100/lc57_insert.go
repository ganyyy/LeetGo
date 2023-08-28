package main

import (
	"fmt"
)

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
			isInsert = combine57(left, newInterval)
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
		if !combine57(left, newInterval) {
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

func combine57(a, b []int) bool {
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

func insert3(intervals [][]int, newInterval []int) [][]int {
	var i, j int
	var insertOk bool
	for ; i < len(intervals); i++ {
		if !insertOk {
			if intervals[i][1] >= newInterval[0] {
				// 合并
				union(intervals[i], newInterval)
				insertOk = true
				j = i
			}
		} else {
			// 合并
			if intervals[j][1] >= intervals[i][0] {
				union(intervals[j], intervals[i])
			} else {
				j++
				intervals[j] = intervals[i]
			}
		}
	}

	return intervals[:j+1]
}

// 将b合并到a
func union(a, b []int) {
	if a[0] > b[0] {
		a[0] = b[0]
	}
	if a[1] < b[1] {
		a[1] = b[1]
	}
}

// 还是要看之前的
// 关于二分:
// 如果界定条件是 left <= right, 那么对应的处理是 left=mid+1, right=mid-1
// 如果界定条件是 left < right, 那么对应的处理是 left=mid+1, right = mid
func insert4(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	// 二分插入, 在合并
	l, r := 0, len(intervals)-1
	for l <= r {
		mid := (l + r) / 2
		if intervals[mid][0] == newInterval[0] {
			l = mid
			break
		}
		if intervals[mid][0] < newInterval[0] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	// 插入到区间中
	intervals = append(intervals, []int(nil))
	copy(intervals[l+1:], intervals[l:])
	intervals[l] = newInterval

	// 遍历合并
	l, r = 0, 1
	for r < len(intervals) {
		left, right := intervals[l], intervals[r]
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
	return intervals[:l+1]
}

func binarySearch1(a []int, t int) int {
	if len(a) == 0 {
		return -1
	}
	// 搜索区间是[left, right]
	var left, right = 0, len(a) - 1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if a[mid] == t {
			return mid
		}
		if a[mid] > t {
			// 搜索区间转变为[left, mid-1]
			right = mid - 1
		} else {
			// 搜索区间转变为[mid+1, left]
			left = mid + 1
		}
	}
	return -1
}

func binarySearch2(a []int, t int) int {
	if len(a) == 0 {
		return -1
	}
	var left, right = 0, len(a)
	var mid int
	// 搜索区间为[left, right)
	for left < right {
		mid = left + (right-left)/2
		if a[mid] == t {
			return mid
		}
		if a[mid] > t {
			// 搜索区间变为[left, mid)
			right = mid
		} else {
			// 搜索区间变为[mid+1, right)
			left = mid + 1
		}
	}
	return -1
}

func insert5(intervals [][]int, newInterval []int) [][]int {
	s, e := newInterval[0], newInterval[1]
	s1, e1 := s, e
	result := make([][]int, 0, len(intervals))
	i := 0
	for _, arr := range intervals {
		if s < arr[0] {
			s1 = s
			break
		} else if s >= arr[0] && s <= arr[1] {
			s1 = arr[0]
			break
		} else {
			result = append(result, arr)
		}
		i++
	}
	for ; i < len(intervals); i++ {
		arr := intervals[i]
		if e < arr[0] {
			e1 = e
			break
		} else if e >= arr[0] && e <= arr[1] {
			e1 = arr[1]
			i++
			break
		}
	}
	result = append(result, []int{s1, e1})

	for ; i < len(intervals); i++ {
		arr := intervals[i]
		result = append(result, arr)
	}
	return result
}

func main() {
	var a = []int{1, 2, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch1(a, 3))
	fmt.Println(binarySearch2(a, 3))
}
