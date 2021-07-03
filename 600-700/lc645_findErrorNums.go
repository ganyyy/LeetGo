package main

func findErrorNumsOn(nums []int) []int {
	// 可以使用 o(n)的空间实现

	var tmp = make([]byte, (len(nums)+7)>>3)

	var set = func(i int) {
		tmp[i>>3] |= 1 << (i & 7)
	}
	var check = func(i int) bool {
		return tmp[i>>3]&(1<<(i&7)) != 0
	}

	var ret [2]int
	for _, v := range nums {
		if check(v - 1) {
			ret[0] = v
		}
		set(v - 1)
	}

	for i := range nums {
		if !check(i) {
			ret[1] = i + 1
			break
		}
	}

	return ret[:]
}

func findErrorNums(nums []int) []int {
	var ret [2]int
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] < 0 {
			// 该数字重复出现了
			ret[0] = v
		} else {
			// 每个出现的数字对应的位置置为负数
			nums[v-1] *= -1
		}
	}

	for i, v := range nums {
		// 那么大于0的那一个就是未出现的数字所在的位置
		if v > 0 {
			ret[1] = i + 1
			break
		}
	}

	return ret[:]
}
