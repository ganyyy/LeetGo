package main

func bestRotation(nums []int) int {
	// 我的评价是: 直接G
	// 题解都看不懂

	// 按照题意:
	// 轮调的方向是向左
	// nums[i] = x, 得分的区间是 [x, n-1] n-x个值,
	// 不得分的区间是[0, x-1], x个值
	// 假设当前已经轮询了k轮, 那么 x 处于 (i-k+n)%n
	// 如果想要得分, 即(i-k+n)%n >= x  ==>  k <= (i-x+n)%n
	// 当x取0时, 有最小下界 k >= (i+1)%n

	// 核心就是差分数组, 将每一个位置可能带来的加分/减分带入进去
	// 这个上下界不太好理解

	n := len(nums)
	diffs := make([]int, n)
	for i, num := range nums {
		low := (i + 1) % n
		high := (i - num + n + 1) % n
		diffs[low]++
		diffs[high]--
		if low >= high {
			diffs[0]++
		}
	}

	// 最后计算总和
	score, maxScore, idx := 0, 0, 0
	for i, diff := range diffs {
		score += diff
		if score > maxScore {
			maxScore, idx = score, i
		}
	}
	return idx
}
