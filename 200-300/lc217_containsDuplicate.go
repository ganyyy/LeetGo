package main

var empty = struct{}{}

func containsDuplicate(nums []int) bool {
	var set = make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = empty
	}
	return false
}
