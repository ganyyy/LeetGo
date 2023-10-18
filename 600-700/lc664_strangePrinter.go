package main

import "math"

func strangePrinter(s string) int {
	// DP啊, 想不出来状态转移方程都白搭

	// 0 <= left, right <= len(s)-1
	// DP[left][right]表示打印s[left:right+1]所需要的最小次数
	length := len(s)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}
	for left := length - 1; left >= 0; left-- {
		// s[left:left+1]至少要一次打印
		dp[left][left] = 1
		for right := left + 1; right < length; right++ {
			if s[left] == s[right] {
				// 如果s[left] == s[right], 那么就可以理解为 s[right]是通过s[left:right-1]一次打印过来的
				// 此时需要的总的打印次数就是 dp[left][right-1]所需要的次数
				// 比如 abc[a], 可以理解为先打印的 aaaa, 然后再打印bc. 此时消耗的次数就是打印 abc 所需要的次数
				// 那么当s[left] == s[right]时, dp[left][right] = dp[left][right-1]
				// 因为依赖于dp[left][right-1], 所以j需要从前往后遍历
				dp[left][right] = dp[left][right-1]
			} else {
				// 此时可以理解为需要将s[left:right+1]分为两部分进行打印
				// 取一个中间值K, 将s[left:right+1]分为 s[left:mid+1] 和 s[mid+1:right+1]
				// 通过遍历的形式获取满足最小值的k
				dp[left][right] = math.MaxInt64
				for mid := left; mid < right; mid++ {
					dp[left][right] = min(dp[left][right], dp[left][mid]+dp[mid+1][right])
				}
			}
		}
	}
	return dp[0][length-1]
}

func strangePrinter2(s string) int {
	length := len(s)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}

	for count := 1; count <= length; count++ {
		// count: 判断字符串的长度
		// left,right: 字串的左右端点(包含, s[left:right+1])
		var left, right int
		for left = 0; left+count-1 < length; left++ {
			right = left + count - 1
			// 先无脑加一次打印. 因为不管是否和中间某个字母相同, 都允许这次单独打印
			// 注意方向: xxxx[XXXX] -> xxx[xXXXX], 往左进位
			var initVal = 1
			if left+1 <= right {
				initVal += dp[left+1][right]
			}
			dp[left][right] = initVal
			for mid := left + 1; mid <= right; mid++ {
				if s[left] == s[mid] {
					//  axxxaxxxx ->
					// [axxxa][xxxx]
					// 注意: [axxxa]这部分的实际打印次数是[axxx],
					// 因为后补的[a]可以理解为归属于[aaaaa]这一次打印的结果
					dp[left][right] = min(dp[left][right], dp[left][mid-1]+dp[mid+1][right])
				}
			}
		}
	}
	return dp[0][length-1]
}
