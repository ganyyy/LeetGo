package main

func countQuadruplets(a []int) (ans int) {
	// 每个值都是[1,100]
	// 那么三个加一块, 范围为[3,300]

	// c1 统计的是单个数字的计数
	// c2 统计的是两个数字的和的计数
	// c3 统计的是三个数字的和的计数
	var c1, c2, c3 [301]int
	for _, v := range a {
		ans += c3[v]
		for w, c := range c2[:201] {
			c3[v+w] += c
		}
		for w, c := range c1[:101] {
			c2[v+w] += c
		}
		c1[v]++
	}
	return
}

func countQuadruplets2(nums []int) (ret int) {
	var ln = len(nums)
	for i := 0; i < ln-3; i++ {
		for j := i + 1; j < ln-2; j++ {
			for k := j + 1; k < ln-1; k++ {
				for m := k + 1; m < ln; m++ {
					if nums[i]+nums[j]+nums[k] == nums[m] {
						ret++
					}
				}
			}
		}
	}
	return ret
}
