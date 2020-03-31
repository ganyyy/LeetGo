package main

import "fmt"

func sortArray(nums []int) []int {
	var quickSort func(int, int)

	quickSort = func(left, right int) {
		if left >= right {
			return
		}
		i, j, v := left, right+1, nums[left]
		for i < j {
			for j--; i < j && nums[j] > v; j-- {
			}
			for i++; i < j && nums[i] < v; i++ {
			}
			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[j], nums[left] = v, nums[j]
		quickSort(left, j-1)
		quickSort(j+1, right)
	}

	quickSort(0, len(nums)-1)
	return nums
}
func main() {
	var a = []int{5, 1, 1, 2, 0, 0}
	sortArray(a)
	fmt.Println(a)
}
