package main

func uniqueOccurrences(arr []int) bool {
	var m = make(map[int]int, len(arr))
	// 第一次循环统计每个数字的个数
	for _, v := range arr {
		m[v]++
	}
	// 第二次循环根据个数查看是否有重复的
	var cm = make(map[int]struct{}, len(m))
	var empty = struct{}{}
	var has bool
	for _, c := range m {
		if _, has = cm[c]; has {
			return false
		}
		cm[c] = empty
	}
	return true
}
