package main

import "fmt"

func removeDuplicates(nums []int) int {
	var j, cnt int
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			cnt = 0
			nums[j] = nums[i]
		} else {
			cnt++
			if cnt < 2 {
				j++
			}
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	cnt := removeDuplicates(arr)
	fmt.Println(arr[:cnt])
}
