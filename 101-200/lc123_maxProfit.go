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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
