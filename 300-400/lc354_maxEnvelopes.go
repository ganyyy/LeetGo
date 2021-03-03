package main

import (
	"fmt"
	"sort"
)

// func maxEnvelopes1(envelopes [][]int) int {
//     var n = len(envelopes)
//     if n == 0 {
//         return 0
//     }

//     sort.Slice(envelopes, func(i, j int) bool {
//         if envelopes[i][0] != envelopes[j][0] {
//             return envelopes[i][0] <= envelopes[j][0]
//         }
//         return envelopes[i][1] <= envelopes[j][1]
//     })

//     // GBN?

//     // 好家伙, 来DP了!

//     var m = 1
//     var dp = make([]int, n)
//     for i := 0; i < n; i++ {
//         dp[i] = 1
//     }
//     for i := 1; i < n; i++ {
//         var ei = envelopes[i]
//         for j := 0; j < i; j++ {
//             var ej = envelopes[j]
//             if ei[0] > ej[0] && ei[1] > ej[1]  {
//                 dp[i] = max(dp[i], dp[j]+1)
//             }
//         }
//         m = max(dp[i], m)
//     }

//     return m
// }

func maxEnvelopes(envelopes [][]int) int {

	// 宽相等的情况下, 按照高的降序进行排序
	// 让宽维持升序, 高维持降序是为了让
	// 同宽情况下, 不同高之间无法构成升序序列
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	var f []int
	for _, e := range envelopes {
		// 可以理解为, 此时的宽已经是升序的了. 只需要从高里查找最大的升序队列即可
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

func main() {
	fmt.Println(maxEnvelopes([][]int{{1, 5}, {1, 4}, {2, 6}, {2, 7}, {3, 8}}))
}
