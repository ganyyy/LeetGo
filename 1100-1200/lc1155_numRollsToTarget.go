package main

func numRollsToTarget(n int, k int, target int) int {
	// 当前迭代的结果

	var current = make([]int, target+1)
	current[0] = 1
	// 上一次迭代的结果
	// var last = make([]int, target+1)
	// last[0]=1

	const MOD = 10e8 + 7

	for num := 1; num <= n; num++ {
		// num枚色子投出的点数 ∈ [num,min(target, num*k)]
		maxTotal := min(target, num*k)
		for total := maxTotal; total >= 0; total-- {
			// 这一步很关键: 每次计算时, 都要将当前位置的值清空
			current[total] = 0
			// 1枚色子的有效上限 ∈ [1,min(total, k)]
			maxPoint := min(total, k)
			for point := 1; point <= maxPoint; point++ {
				// 为什么可以复用一个数组? 因为每次计算时, 都是从后往前计算的
				// 并且后态的计算结果只依赖于前态的计算结果, 且后态的结果不会影响前态的计算结果
				current[total] = (current[total] + current[total-point]) % MOD
			}
		}
		// last, current = current, last
	}

	return current[target] % MOD
}
