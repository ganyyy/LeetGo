package main

import "strconv"

func summaryRanges(nums []int) []string {
	// 双指针呗
	if len(nums) == 0 {
		return nil
	}

	var l, r = 0, 1

	var res = make([]string, 0, len(nums))
	for ; r < len(nums); r++ {
		if nums[r]-1 != nums[r-1] {
			res = append(res, getRes(nums[l], nums[r-1]))
			l = r
		}
	}
	res = append(res, getRes(nums[l], nums[r-1]))

	return res

}

func getRes(a, b int) string {
	if a == b {
		return strconv.Itoa(a)
	} else {
		return strconv.Itoa(a) + "->" + strconv.Itoa(b)
	}
}
