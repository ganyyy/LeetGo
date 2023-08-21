package main

func majorityElement(nums []int) int {
	var num int
	var cnt int

	for _, v := range nums {
		if cnt == 0 {
			num = v
			cnt++
		} else if cnt > 0 && v == num {
			cnt++
		} else {
			cnt--
		}
	}

	return num
}
