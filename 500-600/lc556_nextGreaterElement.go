package main

import (
	"fmt"
	"math"
)

func nextGreaterElement(n int) int {
	// 从后向前, 找到第一个非逆序的数字, 替换最后一个, 反转后半部分
	// 如果和原始值相同, 就返回-1

	var arr []int
	var org = n
	for n != 0 {
		arr = append(arr, n%10)
		n /= 10
	}

	if len(arr) == 0 {
		return 0
	}

	var reverse = func(arr []int) {
		for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}

	reverse(arr)

	var l int
	for l = len(arr) - 1; l > 0; l-- {
		if arr[l] > arr[l-1] {
			break
		}
	}
	fmt.Println(arr)
	if l == 0 { // 已经是最大的序列了, 这里直接返回就行
		return -1
	}

	// 从arr[l:]中找到最后一个大于arr[l-1]的值, 并进行互换
	var mi, j = arr[l-1], l
	for j < len(arr) && mi < arr[j] {
		j++
	}
	arr[l-1], arr[j-1] = arr[j-1], arr[l-1]

	fmt.Println(arr)

	// 后半段逆序
	reverse(arr[l:])

	fmt.Println(arr)

	var after int
	for _, v := range arr {
		after = after*10 + v
	}

	if after <= org || after > math.MaxInt32 {
		return -1
	}

	return after
}
