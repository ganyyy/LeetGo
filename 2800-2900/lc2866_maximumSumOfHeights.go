package main

func maximumSumOfHeights2(maxHeights []int) int64 {
	length := len(maxHeights)
	// prefix[i]表示以i为结尾的最大值, [0,i]
	// suffix[i]表示以i为开头的最大值, [i, n-1]
	prefix, suffix := make([]int, length), make([]int, length)
	// 递增栈
	var stack []int

	// 不管是左还是右, 核心都是构建一个非递减栈
	// 在构建栈的过程中, 如果遇到了比栈顶小的元素, 那么就可以计算一次结果了
	// 相当于这两个位置之前都填充当前的高度(当没有比当前高度小的时候, 就相当于从0开始)
	for i := 0; i < length; i++ {
		height := maxHeights[i]
		for len(stack) > 0 && height < maxHeights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			prefix[i] = (i + 1) * height
		} else {
			// [0,1,2], 5 -> (3,4,5)
			last := stack[len(stack)-1]
			prefix[i] = prefix[last] + (i-last)*height
		}
		stack = append(stack, i)
	}
	stack = stack[:0]
	var ret int
	for i := length - 1; i >= 0; i-- {
		height := maxHeights[i]
		for len(stack) > 0 && height < maxHeights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			suffix[i] = (length - i) * height
		} else {
			// [7,6,5], 2 -> (4,3,2)
			last := stack[len(stack)-1]
			suffix[i] = suffix[last] + (last-i)*height
		}
		stack = append(stack, i)
		ret = max(ret, suffix[i]+prefix[i]-height)
	}
	return int64(ret)
}
