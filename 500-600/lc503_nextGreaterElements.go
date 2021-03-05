package main

func nextGreaterElements(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}

	// 通过一个单调递增栈进行处理
	// 栈顶存放最小值的索引, 栈底存放最大值的索引
	var stack []int
	// 针对循环队列怎么处理?
	var res = make([]int, len(nums))

	var addRes = func(i int) {
		// 如果栈不为空, 并且栈顶元素小于当前元素, 那么就将栈顶元素出栈, 并更新结果的值
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
			var top = stack[len(stack)-1]
			res[top] = nums[i]
			stack = stack[:len(stack)-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		addRes(i)
		// 当前元素入栈
		stack = append(stack, i)
	}

	// 这里需要循环第二轮吗?
	// 这种处理方式不适合循环数组
	// 如果时一趟的还可以

	// 如何处理环呢?
	// 将这个数组和前半段联系起来, 重来一次
	if len(stack) != 0 {
		// 这个绝对是最大的了, 不用管了
		var endIdx = stack[0]
		res[endIdx] = -1
		stack = stack[1:]
		// 再来一次吧...
		for i := 0; i <= endIdx; i++ {
			addRes(i)
		}
	}

	// 到这里, 还剩下的一定都是最大值了
	// 直接赋值为-1即可
	for len(stack) != 0 {
		res[stack[0]] = -1
		stack = stack[1:]
	}

	return res
}

func example(nums []int) []int {
	// 先写一个O(n^2)的吧
	var ln = len(nums)
	var res = make([]int, len(nums))
	for i := 0; i < ln; i++ {
		var j = (i + 1) % ln
		for j != i {
			if nums[j] > nums[i] {
				break
			}
			j = (j + 1) % ln
		}
		if j == i {
			res[i] = -1
		} else {
			res[i] = nums[j]
		}
	}

	return res
}
