package main

func containsNearbyDuplicate(nums []int, k int) bool {
	var m = make(map[int]int)

	for i, v := range nums {
		if idx, ok := m[v]; ok {
			if abs(i-idx) <= k {
				return true
			}
		}
		// 更新最近的位置
		m[v] = i
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
