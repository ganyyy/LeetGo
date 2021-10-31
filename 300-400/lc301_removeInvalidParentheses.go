package main

func removeInvalidParentheses(s string) []string {
	// DFS + 记忆化查询?

	var tmp = make([]byte, 0, len(s))
	var ret []string
	var helper func(idx int)

	var check = func(s []byte) (int, int) {
		var rl, rr int
		for i := range s {
			switch s[i] {
			case '(':
				rl++
			case ')':
				if rl > 0 {
					rl--
				} else {
					rr++
				}
			}
		}
		return rl, rr
	}

	var isValid = func(a, b int) bool {
		return a == 0 && b == 0
	}

	// 计算不合法的左括号/右括号数量, 这个就是要删除的括号
	var left, right = check([]byte(s))

	var set = make(map[string]bool)

	// 枝减做的不够好啊
	helper = func(idx int) {
		if idx == len(s) {
			if left == 0 && right == 0 && isValid(check(tmp)) && !set[string(tmp)] {
				set[string(tmp)] = true
				ret = append(ret, string(tmp))
			}
			return
		}

		var i = idx
		if s[i] != '(' && s[i] != ')' {
			tmp = append(tmp, s[i])
			helper(i + 1)
			tmp = tmp[:len(tmp)-1]
		} else {
			if s[i] == '(' {
				if left > 0 {
					left--
					helper(i + 1)
					left++
				}
			} else if s[i] == ')' {
				if right > 0 {
					right--
					helper(i + 1)
					right++
				}
			}
			tmp = append(tmp, s[i])
			helper(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}

	helper(0)

	return ret
}

func isValid(str string) bool {
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func helper(ans *[]string, str string, start, lRemove, rRemove int) {
	if lRemove == 0 && rRemove == 0 {
		if isValid(str) {
			*ans = append(*ans, str)
		}
		return
	}

	for i := start; i < len(str); i++ {
		if i != start && str[i] == str[i-1] {
			continue
		}
		// 如果剩余的字符无法满足去掉的数量要求，直接返回
		if lRemove+rRemove > len(str)-i {
			return
		}
		// 尝试去掉一个左括号
		if lRemove > 0 && str[i] == '(' {
			helper(ans, str[:i]+str[i+1:], i, lRemove-1, rRemove)
		}
		// 尝试去掉一个右括号
		if rRemove > 0 && str[i] == ')' {
			helper(ans, str[:i]+str[i+1:], i, lRemove, rRemove-1)
		}
	}
}

func removeInvalidParenthesesFast(s string) (ans []string) {
	lRemove, rRemove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lRemove++
		} else if ch == ')' {
			if lRemove == 0 {
				rRemove++
			} else {
				lRemove--
			}
		}
	}

	helper(&ans, s, 0, lRemove, rRemove)
	return
}
