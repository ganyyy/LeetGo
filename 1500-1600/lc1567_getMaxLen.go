package main

func getMaxLen(nums []int) int {
	var mi, ma int
	var ret int

	// 正数正常计算
	// 遇到0重置
	// 负数交换位置
	//

	for _, v := range nums {
		if v == 0 {
			mi, ma = 0, 0
		} else if v > 0 {
			ma++
			if mi > 0 {
				mi++
			}
		} else {
			mi, ma = ma, mi
			mi++
			if ma > 0 {
				ma++
			}
		}
		ret = max(ret, ma)
	}
	return ret
}
