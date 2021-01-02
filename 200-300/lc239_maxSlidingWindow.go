package main

func maxSlidingWindow(nums []int, k int) []int {
	var ln = len(nums)
	if ln == 0 {
		return nil
	}
	// 维护一个队列, 对头索引对应的元素最大
	var queue = make([]int, 0, k)
	var res = make([]int, 0, ln)
	for i, v := range nums {
		for len(queue) > 0 && v >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)

		// 如果已经超过了窗口大小, 就出队
		if queue[0] == i-k {
			queue = queue[1:]
		}

		if i >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}

	return res
}
