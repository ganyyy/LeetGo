package main

func uniqueLetterString(s string) (ans int) {

	// sum: 到当前为止, ans中字符串调用countUniqueChars的返回结果
	// ABC -> 3
	// ABCA -> 2, 3-2*0+(-1) = 2
	sum, last := 0, [26][2]int{}

	// last[c][1]:上上次出现的位置
	// last[c][0]:上次出现的位置
	for i := range last {
		// 初始情况下, 全是-1
		last[i] = [2]int{-1, -1}
	}

	// 题解
	// https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/solution/by-endlesscheng-ko4z/
	// 简而言之, 就是计算

	for i, c := range s {
		c -= 'A'
		// (i-[0])-([0]-[1])
		// 好神奇啊

		sum += i - last[c][0]*2 + last[c][1]
		ans += sum
		last[c][1] = last[c][0]
		last[c][0] = i
	}
	return
}

func main() {
	uniqueLetterString("ABCA")
}
