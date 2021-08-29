package main

func sumOddLengthSubarrays(arr []int) int {
	// 统计每一个位置上在奇数子数组上出现的次数

	// 整体分为两部分.
	// 左边取奇数个数字, 右边也取奇数个数字, 加上自己 = 奇数
	// 左边取偶数个数字, 右边也取偶数个数字, 加上自己 = 奇数

	// 对于下标为i的数而言, 左边有(i+1)个数, 右边有(len-i)个数
	// 取奇数个数字的取法有 num/2 ([0,1,2] => 1)
	// 取偶数个数字的取法有 (num+1)/2 ([0,1,2] => 0, 2)

	var ln = len(arr)
	var ret int
	for i := 0; i < ln; i++ {
		var left, right = i + 1, ln - i
		var lo, ro, le, re = left / 2, right / 2, (left + 1) / 2, (right + 1) / 2
		ret += (lo*ro + le*re) * arr[i]
	}
	return ret
}

func sumOddLengthSubarrays2(arr []int) int {
	var sum = make([]int, len(arr)+1)
	sum[1] = arr[0]
	for i := 1; i < len(arr); i++ {
		sum[i+1] = sum[i] + arr[i]
	}

	var ret int
	for i := 1; i <= len(arr); i += 2 {
		for j := i; j <= len(arr); j += 1 {
			ret += sum[j] - sum[j-i]
		}
	}

	return ret
}
