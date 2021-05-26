package main

import "fmt"

func reverseParentheses(s string) string {
	var stack1 [][]byte
	var bs = []byte(s)

	stack1 = append(stack1, bs[:0])
	for i := 0; i < len(bs); i++ {
		switch bs[i] {
		case '(':
			stack1 = append(stack1, bs[i+1:i+1])
		case ')':
			var top = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			reverse(top)
			stack1[len(stack1)-1] = append(stack1[len(stack1)-1], top...)
		default:
			stack1[len(stack1)-1] = append(stack1[len(stack1)-1], bs[i])
		}
	}

	return string(stack1[0])
}

func reverseParenthesesFail(s string) string {
	var bs = []byte(s)
	var stack [][]byte
	var depth, pre int
	for i := range bs {
		switch bs[i] {
		case ')':
			if len(stack) != depth+1 {
				reverse(bs[pre:i])
				stack[len(stack)-1] = append(stack[len(stack)-1], bs[pre:i]...)
			} else {
				var top = stack[len(stack)-1]
				top = append(top, bs[pre:i]...)
				stack = stack[:len(stack)-1]
				reverse(top)
				stack[len(stack)-1] = append(stack[len(stack)-1], top...)
			}

			depth--
			pre = i + 1
		case '(':
			if depth != len(stack)+1 {
				stack = append(stack, bs[pre:i])
			} else {
				stack[len(stack)-1] = append(stack[len(stack)-1], bs[pre:i]...)
			}
			depth++
			pre = i + 1
		}
	}
	if len(stack) == 0 {
		stack = append(stack, []byte(nil))
	}
	stack[0] = append(stack[0], bs[pre:]...)
	return string(stack[0])
}

func reverse(bs []byte) {
	fmt.Println("before: ", string(bs))
	for l, r := 0, len(bs)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	fmt.Println("after: ", string(bs))
}

func main() {
	println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
}
