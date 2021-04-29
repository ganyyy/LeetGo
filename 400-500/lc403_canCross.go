package main

func canCross(stones []int) bool {

	var ln = len(stones)

	// dp[i][j] 表示第i个石头可以跳j步
	var dp = make([][]bool, ln)
	for i := range dp {
		// 为啥是ln+1 呢? 实际上可跳的范围应该是[0,ln-1], 这里+1是为了方便计算, 所以整体范围提升到了[1, ln]
		dp[i] = make([]bool, ln+1)
	}

	// 第0个石头可以跳1步
	dp[0][1] = true

	for i := 1; i < ln; i++ {
		var flag bool

		for j := i - 1; j >= 0; j-- {
			// i和j之间的差值, 决定了j能跳的步数
			// 这个步数只有在小于 i的情况下才有意义
			var diff = stones[i] - stones[j]

			// 如果差值过大, 说明跳不过去
			if diff > i {
				break
			}
			// 能跳过去的情况下, 差值必定在 < ln
			// 因为i的最大值是ln-1

			if dp[j][diff] {
				dp[i][diff] = true
				dp[i][diff+1] = true
				dp[i][diff-1] = true
				flag = true
			}
		}
		// 到了最后一步, 如果跳不过去说明不存在合理的路径
		if i == ln-1 && !flag {
			return false
		}
	}
	return true
}

func canCrossBigLao(stones []int) bool {
	// 第二块石头只能是1, 否则直接返回错误即可
	if stones[1] > 1 {
		return false
	}
	// 记录每块石头之间距离差值
	s := make([]int, len(stones)-1)
	for i := 0; i < len(stones)-1; i++ {
		s[i] = stones[i+1] - stones[i]
	}
	// 从后向前走?
	// last是当前石头到末尾石头之间的差值
	last := s[len(s)-1]
	// sum是总的步数?
	sum := 0
	for i := len(s) - 2; i >= 0; i-- {
		// 这一个不太可能会出现吧..
		if i == 0 && last > 2 {
			return false
		}
		if s[i]+sum+1 >= last && s[i]+sum-1 <= last {
			// 如果i能跳到i+1上
			last = s[i] + sum
			sum = 0
		} else if s[i]+s[i-1]+sum >= last {
			// 跳过i, 从i-1 -> i+1
			last += s[i]
		} else {
			// 保留一下累计的值, 等待 < i 的某个数进行调用
			sum += s[i]
		}
	}
	return sum == 0
}
