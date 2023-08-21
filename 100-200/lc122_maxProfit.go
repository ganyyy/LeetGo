//go:build ignore

package main

func maxProfit(prices []int) int {
	m := 0
	// 核心这样的: 当天买了还能卖出去, 当天卖了还能买回来
	// 所以如果后边比前边大, 直接加上即可
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			m += prices[i+1] - prices[i]
		}
	}
	return m
}

func maxProfit122Normal(prices []int) int {
	// 前一天买了, 当天卖掉, 然后接着买, 第二天再卖掉 等同于 前一天买, 第二天卖
	// 1, 2, 3       2-1 + 3-2 = 3-1
	var res int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

// 尝试使用dp 进行处理
func maxProfit122Dp(prices []int) int {
	// dp[i] 表示第 i 天的状态, 其中
	// i[0] 表示持有股票
	// i[1] 表示未持有股票
	var dp = make([][2]int, len(prices))

	// 初始化
	dp[0][0], dp[0][1] = -prices[0], 0
	for i := 1; i < len(prices); i++ {
		// 当天持有股票, 最大收益为 昨天持有 和 昨天未持有,今天买入 的最大值
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		// 当天不持有股票, 最大收益为 昨天不持有 和 昨天持有,今天卖出 的最大值
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	// 返回最后一天不持有股票的最大收益
	return dp[len(prices)-1][1]
}

func main() {

}
