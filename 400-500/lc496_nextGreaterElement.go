package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 相当于在nums2查询大于当前值的下一个最大值

	var tmp = make(map[int]int, len(nums2))

	var stack []int

	for _, v := range nums2 {
		// 栈为空, 或者栈顶元素大于当前值时, 元素入栈
		// 不用考虑相等的情况
		if len(stack) == 0 || stack[len(stack)-1] > v {
			stack = append(stack, v)
		} else {
			// 持续出栈, 直到栈顶元素小于v
			// 此时出栈的所有元素的下一个值都是v
			for len(stack) != 0 && stack[len(stack)-1] < v {
				var top = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				tmp[top] = v
			}
			// 将v入栈
			stack = append(stack, v)
		}
	}

	// 原地处理, 理论上栈中还可能存在一些元素
	// 这些元素不会出现在map中, 即没有比他更大的值出现在它自身后面
	for i, v := range nums1 {
		if next, ok := tmp[v]; ok {
			nums1[i] = next
		} else {
			nums1[i] = -1
		}
	}

	return nums1

}
