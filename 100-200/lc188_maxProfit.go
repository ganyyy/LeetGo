package main

import "math"

func maxProfit(k int, prices []int) int {
	// mark
	if k == 0 || len(prices) == 0 {
		return 0
	}
	// 如果大于一半, 就相当于不限制出手次数的购买
	// 使用122的逻辑
	if k >= len(prices)/2 {
		var m int
		for i := 1; i < len(prices); i++ {
			if prices[i]-prices[i-1] > 0 {
				m += prices[i] - prices[i-1]
			}
		}
		return m
	}

	// 借用123的用法
	total := k * 2
	// 奇数表示第i 手买
	// i+1表示第i手卖
	ret := make([]int, total)
	for i := 0; i < total; i += 2 {
		ret[i] = math.MinInt32
	}

	for _, v := range prices {
		// 第一手买, 取历史最小和当前最小的最小值
		ret[0] = max(ret[0], -v)
		// 第一手卖, 取历史最大值和当前卖的最大值
		ret[1] = max(ret[1], v+ret[0])
		for i := 2; i < total; i += 2 {
			// 上一手卖出的价格-当前买入的价格
			ret[i] = max(ret[i], ret[i-1]-v)
			// 当前手买的价格+卖出的价格
			ret[i+1] = max(ret[i+1], ret[i]+v)
		}
	}

	// 返回最后结果
	return ret[total-1]
}

func main() {

}
