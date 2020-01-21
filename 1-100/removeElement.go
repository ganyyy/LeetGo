package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{3, 3, 3, 1}
	l := removeElement(nums, 3)
	fmt.Println(nums[:l])
}
