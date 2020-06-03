package main

/**
要求:
1. 超过K将不能抽牌
2. 每次抽牌只能抽取到[1:W] 之间的值
3. 如果结束后手中牌的和小于N就为胜利

求:
胜利的概率
*/

func new21Game(N int, K int, W int) float64 {
	// 把整体分成两段, [0:K), [K:K+W)
	// 允许抽取的最大值是 当前抽到 K-1时, 所以能抽到的最大值是 K-1+W

	// dp[x] 为 手上牌为 x时, 获胜的概率
	// 很明显, dp[x] = sum(dp[x+1: x+w])/W
	// 因为从1到W之间的概率为均等的,
	// x最大值为 K-1, 超过K就会停止抽牌
	dp := make([]float64, K+W)
	var s float64 = 0

	// 此时s保存的是 dp[K:K+W]之间的胜率和
	for i := K; i < K+W; i++ {
		if i <= N {
			dp[i] = 1
		}
		s += dp[i]
	}
	for i := K - 1; i >= 0; i-- {
		// 当 i = K-1 时, 此时可以得到的胜率就是
		// s/W
		dp[i] = s / float64(W)
		// i每移动一次, 只需要减掉最右边的在加上当前值,
		// 就能一直保持 长度为 从i开始长度为W的胜利次数的和
		s = s - dp[i+W] + dp[i]
	}
	// 最终胜利的概率就是手中一张牌没有时胜利的概率
	return dp[0]
}

func main() {

}
