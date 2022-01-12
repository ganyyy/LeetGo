package main

func dominantIndex(nums []int) int {
	var m1, m2 = -1, -1

	for i, v := range nums {
		if m1 == -1 {
			m1 = i
			continue
		}
		if nums[m1] < v {
			m1, m2 = i, m1
			continue
		}
		if m2 == -1 || nums[m2] < v {
			m2 = i
		}
	}

	if m2 == -1 || nums[m1] >= nums[m2]<<1 {
		return m1
	}
	return -1
}
