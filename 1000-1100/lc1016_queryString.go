//go:build ignore

package main

import (
	"math/bits"
	"strconv"
	"strings"
)

func queryString(s string, n int) bool {
	if n == 1 {
		return strings.Contains(s, "1")
	}

	// 使用区间 [4, 7]进行距离, 长度均为3, 如果有一个字符串可以同时满足这个区间内的4个数字,
	// 考虑到可以存在重叠的部分, 那么其长度最少为
	// 3 + (7-4+1-1) = 6. (7-4+1 是区间内的数字个数, -1是去掉了重叠的部分)
	// 1011100, 实际貌似得是这个才能覆盖整个区间(?)

	m := len(s)                // s的长度等同于二进制位数
	k := bits.Len(uint(n)) - 1 // 获取到最高位的1所处的二进制的位置
	// 此时可以按照 2^k 将 n 分为三部分
	// 1. [2^k, n], 这是因为n可能不是2的整数次幂, 并且其长度都是 k+1.
	//    所以需要满足 m >= (k+1) + (n-2^k + 1 - 1) = n-2^k+k-1
	// 2. [2^(k-1), 2^k - 1], 这段区间内的数字的 长度都为 k
	//    所以需要满足 m >= k + (2^(k-1)-1) = 2^(k-1) + k - 1
	// 3. [0, 2^(k-1)-1], 这一部分可以通过第二部分不停的右移获取到,
	//    如果第二部分是存在的, 那么这一部分也必然是满足的

	// 首先判断以下是否满足最低的长度需求
	if m < max(n-1<<k+k+1, 1<<(k-1)+k-1) {
		return false
	}

	// 对于长为 k 的在 [lower, upper] 内的二进制数，判断这些数 s 是否都有
	check := func(k, lower, upper int) bool {
		if lower > upper {
			return true
		}
		seen := map[int]struct{}{}
		mask := 1<<(k-1) - 1
		// 当成滑动窗口来看, 窗口的长度为k, 然后依次将其对应的二进制数字放入到seen中
		v, _ := strconv.ParseUint(s[:k-1], 2, 64)
		x := int(v)
		for _, c := range s[k-1:] {
			// &mask 可以去掉最高比特位，从而实现滑窗的「出」
			// <<1 | int(c-'0') 即为滑窗的「入」
			x = x&mask<<1 | int(c-'0')
			// 保证上下界的合法性
			if lower <= x && x <= upper {
				seen[x] = struct{}{}
			}
		}
		return len(seen) == upper-lower+1
	}

	// 分别见检查 1, 2两部分是否合法
	return check(k, n/2+1, 1<<k-1) && check(k+1, 1<<k, n)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
