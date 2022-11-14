package main

func maxProfit(prices []int) int {
	// 标记
	ln := len(prices)
	if ln == 0 {
		return 0
	}
	// 第i天售卖的最大收益 要么前一天买了今天卖, 要么不卖 sell[i] = max(buy[i-1]+prices[i], sell[i-1])
	sell := make([]int, ln)
	// 第i天购买的最大收益  要么前一天冻着今天买, 今天不买 buy[i] = max(cool[i-1]-prices[i], buy[i-1])
	buy := make([]int, ln)
	// 第i天冻结的最大收益 要么继续冻着, 要么买一下试试 cool[i] = max(sell[i-1], cool[i-1])
	cool := make([]int, ln)

	buy[0] = -prices[0]

	for i := 1; i < ln; i++ {
		sell[i] = max(buy[i-1]+prices[i], sell[i-1])
		buy[i] = max(cool[i-1]-prices[i], buy[i-1])
		cool[i] = max(sell[i-1], cool[i-1])
	}

	return sell[ln-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
