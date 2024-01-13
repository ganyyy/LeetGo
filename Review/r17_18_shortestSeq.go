package main

import (
	"math"
	"strings"
)

func shortestSeq(big []int, small []int) []int {
	var count = len(small)
	var smallMap = make(map[int]int, count)

	strings.TrimPrefix("abc", "a")

	for _, sn := range small {
		smallMap[sn]++
	}
	var minStart = math.MaxInt32
	var minLength = math.MaxInt32
	var start int
	var curLength int
	for end, val := range big {
		if num, ok := smallMap[val]; ok {
			num--
			if num == 0 {
				curLength++
			}
			smallMap[val] = num
		} else if !ok && curLength == 0 {
			start = end + 1
		}
		if curLength == count {
			for num, ok := smallMap[big[start]]; !ok || num < 0; num, ok = smallMap[big[start]] {
				// 如果
				// 不存在的值, 或者超标的值都可以跳过去
				if ok {
					// 窗口中有多余的值, 这个可以跳过
					smallMap[big[start]]++
				}
				start++
			}
			// 满足条件了
			// 从窗口中剔除一个数字
			if length := end - start + 1; length < minLength {
				minStart = start
				minLength = length
			}
			smallMap[big[start]]++
			curLength--
			for start++; start <= end; start++ {
				if _, ok := smallMap[big[start]]; ok {
					break
				}
			}
		}
	}
	if minStart == math.MaxInt32 {
		return nil
	}
	return []int{minStart, minStart + minLength - 1}
}
