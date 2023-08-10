package main

func lastStoneWeightII(stones []int) int {
	// 如何转换呢..?

	// 每次都崩了值最大的, 差值最小的的一组

	// 排序, [1,1,2,4,7,8]

	// 不管怎么算, 都要先把大的给碰了. 不然砰哪个都不能保证会比最大-次大

	// 76 / 86
	// var ln = len(stones)
	// for ln != 1 {
	//     sort.Ints(stones)
	//     stones[ln-2] = stones[ln-1]-stones[ln-2]
	//     stones = stones[:ln-1]
	//     ln--
	// }
	// return stones[0]

	// 每个数字选or 不选, 保证差值最小

	var total int
	for _, v := range stones {
		total += v
	}

	// 状态转移方程是?

	// total - 2 * 正子集 的最小值
	// 又是0/1背包. 我还是太菜了
	// 将集合分为 sum(P), sum(N), 最终结果就是求 min(abs(sum(P)-sum(N)))
	// sum(P) + sum(N) = total
	// min(total - 2*sum(P))

	// 想要保证 sum(P) + sum(N) == total, 同时还要保证 (sum(P)-sum(N))最小
	// 那么很冥想, sum(P)的最大值不会超过total/2. 因为 N和P之间本身是互补的关系!
	// 背包最大值是 total / 2(?) --划重点, 要考的--

	var capacity = total >> 1
	// dp[i]是背包容量为i时, 可以放入的石头的最大价值???
	var dp = make([]int, capacity+1)
	for _, v := range stones {
		// 就是一个0/1背包. 选或者不选而已
		for j := capacity; j >= v; j-- {
			//  v<=j<=capacity, 0<=j-v <= capacity
			dp[j] = max(dp[j], dp[j-v]+v)
		}
	}

	return total - 2*dp[capacity]
}
