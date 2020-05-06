package main

func mincostTickets(days []int, costs []int) int {
	// dp, 分别取 dp[i-1]+cost[0], dp[i-7]+cost[1], dp[i-30]+cost[2]的最小值
	dp := [366]int{} // 从1开始
	for _, v := range days {
		// 标记当天去旅行
		dp[v] = -1
	}

	for i := 1; i <= 365; i++ {
		if 0 == dp[i] {
			dp[i] = dp[i-1]
		} else {
			// 一天的值
			var min = dp[i-1] + costs[0]
			// 七天的最小值
			if i-7 >= 0 {
				min = getMin(min, dp[i-7]+costs[1])
			} else {
				min = getMin(min, costs[1])
			}
			// 30天的最小值
			if i-30 >= 0 {
				min = getMin(min, dp[i-30]+costs[2])
			} else {
				min = getMin(min, costs[2])
			}
			dp[i] = min
		}
	}
	return dp[365]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
