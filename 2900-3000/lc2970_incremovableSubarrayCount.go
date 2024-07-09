package main

func incremovableSubarrayCount(a []int) int {
	n := len(a)
	i := 0
	for i < n-1 && a[i] < a[i+1] {
		i++
	}
	if i == n-1 { // 每个非空子数组都可以移除
		return n * (n + 1) / 2
	}

	ans := i + 2 // 不保留后缀的情况，一共 i+2 个
	// 枚举保留的后缀为 a[:X+1] + a[j:]
	// 因为要求是连续的, 所以保留后缀的前提是后缀本身得需要是递增的, 所以找到第一个非递增的位置就需要停止迭代
	for j := n - 1; j == n-1 || a[j] < a[j+1]; j-- {
		for i >= 0 && a[i] >= a[j] {
			// 当 i == -1 时, 说明前缀 a[:j] 都可以移除, 这是一个整体 此时 i+2 == 1
			i--
		}
		// 可以保留前缀 a[:i+1], a[:i], ..., a[:0] 一共 i+2 个
		// fmt.Println(a[i+1:j])
		ans += i + 2
	}
	return ans
}
