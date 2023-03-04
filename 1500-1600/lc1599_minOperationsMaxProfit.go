package main

func minOperationsMaxProfit(customers []int, boardingCost, runningCost int) int {

	// 如果四人带来的收益还不如启动的开销, 则直接跳过即可
	profitEachTime := boardingCost*4 - runningCost
	if profitEachTime <= 0 {
		return -1
	}

	ans := -1

	var maxProfit int      // 最大利润
	var totalProfit int    // 当前利润
	var operations int     // 操作次数
	var customersCount int // 乘客的数量
	// 这一步, 是统计所有人都开始排队时, 带来的最大收益
	// 为什么要这么做呢? 因为刚开始是不一定可以坐满的, 可能存在空缺
	for _, c := range customers {
		operations++                                           // 轮转
		customersCount += c                                    // 总人数
		curCustomers := min(customersCount, 4)                 // 本次可以上车的人数
		customersCount -= curCustomers                         // 剩余的人数
		totalProfit += boardingCost*curCustomers - runningCost // 总利润 = 之前的总利润+本次轮转的利润-一次轮转的开销
		if totalProfit > maxProfit {
			maxProfit = totalProfit // 更新最大利润和最小的次数
			ans = operations
		}
	}
	if customersCount == 0 { // 没有剩余乘客了
		return ans
	}
	if customersCount > 0 {
		// 当前剩余的人数
		fullTimes := customersCount / 4           // 四人一组, 上座
		totalProfit += profitEachTime * fullTimes // 整组带来的收益
		operations += fullTimes
		if totalProfit > maxProfit {
			maxProfit = totalProfit
			ans = operations
		}
		// 判断一下, 最后的 [0,3]人带来的收益
		remainingCustomers := customersCount % 4
		remainingProfit := boardingCost*remainingCustomers - runningCost
		totalProfit += remainingProfit
		if totalProfit > maxProfit {
			maxProfit = totalProfit
			operations++
			ans++
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
