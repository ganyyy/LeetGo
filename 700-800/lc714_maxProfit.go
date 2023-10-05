package main

func maxProfit(prices []int, fee int) int {
	// 标记
	if len(prices) == 0 {
		return 0
	}
	// dp[i][0] 表示第i天持有股票的利润, dp[i][1]表示第i天不持有股票的利润
	// var dp = make([][2]int, len(prices))

	// 第一天买股票
	// dp[0][0] = -prices[0]
	// 状态转移方程:
	// 第i天持有股票的最大利润为 继续持有或者前一天不持有, 今天购买的最大值
	// 第i天不持有股票的最大利润为 继续不持有或者前一天持有,今天卖出的最大值2

	// for i := 1; i < len(prices); i++ {
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
	// dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0]-fee)
	// }

	// 很明显, 当前状态只和上一个状态相关, 可以执行压缩
	// 这里的状态是不持有和持有. 和买卖无关
	// 不持有: 继续不持有, 昨天持有今天卖出
	// 持有: 继续持有, 昨天不持有今天买入
	var cash, hold = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		// 因为可以无限次的交易, 所以当天先卖或者先买都行
		hold = max(hold, cash-prices[i])
		cash = max(cash, prices[i]+hold-fee) // 卖出的时候扣手续费
	}
	// return dp[len(prices)-1][1]
	return cash
}
