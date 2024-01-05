package main

func canSeePersonsCountBig(heights []int) []int {
	// 向右找到第一个比自身大的数
	// 中间记录最小值
	var findCount = func(heights []int) int {
		if len(heights) < 2 {
			return 0
		}
		var start = heights[0]
		var count int
		var last = -1
		for _, height := range heights[1:] {
			if last == -1 || last < height {
				last = height
				count++
			}
			if height > start {
				break
			}
		}
		return count
	}

	ret := make([]int, len(heights))
	for i := range heights {
		ret[i] = findCount(heights[i:])
	}
	return ret
}

func canSeePersonsCount(heights []int) []int {

	// var findReverse = func(stack []int, target int) int {
	// 	// [10, 5, 1], 3
	// 	// 逆序的
	// 	var left, right = 0, len(stack)
	// 	for left < right {
	// 		mid := left + (right-left)/2
	// 		if stack[mid] >= target {
	// 			left = mid + 1
	// 		} else {
	// 			right = mid
	// 		}
	// 	}
	// 	if left == 0 {
	// 		return len(stack)
	// 	}
	// 	return len(stack) - left + 1
	// }
	//
	// _ = findReverse

	// 倒序, 递减栈
	ret := make([]int, len(heights))
	var stack = make([]int, 0, len(heights)/2)
	for i := len(heights) - 1; i >= 0; i-- {
		height := heights[i]

		// 从递减栈中找到首个大于当前值的位置
		// count := findReverse(stack, height)
		var count int
		for len(stack) > 0 && stack[len(stack)-1] < height {
			count++
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			// 再加一个比当前值大的
			count++
		}
		// fmt.Println(i, stack, count)
		stack = append(stack, height)
		ret[i] = count
	}
	return ret
}
