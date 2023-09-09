package main

import "fmt"

func strStr(haystack, needle string) int {
	// Mark: kmp算法
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		// next 数组的意义: next[i] 表示 next[0:next[i]] == next[i-next[i]+1: i+1]
		// 所以一旦指定位置上的字符不匹配, 可以回退到next[i]上继续尝试进行匹配
		// 这也是一种自匹配.
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	fmt.Println(pi)
	for i, j := 0, 0; i < n; i++ {
		// 这里就很明显了, 如果某一个位置上不匹配了, 此时就需要进行回退
		// 假设现在 j == 10, next[j] == 4,
		// 说此时needle[:4] == needle[7:11]. 但是很明显, 因为 needle[10] != haystack[i], 所以此时需要进行回退
		// 此时 haystack[i-3:i] == needle[7:10]. 很明显, 可以回退的最远位置就是 next[9],
		// 即从开头开始, 满足 needle[0:next[9]] == needle[9-next[9]+1 : 10],
		// 迭代查询, 直到找到了可以回退的位置或者 从头开始 确定一个最终的 j
		// 下一步就是对比haystack[i] 和 needle[j]是否相等
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func main() {
	strStr("12345", "aabcaab")
}
