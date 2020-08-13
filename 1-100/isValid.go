package main

import "fmt"

func isValid(s string) bool {
	var mapChar = map[uint8]uint8{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	var stack = make([]uint8, 0)

	for _, _v := range s {
		if len(stack) == 0 || mapChar[stack[len(stack)-1]] != uint8(_v) {
			// 空栈 或者栈顶元素和当前符号不匹配
			stack = append(stack, uint8(_v))
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 怎么感觉还退步了?
func isValid2(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		t := s[i]
		if t == '(' || t == '{' || t == '[' {
			stack = append(stack, t)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch t {
			case ')':
				if top != '(' {
					return false
				}
			case ']':
				if top != '[' {
					return false
				}
			case '}':
				if top != '{' {
					return false
				}
			default:
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("([)]"))
}
