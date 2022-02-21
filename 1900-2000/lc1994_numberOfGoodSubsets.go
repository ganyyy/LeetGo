package main

var primes = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

func numberOfGoodSubsets(nums []int) (ans int) {
	const mod int = 1e9 + 7
	freq := [31]int{}
	for _, num := range nums {
		freq[num]++
	}

	// 每一个质数都存在两种可能: 选或者不选
	// 所以整体的DP范围会扩大到两倍
	f := make([]int, 1<<len(primes))

	// 1 默认存在一次不选的可能性
	f[0] = 1
	// 1 可以添加到任意子集中, 可选可不选, 所以对应的可能性为两倍关系
	for i := 0; i < freq[1]; i++ {
		f[0] = f[0] * 2 % mod
	}
next:
	for i := 2; i < 31; i++ {
		// 没出现的数不考虑
		if freq[i] == 0 {
			continue
		}

		// 检查 i 的每个质因数是否均不超过 1 个
		subset := 0
		for j, prime := range primes {
			// 跳过完全平方数
			if i%(prime*prime) == 0 {
				continue next
			}
			// 记录质因数
			if i%prime == 0 {
				subset |= 1 << j
			}
		}

		// 动态规划, 状态转移方程:
		// 每一个数都是可选可不选的
		for mask := 1 << len(primes); mask > 0; mask-- {
			if mask&subset == subset { // mask包含的值的范围 >= subset
				// f[mask]表示不进行组合
				// f[mask^subset]表示进行组合
				f[mask] = (f[mask] + f[mask^subset]*freq[i]) % mod
			}
		}
	}

	for _, v := range f[1:] {
		ans = (ans + v) % mod
	}
	return
}
