package main

func maxScore(cardPoints []int, k int) int {

	var ln = len(cardPoints)
	// 求所有的和, 然后维护一个长度为ln-k的窗口, 保证窗口的和是最小的
	var sum int
	for _, v := range cardPoints {
		sum += v
	}
	var cur int
	for i := 0; i < ln-k; i++ {
		cur += cardPoints[i]
	}
	var mmin = cur
	for i := ln - k; i < len(cardPoints); i++ {
		cur += cardPoints[i] - cardPoints[i+k-ln]
		mmin = min(cur, mmin)
	}
	// 最后用总和减去这个值就行
	return sum - mmin
}
