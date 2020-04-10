package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	t := strings.Split(s, " ")

	var b strings.Builder

	for i := len(t) - 1; i >= 0; i-- {
		if t[i] != "" {
			b.WriteByte(' ')
			b.WriteString(t[i])
		}
	}

	res := b.String()
	if len(res) > 1 {
		return res[1:]
	} else {
		return ""
	}
}

func reverseWords2(s string) string {
	// 三步走
	if len(s) == 0 {
		return ""
	}
	bs := []byte(s)
	// 1. 整体反转
	reverse(bs, 0, len(bs)-1)
	// 2. 反转单词
	reverseWord(bs)
	// 3. 去掉多余空格
	return string(bs[:cancelSpace(bs)])
}

func reverse(s []byte, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseWord(s []byte) {
	var i, j int

	for j < len(s) {
		for i < len(s) && s[i] == ' ' {
			i++
		}
		j = i
		for j < len(s) && s[j] != ' ' {
			j++
		}
		reverse(s, i, j-1)
		i = j
	}
}

func cancelSpace(s []byte) int {
	var i, j int
	for j < len(s) {
		if s[j] != ' ' {
			s[i] = s[j]
			i++
		} else {
			if i > 0 && s[i-1] != ' ' {
				s[i] = ' '
				i++
			}
		}
		j++
	}
	if i > 0 && s[i-1] == ' ' {
		return i - 1
	} else {
		return i
	}
}

func main() {
	fmt.Println(reverseWords2("a good   example"))
}
