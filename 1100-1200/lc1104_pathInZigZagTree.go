package main

import "math/bits"

//pathInZigZagTreeLLL 这是我的做法
func pathInZigZagTreeLLL(label int) []int {
	// 这是一颗满二叉树

	// 所以当 label >

	// 1 判断所处的层数

	var i int
	for (1<<(i+1))-1 < label {
		i++
	}

	var ret = make([]int, i+1)

	// 2 根据层数获取原始值和实际值
	var tVal = getTVal(i, label)
	for t := i; t >= 0; t-- {
		ret[t] = getTVal(t, tVal)
		tVal /= 2
	}
	return ret
}

func getTVal(i int, val int) int {
	// 判断奇数层还是偶数层
	if i&1 != 0 {
		// 奇数, 那么对应的层数是偶数
		val = (1 << i) + (1 << (i + 1)) - 1 - val
	}
	return val
}

func pathInZigZagTree(label int) []int {
	var res []int

	for label != 1 {
		res = append(res, label)
		label >>= 1
		l := bits.Len(uint(label))
		label ^= (1 << (l - 1)) - 1
	}
	res = append(res, 1)
	reverse(res)
	return res
}

func reverse(nums []int) {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
