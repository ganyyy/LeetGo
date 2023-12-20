package main

func maximumSumOfHeights(maxHeights []int) int64 {
	length := len(maxHeights)
	// prefix[i]表示以i为结尾的最大值, [0,i]
	// suffix[i]表示以i为开头的最大值, [i, n-1]
	prefix, suffix := make([]int, length), make([]int, length)
	// 递增栈
	var stack []int

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

		// 这里的height是必须要减去的, 因为prefix[i]和suffix[i]都包含了height, 但是height只能被计算一次
		ret = max(ret, suffix[i]+prefix[i]-height)
	}
	return int64(ret)
}
