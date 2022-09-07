package main

import "strings"

func reorderSpaces(text string) string {
	var words []string
	var cnt int
	var l, r int
	for l, r = 0, 0; r < len(text); r++ {
		if text[r] == ' ' {
			cnt++
			if l != r {
				words = append(words, text[l:r])
			}
			l = r + 1
		}
	}

	if l != r {
		words = append(words, text[l:])
	}

	rep := func(c int) string { return strings.Repeat(" ", c) }

	if len(words) == 0 || cnt == 0 {
		return ""
	}
	if len(words) == 1 {
		return words[0] + rep(cnt)
	}

	block := cnt / (len(words) - 1)

	// fmt.Println(words, cnt, block)

	var sb strings.Builder
	var cur int
	for i := 0; i < len(words)-1; i++ {
		sb.WriteString(words[i])
		sb.WriteString(rep(block))
		cur += block
	}
	sb.WriteString(words[len(words)-1])
	if cur != cnt {
		sb.WriteString(rep(cnt - cur))
	}

	return sb.String()
}
