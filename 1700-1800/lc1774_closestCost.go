//go:build ignore

package main

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	x := baseCosts[0]
	// 最小值大于target
	// 无需计算下一步
	for _, c := range baseCosts {
		x = min(x, c)
	}
	//提前枝减恰好相等的情况(不加任何配料)
	if x >= target {
		return x
	}

	// 大于target的可能采用贪心
	// 小于target的可能采用01背包

	// 标记可以增加辅料的base
	// can[i]表示存在一个到达i的路径,
	can := make([]bool, target+1)
	// ans表示的是大于target中的最小值
	ans := 2*target - x // x一定小于target, ans一定大于target
	for _, c := range baseCosts {
		if c <= target {
			can[c] = true
		} else {
			// 从不能添加辅料的base中, 找到最小值
			ans = min(ans, c)
		}
	}
	for _, c := range toppingCosts {
		// 统计所有辅料
		// 怎么体现出count两次的?
		// 零一背包的思路, 通过累计标记可达路径
		for count := 0; count < 2; count++ {
			// 第一次为 i = XXX-2c 的时候有效, 这个结果可以在下一次迭代c的时候生效
			// 第二次就是在 i = XXX-c的时候生效了
			// 最多增加两次
			// 如果小于target的更为接近的话(?)
			for i := target; i > 0; i-- {
				// 迭代每一种可能,
				// 大于target更新min
				// 小于target更新路径
				if can[i] && i+c > target {
					// 如果可以加, 并且超过了target
					ans = min(ans, i+c)
				}
				if i-c <= 0 {
					continue
				}
				can[i] = can[i] || can[i-c]
			}
		}
	}
	// 这里就隐含的处理了 ans == target的情况
	// 这里隐含的就是 ans是距离target最近且大于target的最大值, 二者的距离是dis
	// 那么迭代[0,dis]中的所有距离(target,target-1,target-2,target-3...,target-dis)
	// 如果存在一条可达路径, 那么它一定是最接近target且小于target的最大值
	for i := 0; i <= ans-target; i++ {
		// 处理小于target的情况(?)
		// 所有小于target中的第一个一定是总消耗最少的
		if can[target-i] {
			return target - i
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
