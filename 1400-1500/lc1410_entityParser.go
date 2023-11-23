package main

import "strings"

func entityParser(text string) string {
	var sb strings.Builder
next:
	for len(text) != 0 {
		// 1. &&&&xxx;;;;;;
		// 需要定位最后一个&和第一个;
		// 2. &x&xxx;
		// 如果不符合标准, 但是接下来还有一个&, 需要二次迭代
		and := strings.Index(text, "&")
		if and == -1 {
			and = len(text)
		} else {
			for ; and+1 < len(text) && text[and+1] == '&'; and++ {
			}
		}
		sb.WriteString(text[:and])
		text = text[and:]
		if len(text) == 0 {
			break
		}
		var split = 1
		for ; split < len(text); split++ {
			if text[split] == ';' {
				break
			}
			if text[split] == '&' {
				sb.WriteString(text[:split])
				text = text[split:]
				continue next
			}
		}
		if split == len(text) {
			sb.WriteString(text)
			break
		}
		var s string
		switch text[:split] {
		case "&quot":
			s = "\""
		case "&apos":
			s = "'"
		case "&amp":
			s = "&"
		case "&gt":
			s = ">"
		case "&lt":
			s = "<"
		case "&frasl":
			s = "/"
		default:
			s = text[:split+1]
		}
		sb.WriteString(s)
		text = text[split+1:]
	}
	return sb.String()
}
