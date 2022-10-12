package main

func maxChunksToSorted2(arr []int) int {
	var group int
	var curMax int

	for i, v := range arr {
		if curMax < v {
			curMax = v
		}
		if curMax == i {
			group++
		}
	}
	return group
}
