package main

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	prefix := make([]int, n)
	suffix := make([]int, n)
	stack1, stack2 := []int{}, []int{}

	// 左侧非递减
	for i := 0; i < n; i++ {
		for len(stack1) > 0 && maxHeights[i] < maxHeights[stack1[len(stack1)-1]] {
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack1) == 0 {
			prefix[i] = (i + 1) * maxHeights[i]
		} else {
			last := stack1[len(stack1)-1]
			prefix[i] = prefix[last] + (i-last)*maxHeights[i]
		}
		stack1 = append(stack1, i)
	}

	res := 0
	//
	for i := n - 1; i >= 0; i-- {
		for len(stack2) > 0 && maxHeights[i] < maxHeights[stack2[len(stack2)-1]] {
			stack2 = stack2[:len(stack2)-1]
		}
		if len(stack2) == 0 {
			suffix[i] = (n - i) * maxHeights[i]
		} else {
			last := stack2[len(stack2)-1]
			suffix[i] = suffix[last] + (last-i)*maxHeights[i]
		}
		stack2 = append(stack2, i)
		res = max(res, prefix[i]+suffix[i]-maxHeights[i])
	}
	return int64(res)
}
