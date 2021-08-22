package main

import "math/bits"

func countArrangement(n int) (ans int) {
	vis := make([]bool, n+1)
	match := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j)
			}
		}
	}

	var backtrack func(int)
	backtrack = func(index int) {
		if index > n {
			ans++
			return
		}
		for _, x := range match[index] {
			if !vis[x] {
				vis[x] = true
				backtrack(index + 1)
				vis[x] = false
			}
		}
	}
	backtrack(1)
	return
}

func countArrangement2(n int) int {
	f := make([]uint16, 1<<n)
	f[0] = 1
	for mask := 1; mask < 1<<n; mask++ {
		// 获取mask对应的1的位数. 可以理解为是完美队列的长度
		num := bits.OnesCount(uint(mask))
		// 这里判断的是, 在第num位上放置i, 所能构成完美队列的可能
		for i := 0; i < n; i++ {
			// mask>>i&1 表示第i+1位已经被选取
			// 从mask已选取的数字中查询可以放置到num位上的组合
			if mask>>i&1 > 0 && (num%(i+1) == 0 || (i+1)%num == 0) {
				f[mask] += f[mask^1<<i]
			}
		}
	}
	return int(f[1<<n-1])
}
