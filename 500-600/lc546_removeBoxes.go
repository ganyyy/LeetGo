package main

func removeBoxes(boxes []int) int {
	// dp[i][j][k] 表示 在[i,j]部分能拿到的最大得分
	// k 表示在[0,i]之间可以结合并一起消失的盒子的个数
	// ... 1,2,2,1,1,2,2 ...
	// 以上表示的一个[i,j]的区间
	// 对于第i个数而言, 它可以和和前边相同的数(k个)一起消失, 也可以和后边的相同的数一起消失

	// helper 接受 i, j, k三个参数  i,j表示边界, k表示 前边有几个要和当前值 一起消失的盒子的个数
	// 返回[i,j]区间内的最大得分
	var helper func(i, j, k int) int

	// 对于当前 i 而言,
	// 1. 如果和前边的K个 boxes[i] 一起消失,
	//      得分为 (k+1)^2 + helper(i+1,j, 0)
	// 2. 如果和后边的 boxex[i] 一起消失的话
	//      - 和后边的第1个 boxes[i](m1)  一起消失, 得分为 helper(i+1,m1-1, 0) + helper(m1, j, k+1)
	//      - 和后边的第2个 boxes[i](m2)  一起消失, 得分为 helper(i+1,m2-1, 0) + helper(m2, j, k+1)
	//      - ....
	//      - 和后边的第N个 boxes[i](mn)  一起消失, 得分为 helper(i+1,mn-1, 0) + helper(mn, j, k+1)
	// 3. 聚合上述的所有情况, 取最大值

	// 根据题目要求, 一次性分配这么多的空间
	var dp [101][101][101]int

	helper = func(i, j, k int) int {
		// 递归结束
		if i > j {
			return 0
		}
		// 有缓存的值, 直接返回
		if dp[i][j][k] != 0 {
			return dp[i][j][k]
		}
		// 计算连续相同盒子的个数和分割索引
		for i < j && boxes[i] == boxes[i+1] {
			i++
			k++
		}

		// 结果1, 和前边相同的盒子进行合并
		res := (k+1)*(k+1) + helper(i+1, j, 0)
		// 结果n, 依次进行分割, 求最大值
		for m := i + 1; m <= j; m++ {
			if boxes[m] == boxes[i] {
				if v := helper(i+1, m-1, 0) + helper(m, j, k+1); v > res {
					res = v
				}
			}
		}
		dp[i][j][k] = res
		return res
	}
	// 返回整个dp的最大值
	return helper(0, len(boxes)-1, 0)
}
