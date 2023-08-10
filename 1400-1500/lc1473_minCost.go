package main

import "math"

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	// 成吧, 果然三重DP

	// dp[i][j][k] 表示i(i ∈ [0, m))个房子涂成第j(j ∈ [1, n])个颜色, 当前处理时的街区数量为k(k ∈ [1:target])个 时的消耗

	// 初始化dp数组
	var dp = make([][][]int, m)
	for i := 0; i < m; i++ {
		var color = make([][]int, n+1)
		for i2 := 0; i2 < n+1; i2++ {
			var basket = make([]int, target+1)
			for i3 := 0; i3 < target+1; i3++ {
				basket[i3] = math.MaxInt32
			}
			color[i2] = basket
		}
		dp[i] = color
	}

	// DP 初始化

	// 第一个房子只能属于第一个街区
	if houses[0] != 0 {
		// 如果首个房子已经有颜色了, 那么第一个房子对应颜色的花费为0
		dp[0][houses[0]][1] = 0
	} else {
		var color = dp[0]
		var costs = cost[0]
		for i := 1; i <= n; i++ {
			// 每种颜色对应的开销进行初始化
			// 第一个房子只能有一个街区
			color[i][1] = costs[i-1]
		}
	}

	// 实际上, 这里只用到了前一个房子的状态, 所以可以进行DP的压缩

	for i := 1; i < m; i++ {
		if houses[i] == 0 {
			// 自由填充
			for i2 := 1; i2 <= n; i2++ {
				for i3 := 1; i3 <= target; i3++ {
					// 街区的数量肯定不能大于当前房子的数量
					if i3 > i {
						break
					}
					// 如果前一个房子不存在i2颜色, i3个街区的组合, 那就跳过
					// 本质上, i-1可能是涂过颜色的, 也可能是没涂过颜色的. 如果是涂过颜色的,
					// 那么就只有前一个有效的颜色才有用, 剩下的都是无效的
					if dp[i-1][i2][i3] == math.MaxInt32 {
						continue
					}
					// i2 代表的是前一个房子的颜色, i22代表的是当前房子的颜色
					for i22 := 1; i22 <= n; i22++ {
						if i22 == i2 {
							// 如果当前房子和前一个房子颜色相同, 那么就只能处于同一个街区
							dp[i][i22][i3] = min(dp[i][i22][i3], dp[i-1][i2][i3]+cost[i][i22-1])
						} else if i3 != target {
							// 如果和前一个颜色不相同, 那么就算是不同的街区
							dp[i][i22][i3+1] = min(dp[i][i22][i3+1], dp[i-1][i2][i3]+cost[i][i22-1])
						}
					}
				}
			}
		} else {
			// 根据前后环境进行填充
			var c = houses[i]
			for i2 := 1; i2 <= n; i2++ {
				for i3 := 1; i3 <= target; i3++ {
					// 街区的数量肯定不能大于当前房子的数量
					if i3 > i {
						break
					}
					if dp[i-1][i2][i3] == math.MaxInt32 {
						continue
					}
					if c == i2 {
						// 同一种颜色的情况下, 没有形成新的分区, 已经涂完的房子不会产生新的消耗,
						// 所以当前消耗和前一个房间对应颜色, 街区消耗的最小值即可
						dp[i][c][i3] = min(dp[i][c][i3], dp[i-1][i2][i3])
					} else if i3 != target {
						// 产生了新的街区, 但是同样也不会存在新的消耗.  依旧和前一个房子指定颜色的街区做比较
						dp[i][c][i3+1] = min(dp[i][c][i3+1], dp[i-1][i2][i3])
					}
				}
			}
		}
	}

	// 计算最后的结果
	var res = math.MaxInt32
	for i2 := 1; i2 <= n; i2++ {
		res = min(res, dp[m-1][i2][target])
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func minCostDaLao(houses []int, cost [][]int, m int, n int, target int) int {
	dp0 := make([][]int, m+1)
	dp1 := make([][]int, m+1)
	// dp[i][j] 表示的是第i个房子, 颜色为j 时, 当前街区数量对应的最低消耗
	for i := 0; i <= m; i++ {
		dp0[i] = make([]int, n+1)
		dp1[i] = make([]int, n+1)
	}
	maxVal := math.MaxInt32
	resetMat(dp0, maxVal)
	for i := 0; i <= n; i++ {
		dp0[0][i] = 0
	}
	for k := 1; k <= target; k++ {
		// 重置dp1, 用来计算当前的状态转移
		resetMat(dp1, maxVal)
		// 当前的房子
		// k个街区最少k个房子
		for i := k; i <= m; i++ {
			// 首先计算一下上一个街区时, 前一个房间最小的开销和其对应的颜色, 以及第二小的颜色的开销
			minC1 := 1           // 最小开销对应的颜色
			minV1 := dp0[i-1][1] // 第一小的开销
			minV2 := maxVal      // 第二小的开销
			for j := 2; j <= n; j++ {
				if minV1 > dp0[i-1][j] {
					minV2 = minV1
					minV1 = dp0[i-1][j]
					minC1 = j
				} else if minV2 > dp0[i-1][j] {
					minV2 = dp0[i-1][j]
				}
			}
			// 一定要创建一个新的街区,
			cc := houses[i-1]
			if cc != 0 {
				// 当前房子已经涂过了颜色, 所以不会产生新的开销
				if cc == minC1 {
					// 那么当前颜色就不能和上一个街区的最小开销颜色一致
					// 所以这里计算最小值是和第二小的开销进行比较
					dp1[i][cc] = min(dp1[i-1][cc], minV2)
				} else {
					// 颜色和上一个街区数量的最小街区不一致, 和最小的开销进行比较
					dp1[i][cc] = min(dp1[i-1][cc], minV1)
				}
				continue
			}
			for j := 1; j <= n; j++ {
				// 同样的道理, 选取的颜色和上一个街区最小消耗颜色不一致的情况下, 取最小值
				// 一致的情况下取第二小的值
				mj := minV1
				if minC1 == j {
					mj = minV2
				}
				dp1[i][j] = min(dp1[i-1][j], mj)
				// 如果获取的的是最大值, 没必要加了, 当前颜色在上一个街区不存在可使用的情况
				if dp1[i][j] < maxVal {
					// 增加该涂装颜色对应的开销
					dp1[i][j] += cost[i-1][j-1]
				}
			}
		}
		// 交换一下. 对于下一个街区而言, dp1就是dp0
		dp0, dp1 = dp1, dp0
	}

	// 计算最小的消耗
	minV := maxVal
	for j := 1; j <= n; j++ {
		if minV > dp0[m][j] {
			minV = dp0[m][j]
		}
	}
	if minV == maxVal {
		return -1
	}
	return minV
}

func resetMat(dp [][]int, v int) {
	m := len(dp)
	n := len(dp[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = v
		}
	}
}
