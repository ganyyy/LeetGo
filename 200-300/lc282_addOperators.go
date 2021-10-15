package main

func addOperators(num string, target int) (ans []string) {
	n := len(num)
	var backtrack func(expr []byte, i, res, mul int)
	backtrack = func(expr []byte, i, res, mul int) {
		if i == n {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}
		signIndex := len(expr)
		if i > 0 {
			expr = append(expr, 0) // 占位，下面填充符号
		}
		// 枚举截取的数字长度（取多少位），注意数字可以是单个 0 但不能有前导零
		for j, val := i, 0; j < n && (j == i || num[i] != '0'); j++ {
			val = val*10 + int(num[j]-'0')
			expr = append(expr, num[j])
			if i == 0 { // 表达式开头不能添加符号
				backtrack(expr, j+1, val, val)
			} else { // 枚举符号
				expr[signIndex] = '+'
				backtrack(expr, j+1, res+val, val)
				expr[signIndex] = '-'
				backtrack(expr, j+1, res-val, -val)
				// 乘法作为一个整体来看, 所以需要先把旧的mul减去, 在加上重新计算的整体数值
				expr[signIndex] = '*'
				backtrack(expr, j+1, res-mul+mul*val, mul*val)
			}
		}
	}
	backtrack(make([]byte, 0, n*2-1), 0, 0, 0)
	return
}
