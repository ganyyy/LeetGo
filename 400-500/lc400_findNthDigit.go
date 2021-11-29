package main

import "math"

func findNthDigit(n int) int {
	d := 1
	// 首先计算d的位数
	for count := 9; n > d*count; count *= 10 {
		n -= d * count
		d++
	}
	// 计算n处在第几位
	index := n - 1
	// 计算d位数的起始位置
	start := int(math.Pow10(d - 1))
	// 获取n对应的数字
	num := start + index/d
	// n在所处的数字中的第几位
	digitIndex := index % d
	// 真尼玛难想
	// 使用 n 对应的数字, 除以 剩余的位数, 就等于 n的index位上的值
	return num / int(math.Pow10(d-digitIndex-1)) % 10
}
