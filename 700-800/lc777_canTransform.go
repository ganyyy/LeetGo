package main

func canTransform(start, end string) bool {
	// 题目等价代换一下
	// XL -> LX, 表示为将L向左移动
	// RX -> XR, 表示为将R向右移动

	// 那么可以得出的结论是
	// 1. X的数量, L的数量, R的数量必须要相同
	// 2. start中任意L的位置都应该处于end的右边, start中任意R的位置都应该处于end的左边

	i, j, n := 0, 0, len(start)
	for i < n && j < n {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}
		if i < n && j < n {
			if start[i] != end[j] {
				return false
			}
			c := start[i]
			if c == 'L' && i < j || c == 'R' && i > j {
				return false
			}
			i++
			j++
		}
	}
	for i < n {
		if start[i] != 'X' {
			return false
		}
		i++
	}
	for j < n {
		if end[j] != 'X' {
			return false
		}
		j++
	}
	return true
}
