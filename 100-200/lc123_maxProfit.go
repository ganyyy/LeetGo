//go:build ignore

package main

import "math"

func maxProfit(prices []int) int {
	// mark
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
	// Mark

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

func maxProfit4(prices []int) int {
	const (
		BUY  = 0
		SELL = 1
	)
	// 一天是可以同时买卖两次的. 所以第一天二者是一样的
	var first, second [2]int
	first[BUY] = -prices[0]
	second = first

	for _, price := range prices[1:] {
		// 第一手购买的收益
		first[BUY] = max(first[BUY], -price)
		// 第一手卖出的收益
		first[SELL] = max(first[SELL], price+first[BUY])
		// 第二手购买
		second[BUY] = max(second[BUY], first[SELL]-price)
		// 第二手卖出
		second[SELL] = max(second[SELL], second[BUY]+price)
	}

	return second[SELL]
}

func main() {

}
