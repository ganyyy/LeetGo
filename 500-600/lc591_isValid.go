package main

import (
	"strings"
	"unicode"
)

func isValid(code string) bool {
	var tags []string
	for code != "" {
		if code[0] != '<' {
			// 验证开头
			if len(tags) == 0 {
				return false
			}
			code = code[1:]
			continue
		}
		if len(code) == 1 {
			// < 开头, 后续要么是start_tag, 要么是end_tag
			return false
		}
		if code[1] == '/' {
			// 接尾标签 </
			j := strings.IndexByte(code, '>')
			if j == -1 {
				return false
			}
			// 对比end_tag和start_tag
			if len(tags) == 0 || tags[len(tags)-1] != code[2:j] {
				return false
			}
			tags = tags[:len(tags)-1]
			code = code[j+1:]
			// 不能有存在于标签外的数据
			if len(tags) == 0 && code != "" {
				return false
			}
		} else if code[1] == '!' {
			// 处理CDATA标签
			if len(tags) == 0 || len(code) < 9 || code[2:9] != "[CDATA[" {
				return false
			}
			// 直接查到末尾, 这段数据不需要检查
			j := strings.Index(code, "]]>")
			if j == -1 {
				return false
			}
			code = code[j+1:]
		} else {
			// 这是一个start_tag
			j := strings.IndexByte(code, '>')
			if j == -1 {
				return false
			}
			// tag合法化校验
			tagName := code[1:j]
			if tagName == "" || len(tagName) > 9 {
				return false
			}
			for _, ch := range tagName {
				if !unicode.IsUpper(ch) {
					return false
				}
			}
			tags = append(tags, tagName)
			code = code[j+1:]
		}
	}
	return len(tags) == 0
}
