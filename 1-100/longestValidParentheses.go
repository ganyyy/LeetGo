package main

import "fmt"

const (
	LEFT  = '('
	RIGHT = ')'
)

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	stack := make([]int, 0, len(s))
	match := make([]int, 0, len(s))
	var res int
	// 当作栈来用
	for _, v := range s {
		if v == LEFT {
			stack = append(stack, len(match))
			match = append(match, 0)
		} else {
			sl := len(stack)
			if sl != 0 {
				// 对应标志位置1, 表示已经匹配上了
				match[stack[sl-1]] = 1
				// 出栈
				stack = stack[:sl-1]
			} else {
				// 不匹配就增加一个非法标志位
				match = append(match, 0)
			}
		}
	}
	var max int
	for i := 0; i < len(match); i++ {
		if match[i] == 1 {
			res++
		} else {
			if max < res {
				max = res
			}
			res = 0
		}
	}
	if res < max {
		res = max
	}
	return res * 2
}

func longestValidParentheses2(s string) int {
	if len(s) <= 1 {
		return 0
	}
	// dp[i]表示的是  s[i] == ')' 并且 最远的有效匹配 括号长度
	dp := make([]int, len(s))
	var res int
	for i := 1; i < len(s); i++ {
		// 进行匹配
		if s[i] == RIGHT {
			// 前边匹配
			pre := i - dp[i-1] - 1
			// 针对 ()
			if pre >= 0 && s[pre] == LEFT {
				// 先加上自己的这两个
				// 比如 ")(())" 的 dp是[0,0,0,2,4]
				val := dp[i-1] + 2
				// 针对 ()(), 前边可能有, 也可能没有
				if pre > 0 {
					// ")()()"的dp是 [0,0,2,0,4]
					val += dp[pre-1]
				}
				// 更新最大值
				if val > res {
					res = val
				}
				dp[i] = val
			}
		}
	}
	return res
}

func main() {
	fmt.Println(longestValidParentheses2("()(()"))
}
