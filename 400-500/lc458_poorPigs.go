package main

import "math"

// 解法查阅 https://leetcode-cn.com/problems/poor-pigs/solution/gong-shui-san-xie-jin-zhi-cai-xiang-xian-69fl/

func poorPigs(buckets, minutesToDie, minutesToTest int) int {

	// 通过 minutesToTest/minutesToDie 可以确定轮数
	// 本质上相当于获取到  以 (minutesToTest/minutesToDie+1) 为进制基数
	// 判断bucket可以使用多少位来进行表示

	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(minutesToTest/minutesToDie+1))))
}
