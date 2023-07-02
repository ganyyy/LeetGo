//go:build ignore

package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumSum(arr []int) int {
	// dp0: 不删除任何元素的arr[:i+1]中最大的子数组和
	//      当前元素不删除, 可以采用 max(dp0, 0) + arr[i]
	// dp1:   删除一次元素的arr[:i+1]中最大的子数组和
	//      可以删arr[i], 也可以不删除arr[i]而删除前边的某个值 max(dp0, dp1+arr[i])

	// dp乃一生之敌, 状态转移方程更是如此
	dp0, dp1, res := arr[0], 0, arr[0]
	for i := 1; i < len(arr); i++ {
		dp0, dp1 = max(dp0, 0)+arr[i], max(dp1+arr[i], dp0)
		res = max(res, max(dp0, dp1))
	}
	return res
}
