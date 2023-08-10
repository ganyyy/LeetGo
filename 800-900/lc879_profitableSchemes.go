package main

// func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
//     // // 背包大小为 n
//     // // 物体体积为 group[i]
//     // // 物体价值为 profit[i]
//     // // 最小实现的价值为 minProfit
//     // // 最终输出的结果为?

//     // var mod int = 10E9 + 7
//     // var ret int

//     // var g, p int

//     // // 搞个bitset ?
//     // // 最大值100, 两个int64就够了

//     // // dp保存的是到当前位置满足大于最小profit的值?
//     // var dp = make([]int, n+1)
//     // for i := range group {
//     //     g, p = group[i], profit[i]

//     //     for i := n; i >= p; i-- {
//     //         dp[i] = max(dp[i], dp[i-g]+p)
//     //         if dp[i] >= minProfit {
//     //             fmt.Println(dp)
//     //             // 假设有一种方法, 可以获取到当前位置是第几种实现的方式的话..?
//     //         }
//     //     }
//     //     // 判断一下当前的利润和是否超过了minProfit
//     // }
//     // return ret%mod

//     const MOD int = 1000000007

//     // 选取group子集G, minProfit子集P, 使其满足
//     // sum(G) <= n; sum(P) >= minProfit
//     // 一共多少种G的选法?

//     // 最大利润
//     var sum int
//     for _, v := range profit {
//         sum += v
//     }

//     var pre = make([][]int, n+1)
//     for i := range pre {
//         pre[i] = make([]int, sum+1)
//     }
//     // 起始条件: 利润0, 一个工作都不做

//     for i := 0; i <= n; i++ {
//         pre[i][0] = 1
//     }

//     // 学会复用内存
//     // dp[i][j][k] 表示 前i种工作, 总人数为j时, 使得利润为k的时候的计划的数量
//     // 如果当前组合可以做, 那么就是 dp[i][j][k] += dp[i-1][j-group[i-1]][k-profit[i-1]]
//     var cur = make([][]int, len(pre))
//     for i := range pre {
//         cur[i] = make([]int, len(pre[i]))
//     }

//     for i := 1; i <= len(group); i++ {
//         cur[0][0] = 1   // 最小利润为0时, 只有一种计划: 什么都不做

//         for j := 1; j <= n; j++ {
//             for k := 0; k <= sum; k++ {
//                 cur[j][k] = pre[j][k]
//                 // 如果人数够, 并且...?
//                 if j-group[i-1] >= 0 && k-profit[i-1] >= 0 {
//                     cur[j][k] = (cur[j][k]+pre[j-group[i-1]][k-profit[i-1]]%MOD)%MOD
//                 }
//             }
//         }

//         // 结束后需要清理一下, 后续继续使用
//         pre, cur = cur, pre
//         for i := range cur {
//             var t = cur[i]
//             for j := range t {
//                 t[j] = 0
//             }
//         }
//     }

//     var ret int

//     for i := minProfit; i <= sum; i++ {
//         ret = (ret+pre[n][i]%MOD)%MOD
//     }

//     return ret
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	g := len(group)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, minProfit+1)
	}
	for i := 0; i < n+1; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < g; i++ {
		// 人数, 利润
		// 整体解题思路和上边一样, 通过逆序来降低维度(只和前态有关, 所以可以压缩)
		manual := group[i]
		money := profit[i]
		for j := n; j >= manual; j-- {
			for k := minProfit; k >= 0; k-- {
				// k-money < 0 说明选择当前任务时, 所获取的利润一定大于minProfit
				dp[j][k] += dp[j-manual][max(0, k-money)]
				dp[j][k] %= 1e9 + 7
			}
		}
	}
	return dp[n][minProfit]
}
