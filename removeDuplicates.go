package main

import "fmt"

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	i := 0
	for s := 1; s < l; s++ {
		if nums[i] != nums[s] {
			i++
			nums[i] = nums[s]
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 4, 4}
	l := removeDuplicates(nums)
	fmt.Println(nums, l)
}
