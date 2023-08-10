package main

func getSteps(cur, n int) (steps int) {
	// 1, 1, 1
	// 10, 19, 11
	// 100, 199 111
	first, last := cur, cur
	for first <= n {
		steps += min(last, n) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return
}

func findKthNumber(n, k int) int {
	cur := 1 // 从第一层开始, 根节点分别为 1,2,3,4,5
	// 下一层分别为 10,11,12...   20,21,22...   30,31,32...
	// 个位数的情况下,
	k--
	// 整体的查找方向, 就是 由左上到右下的方向查找
	for k > 0 {
		steps := getSteps(cur, n) // 先确定大概的层数, 在确定对应的步数
		if steps <= k {           // 如果k >= steps, 说明要找下一个根
			k -= steps
			cur++
		} else { // 否则, 就需要找到具体的层
			cur *= 10
			k--
		}
	}
	return cur
}

func findKthNumber2(n int, k int) int {
	var cur = 1 // 第一个数
	k--         // 算上第一个

	for k > 0 {
		var step = getSteps(cur, n)
		if step <= k {
			// fmt.Println("less K:", step, cur, k)
			k -= step // 第k个数不在当前基数对应的子树中,
			cur++     // 查找下一个基数
		} else {
			// 这种情况啥时候会出现呢?
			// 因为getSteps获取的到达某一层首个数字经历的步数
			// 举个栗子: 就是某一层在前n中, 但是k 不属于下一个基数的子树
			// fmt.Println("big K:", step, cur, k)
			cur *= 10 // 第k个数在当前基数中的某个子树中,
			k--       // 取当前层的第0个, 所以减少了一步
		}
	}
	return cur
}
