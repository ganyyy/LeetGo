package main

import "strconv"

func maximumSwap(num int) int {

	s := []byte(strconv.Itoa(num))
	n := len(s)
	maxIdx, idx1, idx2 := n-1, -1, -1
	for i := n - 1; i >= 0; i-- {
		if s[i] > s[maxIdx] {
			maxIdx = i
		} else if s[i] < s[maxIdx] {
			// maxIdx是最大值对应的位置
			// idx1是小于[maxIdx]的最高位
			idx1, idx2 = i, maxIdx
		}
	}
	if idx1 < 0 {
		// 所有的高位都大于等于低位, 无需交换
		return num
	}
	s[idx1], s[idx2] = s[idx2], s[idx1]
	v, _ := strconv.Atoi(string(s))
	return v
}
