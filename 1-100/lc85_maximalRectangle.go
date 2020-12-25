package main

const (
	Empty byte = '0'
	Fill  byte = '1'
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxArea(heights []int) int {
	stack := make([]int, 0, len(heights)+1)
	stack = append(stack, -1)
	var mArea int
	// 和84题一样的套路
	for i := 0; i < len(heights); i++ {
		for t := len(stack) - 1; t > 0 && heights[stack[t]] >= heights[i]; t-- {
			// 找最大值,
			mArea = max(mArea, heights[stack[t]]*(i-stack[t-1]-1))
			// 最后一个出栈
			stack = stack[:t]
		}
		stack = append(stack, i)
	}
	// 看看栈是否为空
	for i, h := len(stack)-1, len(heights); i > 0; i-- {
		mArea = max(mArea, heights[stack[i]]*(h-stack[i-1]-1))
	}
	return mArea
}

func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	mArea := 0
	// 标记每一行的高度
	dp := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == Empty {
				dp[j] = 0
			} else {
				dp[j] += 1
			}
		}
		// 将其每一行转换为一个柱状图进行求解, 同样的, 下一行会继承上一行的长度,
		// 如果一直为'1', dp[j]就会一直累加, 如果出现了一个0, 就会清零从新计算
		mArea = max(mArea, getMaxArea(dp))
	}
	return mArea
}

func maximalRectangle2(matrix [][]byte) int {
	// 对每一行按照上一题的做法进行处理
	if len(matrix) == 0 {
		return 0
	}
	var res int
	var tmp = make([]int, len(matrix[0]))
	for _, row := range matrix {
		for i, b := range row {
			if b == '1' {
				tmp[i]++
			} else {
				tmp[i] = 0
			}
		}
		res = max(res, largestRectangleArea23(tmp))
	}

	return res
}

func largestRectangleArea23(heights []int) int {
	// 单调递增栈处理
	var stack = make([]int, 1, len(heights)+1)
	stack[0] = -1

	var res int
	for i, v := range heights {
		for t := len(stack) - 1; t != 0 && heights[stack[t]] >= v; t-- {
			// 长度就是栈中的最后一个数值
			// 宽度就是 i 距离 栈顶前一个元素的距离
			res = max(res, heights[stack[t]]*(i-stack[t-1]-1))
			stack = stack[:t]
		}
		stack = append(stack, i)
	}

	// 计算剩余的递增栈
	for t, h := len(stack)-1, len(heights); t != 0; t-- {
		res = max(res, heights[stack[t]]*(h-stack[t-1]-1))
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
