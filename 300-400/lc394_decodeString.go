package main

import (
	"fmt"
	"strings"
)

func decodeString(s string) string {
	// 两个栈, 一个数字栈, 一个字符栈
	sStack := make([]string, 0)
	nStack := make([]int, 0)

	// "3[a]2[bc]", 返回 "aaabcbc".
	var cur strings.Builder
	var num int
	for _, v := range s {
		if v >= '0' && v <= '9' {
			num = num*10 + int(v) - '0'
			continue
		}
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			cur.WriteByte(byte(v))
			continue
		}
		if v == '[' {
			nStack = append(nStack, num)
			num = 0
			sStack = append(sStack, cur.String())
			cur.Reset()
			continue
		}
		if v == ']' {
			// 取出要重复多少次
			n := nStack[len(nStack)-1]
			nStack = nStack[:len(nStack)-1]
			// 取当前要重复的字符
			t := cur.String()
			cur.Reset()
			// 取栈顶的字符串, 先写入, 再出栈
			cur.WriteString(sStack[len(sStack)-1])
			sStack = sStack[:len(sStack)-1]
			// 写入指定次数的 当前值
			for i := 0; i < n; i++ {
				cur.WriteString(t)
			}
		}
	}
	return cur.String()
}

func main() {
	/*
		s = "3[a]2[bc]", 返回 "aaabcbc".
		s = "3[a2[c]]", 返回 "accaccacc".
		s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
	*/
	testCase := []string{
		//"3[a]2[bc]",
		//"3[a2[c]]",
		//"2[abc]3[cd]ef",
		//"cd2[a3[b]]ef",
		"100[leetcode]",
	}

	for _, v := range testCase {
		fmt.Println(decodeString(v))
	}
}
