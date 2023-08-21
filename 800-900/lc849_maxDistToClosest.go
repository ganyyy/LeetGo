package main

import "math"

func maxDistToClosest(seats []int) int {
	res := 0
	l := 0
	// 去掉前缀
	for l < len(seats) && seats[l] == 0 {
		l++
	}
	res = max(res, l)
	for l < len(seats) {
		r := l + 1
		for r < len(seats) && seats[r] == 0 {
			r++
		}
		if r == len(seats) {
			// 到了末尾, 直接计算
			res = max(res, r-l-1)
		} else {
			// 没到末尾, 计算中间的距离
			res = max(res, (r-l)/2)
		}
		l = r
	}
	return res
}

func maxDistToClosest2(seats []int) int {
	// 这个做法不太好, 因为需要额外的空间
	// 左边, 右边?
	ln := len(seats)
	var right = make([]int, ln)

	var pre = math.MaxInt
	for i := ln - 1; i >= 0; i-- {
		v := seats[i]
		if v == 1 {
			pre = 0
		} else if pre != math.MaxInt {
			pre++
		}
		right[i] = pre
	}

	var ret int

	pre = math.MaxInt
	for i, v := range seats {
		if v == 1 {
			pre = 0
		} else if pre != math.MaxInt {
			pre++
		}
		ret = max(ret, min(pre, right[i]))
	}
	return ret
}
