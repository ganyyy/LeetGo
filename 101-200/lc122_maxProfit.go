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

func main() {

}
