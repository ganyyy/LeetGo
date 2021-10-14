package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	pre := countAndSay(n - 1)

	head := pre[0]
	count := 1
	res := make([]byte, 0)
	for i := 1; i < len(pre); i++ {
		if pre[i] != head {
			// X个X
			res = append(res, byte(count)+'0', head)
			head = pre[i]
			count = 1
		} else {
			count++
		}
	}
	res = append(res, byte(count)+'0', head)
	return string(res)
}

func countAndSay2(n int) string {
	if n == 1 {
		return "1"
	}
	var pre = countAndSay2(n - 1)
	var cur = pre[0]
	var cnt int
	var sb strings.Builder
	for i := range pre {
		if cur == pre[i] {
			cnt++
		} else {
			sb.WriteString(strconv.Itoa(cnt) + pre[i-1:i])
			cur = pre[i]
			cnt = 1
		}
	}
	if cnt != 0 {
		sb.WriteString(strconv.Itoa(cnt) + string(cur))
	}
	return sb.String()
}

func main() {
	fmt.Println(countAndSay(5))
}
