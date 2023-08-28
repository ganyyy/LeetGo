//go:build ignore

package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	// mark
	if len(prices) < 2 {
		return 0
	}
	// 最小值索引和最大利润
	mi, max := prices[0], 0
	for i := 1; i < len(prices); i++ {
		// 如果利润最大还大, 就更新一下
		if max < prices[i]-mi {
			max = prices[i] - mi
		}
		// 如果当前值比最小值还小, 就更新一下
		if prices[i] < mi {
			mi = prices[i]
		}
	}
	return max
}

func maxProfit2(prices []int) int {
	buy, sell := math.MinInt32, 0
	for _, v := range prices {
		// 买入
		buy = max(buy, -v)
		// 卖出
		sell = max(sell, v+buy)
	}
	return sell
}

func main() {
	fmt.Println(maxProfit2([]int{7, 6, 4, 3, 1}))
}
