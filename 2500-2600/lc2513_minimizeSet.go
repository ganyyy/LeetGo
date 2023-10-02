package main

import "sort"

func minimizeSet(d1, d2, uniqueCnt1, uniqueCnt2 int) int {
	lcm := d1 / gcd(d1, d2) * d2

	/*
		[1, n]中的数可以分为四部分:
		1. 可以整除d1的数(n/d1)
		2. 可以整除d2的数(n/d2)
		3. 既可以整除d1, 也可以整除d2的数(n/lcm)
		4. 既不可以整除d1, 也不可以整除d2的数(n - n/d1 - n/d2 + n/lcm)

		属于d1独占的部分包括: 2 - 3
		属于d2独占的部分包括: 1 - 3
	*/

	// 根据题意, 可得:
	// arr1中不包含可以整除d1的元素
	// arr2中不包含可以整除d2的元素
	// 设d1/d2的最小公倍数为lcm
	// 假设x是数组中的最大元素, 那么相当于元素会从[1,x]中选取, 上限是x个
	// 由此可得: len(arr1)+len(arr2) == arr1中独占的部分+arr2中独占的部分+两个数组共享(按需分配)的部分
	// arr1独占: x/d2 - x/lcm
	// arr2独占: x/d1 - x/lcm
	// 二者共享: x - x/d1 - x/d2 + x/lcm (这部分给谁都行, 看uniqueCnt)
	// 理想情况下: 二者共享的部分 = 数组1去掉独占后的部分+数组2去掉独占后的部分
	// 如果共享的数更多, 说明还有下降空间; 否则就需要提升
	// 共享的数多说明x偏大

	// 关于上界问题: 由题意可知, 当满足 对应可以整除d1,d2的数越多的时候, 需要的x就得更大才能满足uniqueCnt的需要
	// 所以d1,d2都是2的时候, 相当于集合中只能存在非2的倍数的数, 那么最大值就是2*uniqueCnt
	// 当d1, d2 >= 2的时候, 所需要的数量只会更小, 所以最大值就是2*uniqueCnt-1
	return sort.Search((uniqueCnt1+uniqueCnt2)*2-1, func(x int) bool {
		// x就是两个集合中的最大值

		// dvd1: 可以整除d1的数
		// dvd2: 可以整除d2的数
		// pub: 既可以整除d1, 也可以整除d2的数
		dvd1, dvd2, pub := x/d1, x/d2, x/lcm

		// 数组1剩余的个数
		left1 := max(uniqueCnt1-(dvd2-pub), 0)
		// 数组2剩余的个数
		left2 := max(uniqueCnt2-(dvd1-pub), 0)
		// 二者共有的部分
		common := x - (dvd1 + dvd2 - pub)

		return common >= left1+left2
	})
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
