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
			// 不然就跳出去了...
			var diff = stones[i] - stones[j]

			// 如果差值过大, 说明无法正好跳跃到终点
			// 如果从j开始就超过了最大范围, 那么 j-1只会比这个值更大, 所以就没有继续看的必要了
			if diff > i {
				break
			}
			// 能跳过去的情况下, 差值必定在 < ln
			// 因为i的最大值是ln-1

			// 如果之前就能跳跃到这里来, 就把下一步可条约的 diff-1, diff, diff+1 置为true
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
		// s[i] 转变成了要跳到下一步所需要的步数
		s[i] = stones[i+1] - stones[i]
	}
	// 从后向前走?

	// last是跳到下一个目标点所需要的步数
	last := s[len(s)-1]

	sum := 0
	for i := len(s) - 2; i >= 0; i-- {
		if i == 0 && last > 2 {
			// last不是一直再累加吗? 不应该小于2啊
			return false
		}

		// 真尼玛难理解的解法

		/*
			1. last代表跳跃到*结尾*所需要的步数
			2. sum 中间存在的空缺的步数

			完全理解不能啊. 不看了, 反正也背不下来..
		*/

		// 从当前位置能跳的步数为
		if total := s[i] + sum; total >= last-1 && total <= last+1 {
			// 如果再允许的范围内 [last-1, last+1]
			// 那么就更新一下 跳到当前位置所需要的步数
			last = s[i] + sum
			sum = 0
		} else if s[i-1]+total >= last {
			// 这就很奇怪, 为什么要依赖于上一个位置所需要的元素判断当前的结果呢?
			// 前一步所需要的步数+当前所需要的步数+中间空缺的步数 比 总体所需要的步数要大(last+1, )
			last += s[i]
		} else {
			// 前一步所需要的步数+当前所需要的步数+中间空缺的步数 比 总体所需要的步数要小[0, last-1)
			sum += s[i]
		}
	}

	// 最终正确的跳到结尾的前提是不存在 剩余的步数
	return sum == 0
}

func canCrossSearch(stones []int) bool {
	// 第二种解法太过于逆天, 还是找个看得懂的吧...
	var m = make(map[int]bool, len(stones))

	var helper func(i, limit int) bool

	helper = func(i, limit int) bool {
		var key = i*1000 + limit
		if m[key] {
			// 这里只需要关注没走过的路, 走过的路无需重复计算
			return false
		} else {
			m[key] = true
		}
		for idx := i + 1; idx < len(stones); idx++ {
			var diff = stones[idx] - stones[i]
			if diff >= limit-1 && diff <= limit+1 {
				// 情况1: 能跳到下一步, 基于进行递归计算
				if helper(idx, diff) {
					return true
				}
			} else if diff > limit+1 {
				// 情况2: 超出了限制的步数, 无需继续计算(因为是递增的)
				break
			}
		}
		// 判断是不是到结尾了
		return i == len(stones)-1
	}

	return helper(0, 0)
}

// helper：上一步跳了k步，来到index处，基于此，能否到达终点
func helper(stones []int, index, k int, hashMap map[int]bool) bool {
	key := index*1000 + k // 构造唯一的key，代表当前子问题
	if hashMap[key] {     // 这个子问题之前遇到过，直接返回false
		return false
	} else { // 第一次遇到这个子问题，在map记录一下
		hashMap[key] = true
	}
	for i := index + 1; i < len(stones); i++ { //枚举出可跳的选择，从下一个石头到最后一个
		gap := stones[i] - stones[index] // 本轮迭代选择跳到i，算出第i石头到当前石头的距离
		if gap >= k-1 && gap <= k+1 {    // 这一步能跳[k-1,k+1]，距离gap在这范围内，就能尝试
			if helper(stones, i, gap, hashMap) { // 如果 基于此的递归返回true，说明能到终点
				return true // 递归压栈压到底，如果true了，则true一层向上返回
			}
		} else if gap > k+1 { // 超出范围，跳不到，它之后的石头更跳不到，不用考察了，break
			break
		} // 这是gap<k-1的情况，说明这个点离太近了，要继续下一轮迭代，看看远一点的石头
	}
	return index == len(stones)-1 // 考察完当前的所有选项，看看来到的index是不是终点
}

func main() {
	canCrossBigLao([]int{0, 1, 3, 5, 6, 8, 12, 17})
}
