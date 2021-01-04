package main

func largeGroupPositions(s string) [][]int {
	// 双指针
	var res [][]int
	var slow, fast = 0, 1
	for ; fast < len(s); fast++ {
		if s[slow] != s[fast] {
			if fast-slow >= 3 {
				res = append(res, []int{slow, fast - 1})
			}
			slow = fast
		}
	}
	if fast-slow >= 3 {
		res = append(res, []int{slow, fast - 1})
	}
	return res
}
