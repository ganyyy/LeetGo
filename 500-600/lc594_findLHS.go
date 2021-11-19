package main

import "sort"

func findLHS(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	sort.Ints(nums)

	var left, right = 0, 1
	var cur int
	var ret int
	var checkFlag bool
	for right < len(nums) {
		if nums[right]-nums[left] <= 1 {
			if nums[right] != nums[left] {
				checkFlag = true
			}
			right++
			cur = right - left
			continue
		}
		if checkFlag {
			ret = max(ret, cur)
		}
		cur = 0
		checkFlag = false
		left++
		for nums[left-1] == nums[left] {
			left++
		}
		if left == right {
			right++
		}
	}
	if left == 0 && nums[left] == nums[right-1] {
		return 0
	}
	if checkFlag {
		ret = max(ret, cur)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLHSBig(nums []int) (ans int) {
	cnt := make(map[int]int, len(nums))
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		if c1, ok := cnt[num+1]; ok && c+c1 > ans {
			ans = c + c1
		}
	}
	return
}

func main() {
	println(findLHS([]int{1, 4, 1, 3, 1, -14, 1, -13}))
}
