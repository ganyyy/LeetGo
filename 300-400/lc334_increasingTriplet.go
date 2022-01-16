package main

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	// 三个数, 当长度不确定时, 就需要通过二分查找定长数组来处理了
	// 经典LIS

	var src = make([]int, 0, 3)
next:
	for _, v := range nums {
		if len(src) == 0 {
			src = append(src, v)
			continue
		}
		for j, v2 := range src {
			if v <= v2 {
				src[j] = v
				continue next
			}
		}
		src = append(src, v)
		if len(src) == 3 {
			return true
		}
	}
	return false
}

func increasingTriplet2(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	var a, b = math.MaxInt32, math.MaxInt32

	for _, v := range nums {
		if a >= v {
			a = v
		} else if b >= v {
			b = v
		} else {
			return true
		}
	}
	return false
}
