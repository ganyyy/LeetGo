package main

func findRepeatedDnaSequencesBad(s string) []string {
	// 来个简单的吧..

	var cnt = make(map[string]int, len(s)-10)

	for i := 0; i <= len(s)-10; i++ {
		cnt[s[i:i+10]]++
	}

	var ret = make([]string, 0, len(cnt)>>1)
	for k, c := range cnt {
		if c > 1 {
			ret = append(ret, k)
		}
	}
	return ret
}

// 四个字母随机组合, 使用两位进行编码
var set = []int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

func findRepeatedDnaSequences(s string) (ans []string) {
	n := len(s)
	if n <= 10 {
		return
	}
	// 20位既可以表示一个10字符组成的字符串
	h := 0

	// 先统计头部的十个字符组成的数字
	for _, b := range s[:10] {
		h = h<<2 | set[b]
	}
	// 当成set使用, 本质上也是计数
	c := [1 << 20]int8{}
	c[h] = 1
	for i := 10; i < n; i++ {
		// 加一个新字符, 去掉旧的字符
		h = (h<<2 | set[s[i]]) % (1 << 20)
		if c[h] < 2 {
			if c[h] == 1 {
				ans = append(ans, s[i-9:i+1])
			}
			c[h]++
		}
	}
	return
}
