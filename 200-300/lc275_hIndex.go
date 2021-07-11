package main

func hIndex3(citations []int) int {

	for i, v := range citations {
		var cnt = len(citations) - i
		if v >= cnt {
			return cnt
		}
	}
	return 0
}
