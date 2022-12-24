package main

func minimumBoxes(n int) int {
	//curr记录目前所放置的盒子数，i为当前第j层的盒子数，j为层数
	curr, i, j := 0, 0, 0

	// 可以理解为: 每次添加的都是最下边的一层
	// 1
	// 1 1
	// 1 1 1 <- 这是新增的那一层

	for curr < n {
		//层数加一
		j++
		//计算当前层的盒子数
		i += j
		//计算总共放置的盒子数
		curr += i
	}

	//如果正好等于，则i就是最底层所放置的盒子数
	if curr == n {
		return i
	}

	//不相等，就就接着放，可以尝试放1，2，3，4，知道curr > n
	//因为curr > n，说明放多了，要减去放的i
	curr -= i
	//因为j层的i不符合，所以退回到上一层的i
	i -= j
	//j此时代表，继续放置盒子的第几次，第i次可以放i个
	j = 0

	for curr < n {
		j++
		curr += j
	}
	//一共放了j次，最底层放了j个盒子，直接加上之前的i返回
	return i + j
}
