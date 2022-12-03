//go:build ignore

package main

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	x := baseCosts[0]
	// 最小值大于target
	// 无需计算下一步
	for _, c := range baseCosts {
		x = min(x, c)
	}
	if x >= target {
		return x
	}

	// 标记可以增加辅料的base
	// can[i]表示存在一个到达i的路径
	can := make([]bool, target+1)
	ans := 2*target - x // 一个相对较大值
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
		for count := 0; count < 2; count++ {
			// 最多增加两次
			// 如果小于target的更为接近的话(?)
			for i := target; i > 0; i-- {
				// 迭代每一种可能,
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
