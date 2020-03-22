package main

import (
	"fmt"
	"sort"
)

func minIncrementForUnique(A []int) int {
	// 需要先排序, 在以最小值从头到尾加一遍?
	sort.Ints(A)
	res := 0
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			sub := A[i-1] - A[i] + 1
			A[i] += sub
			res += sub
		}
	}
	return res
}

// 线性探测+路径压缩, 不需要进行排序
func minIncrementForUnique2(A []int) int {
	// 这里需要申请很大一块内存, 防止溢出
	// 感觉不如排序? 要的内存太多了...
	// path[i] 保存的是寻址得到的空位
	path := make([]int, 40001)

	var findPath func(int) int
	findPath = func(a int) int {
		// 先看一下当前位置有没有人占用
		b := path[a]
		if b == 0 {
			path[a] = a
			return a
		}
		// 如果被占了, 就去找当前位置的下一个空位
		b = findPath(b + 1)
		// 更新当前值对应的位置
		// 因为是递归调用, 所以一旦出现重复值, 都会将其挨个更新
		path[a] = b
		return b
	}
	res := 0
	for _, v := range A {
		res += findPath(v+1) - v - 1
	}
	return res
}

func main() {
	fmt.Println(minIncrementForUnique2([]int{0, 0}))
}
