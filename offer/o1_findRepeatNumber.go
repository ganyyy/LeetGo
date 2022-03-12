package main

import (
	"fmt"
	"math/rand"
)

func findRepeatNumber(nums []int) int {
	// 这需要消耗on的空间,
	// 如果对空间有需求, 可以
	set := make([]bool, len(nums))
	for _, v := range nums {
		if !set[v] {
			set[v] = true
		} else {
			return v
		}
	}
	return -1
}

func findRepeatNumber2(nums []int) int {
	for i, v := range nums {
		for i != v {
			var next = nums[v]
			if next == v {
				return v
			}
			nums[i], nums[v] = nums[v], nums[i]
			v = next
		}
	}
	return -1
}

func main() {
	var src = []int{5, 0, 1, 2, 3, 4, 2, 3}
	rand.Shuffle(len(src), func(i, j int) {
		src[i], src[j] = src[j], src[i]
	})
	var ret = findRepeatNumber2(src)
	fmt.Println(ret, src)
}
