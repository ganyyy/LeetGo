package main

import "strconv"

func findIntegersBad(n int) int {
	var ret int
	for i := 0; i <= n; i++ {
		if i&(i>>1) == 0 {
			ret++
		}
	}
	return ret
}

func findIntegers(n int) (ans int) {
	// 这个做法很...

	/*
	   还是有点想不明白
	   简单而言, 就是把数字分解成一个由 0, 1组成的字典树
	   很明显, 任意节点要么为0, 要么为1
	   其包含的 非连续1的数字个数就是从叶子节点上升到该根节点中不存在连续1的路径数量
	   如果一个根节点存在两个子节点, 那么其左子树(0)一定是一个满二叉树, 右子树一定是一个完全二叉树
	   如果只有一个子节点, 那么该子节点一定是一个完全二叉树

	   以根节点为0的子树而言, 其高度为t
	   那么, 左右子树高度为 t-1, 左子树的数据可以直接加上, 右子树只能加入其左半部分(高度为t-2)
	   所以, dp[t] = dp[t-1] + dp[t-2]

	*/

	// 标记

	dp := [31]int{1, 1}
	for i := 2; i < 31; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	for i, pre := 29, 0; i >= 0; i-- {
		val := 1 << i
		if n&val > 0 {
			// 该位置值为1, 那么就相当于加上该高度对应的根节点为0的 非连续1 的个数
			ans += dp[i+1]
			if pre == 1 {
				// 如果出现了连续的两个1, 那么就不需要再继续查找了. 因为对应的左半部分已经在之前计算过了
				break
			}
			pre = 1
		} else {
			pre = 0
		}
		if i == 0 {
			// 叶子节点, 需要单独处理
			ans++
		}
	}
	return
}

func findIntegers2(n int) int {
	s := strconv.FormatInt(int64(n), 2)
	m := len(s)
	// 每一位有两种可能: 0/1
	dp := make([][2]int, m)
	for i := range dp {
		dp[i] = [2]int{-1, -1}
	}
	// 数位dp
	var f func(int, int8, bool) int
	f = func(i int, pre1 int8, isLimit bool) (res int) {
		if i == m {
			return 1
		}
		if !isLimit {
			dv := dp[i][pre1]
			if dv >= 0 {
				return dv
			}
			defer func() { dp[i][pre1] = res }()
		}
		up := 1
		if isLimit {
			up = int(s[i] & 1)
		}
		// 无论前边是什么, 填0总是对的
		// isLimit应该怎么理解呢? 可以认为这一次迭代是不是该位置的上限.
		// 因为这里只有0/1, 所以是二元的. 但是, 在十进制的框架下, 非limit可选0-9, limit只能选到s[i]作为其上限
		res = f(i+1, 0, isLimit && up == 0) // 填 0
		if pre1 == 0 && up == 1 {           // 可以填 1
			// 前边是0的情况下, 才可以填1
			res += f(i+1, 1, isLimit) // 填 1
		}
		return
	}
	return f(0, 0, true)
}
