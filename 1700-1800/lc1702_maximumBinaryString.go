package main

import "strings"

func maximumBinaryString1(binary string) string {
	n := len(binary)
	s := []byte(binary)
	j := 0
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			// 往后找到首个0, 然后移到[i+1]上, 在将00替换成10
			for j <= i || (j < n && s[j] == '1') {
				j++
			}
			if j < n {
				// 111110 -> 011111
				s[j] = '1'
				// 00 -> 10
				s[i] = '1'
				s[i+1] = '0'
			}
		}
	}
	return string(s)
}

func maximumBinaryString(binary string) string {
	n := len(binary)
	i := strings.Index(binary, "0")
	if i < 0 {
		// 没有0
		return binary
	}
	// 经过整体替换之后, 至多保留1个0
	zeros := strings.Count(binary, "0")
	// 这个道理是啥呢?
	s := []byte(strings.Repeat("1", n))
	// 尽可能的把后边的0移到前边, 并转换成1
	// 那么可以转换的0的个数就是总体0的个数-1
	// 所处的位置也是从首个0出现的位置开始+总数-1位置结束
	s[i+zeros-1] = '0'
	return string(s)
}
