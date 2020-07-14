package main

func numTrees(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	// 对于 n 而言, 可以分别从 1,2,3...n作为根节点
	// 当1是根节点时, 左子树是0个, 右子树为 n-1个
	// 当2是根节点时, 左子树是1个, 右子树为 n-2个
	// ...
	// 当n-1是根节点时,左子树是n-2个,右子树为 1个
	// 当n是根节点时, 左子树是n-1个, 右子树为 0个
	// 所以 G(n) = G(0)*G(n-1) + G(1)*G(n-2) + ... + G(n-2)*G(1) + G(n-1)*G(0)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
