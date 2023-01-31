package main

func maxHappyGroups(batchSize int, groups []int) (ans int) {
	const kWidth = 5
	const kWidthMask = 1<<kWidth - 1

	var max = func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	cnt := make([]int, batchSize)
	for _, x := range groups {
		cnt[x%batchSize]++
	}

	/*
	   1 <= batchSize <= 9
	   1 <= groups.length <= 30
	   1 <= groups[i] <= 10e9

	*/

	// 每5bit([0, 31])表示一个batch中所有元素的数量, 上限是 5*9 = 45bit, 可以使用一个64位数完全表示

	start := 0
	// 因为cnt[0]天然就是满意的, 所以在最后累加, 这里不需要计数
	// 这里按照余数从大到小的逆序累计
	for i := batchSize - 1; i > 0; i-- {
		start = start<<kWidth | cnt[i]
	}

	// 每个mask对应的计算缓存
	memo := map[int]int{}
	var dfs func(int) int
	dfs = func(mask int) (best int) {
		if mask == 0 {
			return
		}
		if res, ok := memo[mask]; ok {
			return res
		}

		// 统计当前mask剩余的总个数
		total := 0
		for i := 1; i < batchSize; i++ {
			amount := mask >> ((i - 1) * kWidth) & kWidthMask
			total += i * amount
		}

		for i := 1; i < batchSize; i++ {
			// 按位置挨个扣
			amount := mask >> ((i - 1) * kWidth) & kWidthMask
			if amount > 0 {
				// 计算剩余数量的满意度
				result := dfs(mask - 1<<((i-1)*kWidth))
				if (total-i)%batchSize == 0 {
					// 如果去掉这个后, 剩余的也是满意的?
					result++
				}
				best = max(best, result)
			}
		}
		memo[mask] = best
		return
	}
	return dfs(start) + cnt[0]
}
