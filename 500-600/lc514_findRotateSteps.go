//go:build ignore

package main

import (
	"fmt"
	"math"
)

// 失败了, 无法处理 相同字符连接字串的问题
func findRotateStepsFail(ring string, key string) int {
	// 不好控制 左右的前进步数
	// 可以考虑搭建一个全映射, 最大内存消耗 1W?
	// 需要按照顺序来, 这样就好处理了

	// 一共26个字符, 最多 26*26 个空间搭建映射

	// m[i][j] 表示 从 i->j所需要的最短路径
	// 1. 每一个位置向两边发散, 统计路途中所遇到的最短的路径
	// 从起点开始, 然后依次旋转?

	var m [26][26]int

	// 单数情况下, 左右路途一致
	// 双数情况下, 选取任意一边多走一格(是否可行?)
	var ln = len(ring)
	// 单次最多走多远

	var left = (ln - 1) >> 1
	var right = (ln - 1) - left

	for c := 0; c < ln; c++ {
		var cur = int(ring[c] - 'a')
		// fmt.Printf("cur:%c\t", cur+'a')
		for i := 1; i <= left; i++ {
			var pos = (i + c + ln) % ln
			// 自己到自己永远是0
			if ring[pos] == ring[c] {
				continue
			}
			// fmt.Printf("%c\t", ring[pos])
			m[cur][ring[pos]-'a'] = min(m[cur][ring[pos]-'a'], i)
		}
		for i := 0; i < right; i++ {
			var pos = (ln - 1 - i + c) % ln
			// 自己到自己永远是0
			if ring[pos] == ring[c] {
				continue
			}
			// fmt.Printf("%c\t", ring[pos])
			m[cur][ring[pos]-'a'] = min(m[cur][ring[pos]-'a'], i+1)
		}
		// fmt.Printf("\t\t")
		// for j, jj := range m[cur] {
		//	if jj == 0 {
		//		continue
		//	}
		//	fmt.Printf("%c-%d \t", j+'a', jj)
		// }
		//
		// fmt.Println()

	}

	// for i, v := range m {
	//	fmt.Printf("cur:%c \t", i+'a')
	//	for j, c := range v {
	//		if c == 0 {
	//			continue
	//		}
	//		fmt.Printf("%c-%d \t", j+'a', c)
	//	}
	//	fmt.Printf("\n")
	// }
	var res int
	var cur = ring[0]
	for _, c := range key {
		res += m[cur-'a'][c-'a'] + 1
		cur = byte(c)
	}

	return res
}

func findRotateSteps(ring, key string) int {
	var inf = math.MaxInt32
	var n, m = len(ring), len(key)

	// ring 每一个字符出现的位置的集合
	var pos [26][]int
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}

	// dp[i][j] 表示 从 拼到第i个key的字符, ring的第j个字符与12:00方向对其的最少步数

	// 初始化每一个位置. 后期考虑压缩成一维数组
	var dp = make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	// 遍历 首个字符 的所有位置
	for _, c := range pos[key[0]-'a'] {
		// 找出ring中key[0]字符 相对于 ring[0](初始12:00的字符) 最短距离
		dp[0][c] = min(c, n-c) + 1
	}

	for i := 1; i < m; i++ {
		// 查找key中每一个字符, 找出相对于 ring 12:00 方向最近操作步骤
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
			}
		}
	}

	return minSlice(dp[m-1]...)
}

func minSlice(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	var m = a[0]
	for _, v := range a[1:] {
		if m > v {
			m = v
		}
	}
	return m
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(findRotateSteps("godding",
		"godding"))
}
