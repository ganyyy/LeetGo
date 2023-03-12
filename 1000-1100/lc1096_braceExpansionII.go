package main

import (
	"sort"
	"strings"
)

func braceExpansionII(expression string) []string {
	var queue []string
	queue = append(queue, expression)

	var set = make(map[string]struct{})
	var sb strings.Builder

	for len(queue) != 0 {
		exp := queue[0]
		queue = queue[1:]

		if strings.Index(exp, "{") == -1 {
			set[exp] = struct{}{}
			continue
		}

		// 计算第一对括号
		var start, end int
		// {a,{a,b,c},c} -> {a,b,c}
		for i, c := range exp {
			if c == '}' {
				end = i
				break
			}
			// 起点需要不停的更新, 这样才可以拆出最内部的括号
			if c == '{' {
				start = i
			}
		}

		before := exp[:start]
		after := exp[end+1:]

		// 这一对括号中的所有字母 {a,b,c} -> [a, b, c], 展开并添加到队列中
		for _, sub := range strings.Split(exp[start+1:end], ",") {
			sb.Reset()
			sb.WriteString(before)
			sb.WriteString(sub)
			sb.WriteString(after)
			queue = append(queue, sb.String())
		}
	}

	var ret = make([]string, 0, len(set))
	for sub := range set {
		ret = append(ret, sub)
	}
	sort.Strings(ret)
	return ret
}
