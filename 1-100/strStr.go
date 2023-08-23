//go:noinline

package main

import (
	"fmt"
)

// dp[j][char] = next
// j 表示是当前状态
// char 要匹配的字符
// next 表示在j状态下遇到char字符的下一个状态
var dm []map[byte]int
var pat string

func makeDP() {
	lp := len(pat)
	dm = make([]map[byte]int, lp)
	m := make(map[byte]struct{})
	x := 0 // 影子状态
	dm[0] = make(map[byte]int)
	p0 := pat[0]
	dm[0][pat[0]] = 1
	m[p0] = struct{}{}
	for i := 1; i < lp; i++ {
		dm[i] = make(map[byte]int)
		c := pat[i]
		// 正确的字符状态+1
		dm[i][c] = i + 1
		for k := range m {
			// 遍历之前已经存在的字符,
			// 依次根据判断当前位置遇到之前存在的字符时
			// 应该回退到那个位置。
			// 这里的x可以理解为 从 0 到 x 是 当前位置的最长
			// 重复子串
			// 如 abab
			// 对应的重复子串是 a, ab, aba, abcd
			// 所以在满足a匹配成功后, 如果下一个字符是 b, 状态+1, 同时 x 变成 bm[0][a] = 1
			if k != c {
				dm[i][k] = dm[x][k]
			}
		}
		_, ok := m[c]
		if !ok {
			m[c] = struct{}{}
		}
		// 更新一下x
		x = dm[x][c]
		fmt.Println(x)
	}
}

func search_(txt string) int {
	lp, lt := len(pat), len(txt)
	var j int // 初始状态
	for i := 0; i < lt; i++ {
		j = dm[j][txt[i]]
		fmt.Println(j, txt[i])
		if j == lp {
			// 匹配成功, 返回起始的索引
			return i - lp + 1
		}
	}
	// 没有匹配成功的数据, 返回-1
	return -1
}

func main() {
	pat = "bcc"
	makeDP()
	fmt.Println(dm)

	fmt.Println(search_("aaaababbbbb"))
}
