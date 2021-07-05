package main

import (
	"sort"
	"strconv"
	"strings"
)

func countOfAtoms(formula string) string {
	// 应该用到栈了吧

	var stack = []map[string]int{{}}
	var ln = len(formula)

	var getNum = func(i int) (next int, num int) {
		var k = i + 1
		for k < ln && formula[k] >= '0' && formula[k] <= '9' {
			k++
		}
		if k != i+1 {
			num, _ = strconv.Atoi(formula[i+1 : k])
		} else {
			num = 1
		}
		next = k - 1
		return
	}
	var num int
	for i := 0; i < ln; i++ {
		// 提取第一个数字
		var c = formula[i]
		switch {
		case c == '(':
			// 入栈一个新的计数
			stack = append(stack, map[string]int{})
		case c == ')':
			i, num = getNum(i)
			var top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var pre = stack[len(stack)-1]
			for i2, v := range top {
				pre[i2] += v * num
			}
		default:
			// 读取原子和其对应的个数
			var top = stack[len(stack)-1]
			var k = i + 1
			for k < ln && formula[k] <= 'z' && formula[k] >= 'a' {
				k++
			}
			var old = i
			i, num = getNum(k - 1)
			top[formula[old:k]] += num
		}
	}

	var m = stack[0]
	var tmp = make([]string, 0, len(m))
	for k := range m {
		tmp = append(tmp, k)
	}
	sort.Strings(tmp)
	var sb strings.Builder
	for _, s := range tmp {
		var v = m[s]
		if v == 1 {
			sb.WriteString(s)
		} else {
			sb.WriteString(s + strconv.Itoa(v))
		}
	}
	return sb.String()
}

func main() {
	println(countOfAtoms("H2O"))
}
