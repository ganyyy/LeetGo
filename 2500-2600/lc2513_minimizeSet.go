package main

import "sort"

func minimizeSet(d1, d2, uniqueCnt1, uniqueCnt2 int) int {
	lcm := d1 / gcd(d1, d2) * d2
	// 根据题意, 可得:
	// arr1中不包含可以整除d1的元素
	// arr2中不包含可以整除d2的元素
	// 设d1/d2的最小公倍数为lcm
	// 假设x是数组中的最大元素, 那么相当于元素会从[1,x]中选取, 上限是x个
	// 由此可得: len(arr1)+len(arr2) == arr1中独占的部分+arr2中独占的部分+两个数组共享的部分
	// arr1独占: x/d2 - x/lcm
	// arr2独占: x/d1 - x/lcm
	// 二者共享: x - x/d1 - x/d2 + x/lcm (这部分给谁都行, 看uniqueCnt)
	// 理想情况下: 二者共享的部分 = 数组1去掉独占后的部分+数组2去掉独占后的部分
	// 如果共享的数更多, 说明还有下降空间; 否则就需要提升
	return sort.Search((uniqueCnt1+uniqueCnt2)*2-1, func(x int) bool {
		// 数组1剩余的个数
		left1 := max(uniqueCnt1-x/d2+x/lcm, 0)
		// 数组2剩余的个数
		left2 := max(uniqueCnt2-x/d1+x/lcm, 0)
		// 二者共有的部分
		common := x - x/d1 - x/d2 + x/lcm

		return common >= left1+left2
	})
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
