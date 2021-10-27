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
	var rl, rr = check([]byte(s))
	var left, right = rl, rr

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
