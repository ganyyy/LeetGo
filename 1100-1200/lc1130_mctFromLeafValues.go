package main

func mctFromLeafValues(arr []int) int {
	// stk: 单调递减栈
	var stk []int
	res := 0
	for _, x := range arr {
		// 合并的要点: 大的放到最后合并, 这样才能保证合并过程中间节点的和最小
		/*
		   [5,4,6] -> [5,4] + [6] = [20(5),6] = [30], 20+30 = 50 √
		   [5,4,6] -> [5] + [4,6] = [5,24(6)] = [30], 24+30 = 54 ×

		   [6,4,5] -> [6] + [5,4] = [6,20(5)] = [30], 20+30 = 50 √
		   [6,4,5] -> [6,5] + [4] = [5,24(6)] = [30], 24+30 = 54 ×
		*/
		for len(stk) > 0 && stk[len(stk)-1] <= x {
			if len(stk) == 1 || stk[len(stk)-2] > x {
				// 只剩一个元素, 或者 x < 栈顶前边的元素
				res += stk[len(stk)-1] * x
			} else {
				// 超过一个元素, 并且 x > 栈顶前边的元素
				res += stk[len(stk)-2] * stk[len(stk)-1]
			}
			// 栈顶元素出栈
			stk = stk[:len(stk)-1]
		}
		// 当前元素入栈
		stk = append(stk, x)
	}
	for len(stk) > 1 {
		// 如果一路递减, 或者还存在未消耗完成的元素, 挨个计算
		res += stk[len(stk)-2] * stk[len(stk)-1]
		stk = stk[:len(stk)-1]
	}
	return res
}
