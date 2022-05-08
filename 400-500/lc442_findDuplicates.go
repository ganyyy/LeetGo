package main

func findDuplicates(nums []int) []int {
	// 原地重排?

	var ret []int

	var abs = func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	for _, v := range nums {
		var n = abs(v)
		if nums[n-1] > 0 {
			nums[n-1] *= -1
		} else {
			ret = append(ret, n)
		}
	}

	return ret
}

func findDuplicatesBad(nums []int) []int {
	// 原地重排?
	var ret []int
	var cnt = make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			// 这个位置上已经有一个数了, 保存结果
			if nums[i] == nums[nums[i]-1] {
				cnt[nums[i]] = true
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	ret = make([]int, 0, len(cnt))
	for k := range cnt {
		ret = append(ret, k)
	}
	return ret
}
