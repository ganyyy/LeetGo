//go:build ignore

package main

import "math"

const (
	T = 243 // 3^5, 范围上限
	N = 5   // 列上限
	M = 6   // 内向/外向人数上限
)

// Score 相邻状态为(x,y)
// 类似于三进制, 每个点位有三种状态
// 0: 没人
// 1: 内向的人
// 2: 外向的人
// (1,1)额外得分-60(-30,-30)
// (1,2)额外得分-10(-30,+20)
// 以此类推
var Score = [3][3]int{
	{0, 0, 0},
	{0, -60, -10},
	{0, -10, 40},
}

// 针对每一行而言, 每个mask对应的指定列里的人是内向的还是外向的
var maskBits [T][N]int

// 针对每一行而言, 每个mask状态对应的已选取的内向人数
var ivCount [T]int

// 针对每一行而言, 每个mask状态对应的已选取的外向人数
var evCount [T]int

// 针对每一行而言, 每个状态对应的行内得分
var innerScore [T]int

// 针对每两行而言, 前置行和当前行每个mask加一块的分值
var interScore [T][T]int

// D 第N-1行的状态为T, 剩下的行还可以放置IV个内向和EV个外向的最大得分
// N的上限是5, T的上限是pow(3, 5) = 243
// IV/EV的上限是6, 所以总和就是[0, 6]一共7个状态
var D [N][T][M + 1][M + 1]int

func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
	// 每个位置有三种状态, 每一行有n列, 将其状态进行编码, 则tot = pow(3, n)
	// 我滴个乖乖, 这个状态压缩的方法太妙了
	var tot = int(math.Pow(3, float64(n)))

	initData := func() {

		// 初始化D: -1表示未赋值
		for i := 0; i < N; i++ {
			for j := 0; j < T; j++ {
				for k := 0; k <= M; k++ {
					for l := 0; l <= M; l++ {
						D[i][j][k][l] = -1
					}
				}
			}
		}

		{
			maskBits = [T][N]int{}
			ivCount = [T]int{}
			evCount = [T]int{}
			innerScore = [T]int{}
			interScore = [T][T]int{}
		}

		// 计算行内分数
		for mask := 0; mask < tot; mask++ {
			// 这个方法妙啊
			// tot 上限是: 3 * 3 * 3 * 3 * 3
			// 对于每个mask, 表示的是一行5个位置的状态
			// 比如如果是1的话, 切换到5行就是 1, 0, 0, 0, 0 相当于首个位置放置了一个内向人
			// 121 表示的是 1, 1, 1, 1, 1 相当于每个位置都放置了一个内向人
			// 242 表示的是 2, 2, 2, 2, 2 相当于每个位置都放置了一个外向人
			var tmpMask = mask
			for i := 0; i < n; i++ {
				x := tmpMask % 3      // 取余得到的是这个位置的状态(0, 1, 2)
				maskBits[mask][i] = x // 给当前mask对应的当前列存储的人物特性进行赋值
				tmpMask /= 3          // 整除是切换到下一个位置
				if x == 1 {
					ivCount[mask]++         // 当前mask内向人数++
					innerScore[mask] += 120 // 行内分数++
				} else if x == 2 {
					evCount[mask]++        // 当前mask外向人数++
					innerScore[mask] += 40 // 行内分数++
				}
				if i > 0 {
					innerScore[mask] += Score[x][maskBits[mask][i-1]] // 同一行内两个相邻的格子, 额外的分数
				}
			}
		}
		// 计算行间分数
		for i := 0; i < tot; i++ {
			for j := 0; j < tot; j++ {
				interScore[i][j] = 0
				for k := 0; k < n; k++ {
					// i代表的是上一行的某个mask, j代表的是当前行的某个mask
					// (其实反过来也没啥问题, 毕竟俩长度是一样的)
					// 行间分数相当于竖向相邻的两个格子, 额外的分数
					interScore[i][j] += Score[maskBits[i][k]][maskBits[j][k]]
				}
			}
		}
	}

	// 初始化
	initData()

	var dfs func(row, preMask, iv, ev int) int

	dfs = func(row, preMask, iv, ev int) int {
		// 到达了行尾, 或者没有剩余的可放置的人
		if row == m || (iv == 0 && ev == 0) {
			return 0
		}
		p := &D[row][preMask][iv][ev]
		if val := *p; val != -1 {
			return val
		}
		// 初始状态
		var cur int
		for mask := 0; mask < tot; mask++ {
			// 验证人数的有效性
			if ivCount[mask] > iv || evCount[mask] > ev {
				continue
			}
			// DFS, 计算下一行对应的得分
			cur = max(cur, dfs(row+1, mask, iv-ivCount[mask], ev-evCount[mask])+
				innerScore[mask]+interScore[preMask][mask],
			)
		}
		*p = cur
		return cur
	}

	return dfs(0, 0, introvertsCount, extrovertsCount)
}
