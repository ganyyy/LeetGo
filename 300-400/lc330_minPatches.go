package main

func minPatches(nums []int, n int) int {
	// 一个逐步扩大的过程

	// 两个下标. 一个委会 从[1,n], 一个维护 nums[i]

	// res 是最终结果, curRange 是当前可以遍历到的最大范围
	var res, curRange int
	var m = len(nums)
	for i, pos := 1, 0; i <= n; {
		// 如果pos 已经超过了nums大小, 说明没有候选的目标了
		// 或者 i < nums[pos], 此时 说明i还在前边, 且此时没有满足计算出i的组合
		// 此时需要添加一个数, 这个数就是i
		if pos >= m || i < nums[pos] {
			// 需要添加一个数
			res++
			// 当前可以计算到的最大的范围
			// 此时可以保证curRange之前的数都可以计算出来
			// 那么 可以计算的数的最大范围就到了 curRange + i上
			curRange += i
		} else {
			// 此时, nums[pos] >= i
			// 直接使用这个数, 重新计算能达到的最大范围
			curRange += nums[pos]
			// 指向下一个候选的数
			pos++
		}
		// 此时, <= curRange的所有的数都存在一个组合可以计算出来
		// 所以下一步就从 curRange+1开始计算
		i = curRange + 1
	}

	return res
}
