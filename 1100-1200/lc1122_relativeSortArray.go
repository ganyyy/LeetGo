package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	//O(n)???

	var m = make(map[int]int, len(arr1))
	for _, v := range arr1 {
		m[v]++
	}
	arr1 = arr1[:0]

	for _, v := range arr2 {
		for n := m[v]; n > 0; n-- {
			arr1 = append(arr1, v)
		}
		delete(m, v)
	}
	var cnt = len(arr1)
	for k, n := range m {
		for ; n > 0; n-- {
			arr1 = append(arr1, k)
		}
	}

	// 末尾升序排序
	sort.Ints(arr1[cnt:])
	return arr1
}
