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
	// 第i天冻结的最大收益 要么继续冻着, 要么卖一下试试 cool[i] = max(sell[i-1], cool[i-1])
	cool := make([]int, ln)

	buy[0] = -prices[0]

	for i := 1; i < ln; i++ {
		sell[i] = max(buy[i-1]+prices[i], sell[i-1])
		buy[i] = max(cool[i-1]-prices[i], buy[i-1])
		cool[i] = max(sell[i-1], cool[i-1])
	}

	return sell[ln-1]
}

func maxProfit2(prices []int) int {
	var ln = len(prices)
	if ln == 0 {
		return 0
	}

	// s: 卖
	// b: 买
	// c: 冻结期
	var s1, s2 int
	var b1, b2 int
	var c1, c2 int

	b1 = -prices[0]

	for _, v := range prices[1:] {
		s2 = max(b1+v, s1) //
		b2 = max(c1-v, b1) // 冻结期后才能买
		c2 = max(s1, c1)   // 前一天买不买都可以进入冻结期

		s1 = s2
		b1 = b2
		c1 = c2
	}

	return s2
}
