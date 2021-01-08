package main

import "math"

func maxProfit(prices []int) int {
	// 每次更新4个值, 原则上允许一天买卖买卖
	var fstSell, secSell int
	fstBuy, secBuy := math.MinInt32, math.MinInt32
	for _, v := range prices {
		// 第一手买
		fstBuy = max(fstBuy, -v)
		// 第一手卖
		fstSell = max(fstSell, fstBuy+v)
		// 第二手买
		secBuy = max(secBuy, fstSell-v)
		// 第二手卖
		secSell = max(secSell, secBuy+v)
	}
	return secSell
}

func maxProfit3(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	// 0代表买, 1代表卖
	var first, second [2]int

	first[0] = -prices[0]
	first[1] = 0
	second = first

	for i := 1; i < len(prices); i++ {
		first[0] = max(first[0], -prices[i])
		first[1] = max(first[1], first[0]+prices[i])

		second[0] = max(second[0], first[1]-prices[i])
		second[1] = max(second[1], second[0]+prices[i])
	}

	return second[1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
