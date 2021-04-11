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

func removeDuplicates2(nums []int) int {
	// 低于两个的不用看了
	if len(nums) <= 2 {
		return len(nums)
	}
	// 双指针走起
	var l, r int

	var pre = nums[0]
	var cnt int
	for ; r < len(nums); r++ {
		if pre != nums[r] {
			pre = nums[r]
			cnt = 1
			nums[l] = nums[r]
			l++
			continue
		}
		cnt++
		if cnt <= 2 {
			nums[l] = nums[r]
			l++
			continue
		}
	}
	return l
}

func removeDuplicates3(nums []int) int {
	i, j := 1, 2
	// 这也是双指针的一种应用, 通过预留一个空位的形式实现了 超过2的移除
	// 实际上如果是N, 这个也能这么用..
	for ; j < len(nums); j++ {
		if nums[j] != nums[i-1] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	cnt := removeDuplicates(arr)
	fmt.Println(arr[:cnt])
}
