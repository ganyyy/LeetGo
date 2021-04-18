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

func removeElement2(nums []int, val int) int {
	var left int
	for right := 0; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

func main() {
	nums := []int{3, 3, 3, 1}
	l := removeElement(nums, 3)
	fmt.Println(nums[:l])
}
