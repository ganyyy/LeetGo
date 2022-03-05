package main

func goodDaysToRobBank(security []int, time int) []int {
	// 确认每个位置连续递增/连续递减的个数
	var num = len(security)

	// 左边有几个比我大的, 右边有几个比我大的?
	var leftAdd = make([]int, num)
	var rightAdd = make([]int, num)
	var cnt int
	for i := 1; i < num; i++ {
		if security[i] <= security[i-1] {
			cnt++
		} else {
			cnt = 0
		}
		leftAdd[i] = cnt
	}

	cnt = 0
	for i := num - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			cnt++
		} else {
			cnt = 0
		}
		rightAdd[i] = cnt
	}

	// 记录合法的天数
	var ret []int
	for i := 0; i < num; i++ {
		if leftAdd[i] >= time && rightAdd[i] >= time {
			ret = append(ret, i)
		}
	}
	return ret
}
