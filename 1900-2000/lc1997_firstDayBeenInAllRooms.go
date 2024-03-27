package main

func firstDayBeenInAllRooms(nextVisit []int) int {
	const mod = 1_000_000_007
	n := len(nextVisit)
	s := make([]int, n)
	// 当到达房间i时, 肯定之前所有的房间都被访问的偶数次!
	// 当首次到达房间i的时候, 第一次访问为奇数, 因此需要回源到房间j(0 <= j <= i)
	// 此时访问房间j的次数就变成了奇数!
	// 如果需要重回i, 就意味着[j, i-1]这个区间内的所有房间都得走一次回访.
	// 对回访的定义: 第偶数次到达该房间所需要的天数. 只需要关注第一个偶数即可, 后续的操作都是可复制的.
	// 定义f[i]为回访所需要的天数, 且 nextVisit[i] = j
	//  f[i] = 2 + sum(f[k]) (k∈[j,i-1])  ①
	// 定义前缀和s[i], s[0] = 0, s[i] = sum(f[k]) (k∈[0, i-1])
	//  f[i] = 2 + s[i] - s[j] ②
	// 前缀和的递推公式可得
	//  s[i+1] = s[i]+f[i]  ③
	// ②和③联立可得
	//  s[i+1] = s[i]*2 +2 - s[j]

	// 访问所有房间, 等同于 f[0] + f[1] + f[2] + ... + f[n-2] + 1. 因为第n-1号房间只需要访问一次就行了
	// 因为第0天不需要访问, 所以总数为 sum(f[k]) (k ∈ [0,n-2]) = s[n-1]
	// 最终结果就是求s[n-1]
	// s[n-1] = s[n-2]*2+2-s[nextVisit[n-2]]
	// ...
	// s[1]   = s[0]*2 + 2-s[nextVisit[0]]
	for i, j := range nextVisit[:n-1] {
		s[i+1] = (s[i]*2 - s[j] + 2 + mod) % mod // + mod 避免算出负数
	}
	return s[n-1]
}
