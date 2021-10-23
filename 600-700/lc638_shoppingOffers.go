package main

func shoppingOffers(price []int, special [][]int, needs []int) int {

	// 总的商品计数
	var n = len(price)

	// 过滤不符合要求的礼包
	var filterIdx = -1
	for i, s := range special {
		// 计算礼包的总价值
		var totalCount, totalPrice int
		for idx, cnt := range s[:n] {
			totalCount += cnt
			totalPrice += price[idx] * cnt
		}

		// 判断礼包是否合法: 存在可用的数量, 并且礼包内商品单独购买的价值大于礼包的价值
		if totalCount > 0 && totalPrice > s[n] {
			filterIdx++
			special[filterIdx], special[i] = special[i], special[filterIdx]
		}
	}
	special = special[:filterIdx+1]

	// 记忆化搜索. 记录一种特定组合格式的最低价格
	var search = make(map[string]int)

	// 基于当前的组合, 获取最小值价格
	var dfs func(curNeeds []byte) (minPrice int)

	dfs = func(curNeeds []byte) (minPrice int) {
		// 如果已有缓存, 直接返回结果
		if res, ok := search[string(curNeeds)]; ok {
			return res
		}

		// 计算不购买礼包的价值
		for i, p := range price {
			minPrice += int(curNeeds[i]) * p // 不购买任何大礼包，原价购买购物清单中的所有物品
		}

		var nextNeeds = make([]byte, n)
	outer:
		// 计算购买礼包所能降低的价格
		for _, s := range special {
			for idx, cnt := range curNeeds {
				// 不允许总数量大于需要的数量
				if cnt < byte(s[idx]) {
					continue outer
				}
				nextNeeds[idx] = cnt - byte(s[idx])
			}
			// 更新一下当前购买方案的最低价格
			minPrice = min(minPrice, dfs(nextNeeds)+s[n])
		}

		// 缓存计算结果
		search[string(curNeeds)] = minPrice
		return
	}

	var curNeeds = make([]byte, n)
	for i := range needs {
		curNeeds[i] = byte(needs[i])
	}

	return dfs(curNeeds)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
