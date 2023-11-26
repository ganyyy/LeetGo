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

func uniqueLetterString2(s string) int {
	type ByteCount struct {
		First, Second int
	}
	var count [26]ByteCount

	for i := range count {
		count[i] = ByteCount{First: -1, Second: -1}
	}

	// xxAxxxxAxxxxA
	// [0] = 2, [1] = 7, i = 12
	// [8, 12]中, ________xxxxA 中, 每个子串的不重复字符 +1(首次出现A)
	// [3, 12]中, ___xxxxA____A 中, 每个子串的不重复字符 -1(重复出现一次A)
	// [0, 12]中, xxA____A____A 中, 每个字串的不重复字符不变(已经出现2次A)
	// 那么从[0, 11] -> [0, 12]时, 总共增加了 12 - 7 = 5次, 减少了 7 - 2次
	// sum += (i-last0) - (last0-last1) = i-2*last0+last1
	var sum int
	var ret int
	for i, c := range s {
		bs := &count[c-'A']
		// (i-bs.First) - (bs.First-bs.Second)
		sum += i - 2*bs.First + bs.Second
		ret += sum
		bs.First, bs.Second = i, bs.First
	}
	return ret
}
