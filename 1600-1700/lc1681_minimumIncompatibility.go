//go:build ignore

package main

import (
	"math"
	"math/bits"
)

func minimumIncompatibility(nums []int, k int) int {
	n := len(nums)
	// 组的数量
	group := n / k
	inf := math.MaxInt32
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	// group对应的mask的最小不兼容性
	values := make(map[int]int)

	for mask := 1; mask < (1 << n); mask++ {
		// 依次迭代每个组合的可能, 要求对应的mask中1的数量和group的数量相同,
		// 此时这个mask就对应了一种可能的组合
		if bits.OnesCount(uint(mask)) != group {
			continue
		}
		minVal := 20
		maxVal := 0
		cur := make(map[int]bool)
		// 获取组内的各个成员
		for i := 0; i < n; i++ {
			// 要求组内的成员不能相同, 所以得进行过滤
			if mask&(1<<i) != 0 {
				if cur[nums[i]] {
					break
				}
				cur[nums[i]] = true
				minVal = min(minVal, nums[i])
				maxVal = max(maxVal, nums[i])
			}
		}
		if len(cur) == group {
			// 这个组合对应的最小不兼容性
			values[mask] = maxVal - minVal
		}
	}

	for mask := 0; mask < (1 << n); mask++ {
		// 迭代所有的mask, 这里的mask分配的可能. 0表示未分配
		// mask为0时, 表示所有元素都未分配
		if dp[mask] == inf {
			continue
		}
		seen := make(map[int]int)
		for i := 0; i < n; i++ {
			// seen中, 相等的元素保留最后一个位置
			if (mask & (1 << i)) == 0 {
				seen[nums[i]] = i
			}
		}
		// 所有未分配的数字的个数不能小于一个分组内的数字的个数
		if len(seen) < group {
			continue
		}
		// 将所有未分配的数字重新组合成sub(mask)
		sub := 0
		for _, v := range seen {
			sub |= (1 << v)
		}
		// sub是所有未分配的数字的组合, 所以是可能大于group的
		// nxt是sub的子集, 且nxt的1的个数和group相同才会被考虑
		nxt := sub
		for nxt > 0 {
			// mask保留的是当前已选的数字的集合
			// nxt 保留的是未选的集合
			// 如果nxt出现在了 values中, 说明这是一个合法的分组组合
			if val, ok := values[nxt]; ok {
				// mask | nxt 表示的是之前已经选取的组合对应的最小差值集合
				// dp[mask] + val 指的是当前匹配的组合对应的最小差值集合
				// 核心就是找到完美组合的最小值...
				dp[mask|nxt] = min(dp[mask|nxt], dp[mask]+val)
			}
			// 从大到小枚举子集
			nxt = (nxt - 1) & sub
		}
	}
	if dp[(1<<n)-1] < inf {
		return dp[(1<<n)-1]
	}
	return -1
}
