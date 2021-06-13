package main

import (
	"math"
	"strconv"
)

func largestNumber(cost []int, target int) string {
	var max = func(s1, s2 string) (ret string) {
		if len(s1) != len(s2) {
			if len(s1) > len(s2) {
				return s1
			}
			return s2
		}
		if s1 > s2 {
			return s1
		}
		return s2
	}

	// 当成一个背包
	// 感觉更像是一个贪心?

	// 每次都选一个开销最小的, 数值最大的
	// 但是不一定可以正好凑成target
	var dp = make([]string, target+1)

	// 要求是整好组合, 所以只能从所有出现的数字中查询
	for i := 1; i < len(dp); i++ {
		dp[i] = "#"
	}

	// 路径的计算保留问题?
	for i := 1; i <= target; i++ {
		for k, v := range cost {
			if i >= v && dp[i-v] != "#" {
				dp[i] = max(dp[i], dp[i-v]+strconv.Itoa(k+1))
			}
		}
	}

	// 这种方法可以求出到达target 时最多能存在的个数. 现在就要回推最大值
	// 位数最多的一定是结果最大的, 区别就在组合上

	if dp[target] == "#" {
		return "0"
	}
	return dp[target]
}

func largestNumber2(cost []int, target int) string {
	// dp[i] 表示恰好消耗成本 == i时可以保留的最大位数

	var dp = make([]int, target+1)

	for i := range dp {
		dp[i] = math.MinInt32
	}
	// 消耗为0没有任何位数, 因为 0<cost[i]
	dp[0] = 0

	for _, v := range cost {
		// 后态依赖于前态, 所以前序遍历
		for j := v; j <= target; j++ {
			dp[j] = max(dp[j], dp[j-v]+1)
		}
	}

	// 没有正好为 target 的解决方案
	if dp[target] < 0 {
		return "0"
	}

	var ans = make([]byte, 0, dp[target])

	// 路径回溯

	// i 表示当前遍历到哪一位上了(一共9位)
	// j 当前剩余的可以资源数
	for i, j := 8, target; i >= 0; i-- {
		// c 表示i对应的消耗, 如果剩余的资源够用并且满足其前置条件相等, 就可以取该值
		for c := cost[i]; j >= c && dp[j] == dp[j-c]+1; j -= c {
			ans = append(ans, byte(i)+'1')
		}
	}

	return string(ans)
}

func main() {
	println(largestNumber2([]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 9))
}
