package main

func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}

	sum := 0
	for _, x := range nums {
		sum += x
	}
	for i := range nums {
		// 消除浮点数带来的影响
		nums[i] = nums[i]*n - sum
	}

	// 折半查找
	m := n / 2
	left := map[int]bool{}
	for i := 1; i < 1<<m; i++ {
		// i代表各种选取的可能
		// m个数一共有 2^m中选取方法
		tot := 0
		for j, x := range nums[:m] {
			// 选取当前i对应的值
			if i>>j&1 > 0 {
				tot += x
			}
		}
		// 如果为0, 那么就表示存在相关的差集也是0
		if tot == 0 {
			return true
		}
		left[tot] = true
	}

	// 右边的总和
	// 能走到这一步, 说明前半部分肯定找不到和为0的组合
	rsum := 0
	for _, x := range nums[m:] {
		rsum += x
	}
	for i := 1; i <= (1<<(n-m))-1; i++ {
		tot := 0
		for j, x := range nums[m:] {
			if i>>j&1 > 0 {
				tot += x
			}
		}
		// 这个 != 几个意思?
		// 必须要保证没有全部选取才可.
		if tot == 0 || rsum != tot && left[-tot] {
			return true
		}
	}
	return false
}
