package main

import "sort"

const MAX = 1e9

var numMap = make(map[int]bool)
var numArr = make([]int, 0)

func init() {
	var a, b = 1, 1
	numMap[1] = true
	numArr = append(numArr, 1)
	for b < MAX {
		a, b = b, a+b
		numArr = append(numArr, b)
		numMap[b] = true
	}
}

func findMinFibonacciNumbers(k int) int {
	// 如果存在, 直接返回即可
	if numMap[k] {
		return 1
	}

	var cnt int

	for {
		// Golang的Sort, 有点怪.
		i := sort.SearchInts(numArr, k) - 1 // k一定不存在与 numArr中, 否则上一轮迭代就直接返回了
		cnt++
		k -= numArr[i]
		i++ // 再次在当前位置计算
		if numMap[k] {
			return cnt + 1
		}
	}

}

func main() {
	println(findMinFibonacciNumbers(4))
}
