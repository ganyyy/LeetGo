package main

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	var cnt [1001]int

	for i := range target {
		cnt[target[i]]++
		cnt[arr[i]]--
	}

	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
