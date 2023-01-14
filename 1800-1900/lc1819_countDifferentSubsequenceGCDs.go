package main

func countDifferentSubsequenceGCDs(nums []int) (ans int) {
	maxVal := 0
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	occur := make([]bool, maxVal+1)
	for _, num := range nums {
		occur[num] = true
	}
	// 统计序列的公约数
	// 不如统计公约数属于那些序列
	for i := 1; i <= maxVal; i++ {
		subGcd := 0
		// 枚举所有的公约数, 和他们的倍数
		for j := i; j <= maxVal; j += i {
			if occur[j] {
				if subGcd == 0 {
					subGcd = j
				} else {
					// gcd需要一直更新, 比如 [4,8] -> 4, [4,8,10]-> 2
					// 如果i==2, 他是 [4,8,10]的最大公约数, 而不是 [4,8]的最大公约数
					// 所以需要不停的迭代最小的gcd(增加子序列的数字的数量)
					subGcd = gcd(subGcd, j)
				}
				// i是最小值, 因为i是单调递增的
				// 这里统计的就是i为子序列的公约数的数量
				if subGcd == i {
					// 因为是统计的是不同的最大公约数, 所以找到一次就退出即可
					ans++
					break
				}
				// 小了, 还有必要继续统计嘛?
				if subGcd < i {
					break
				}
			}
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func gcd(num1, num2 int) int {
	for num1 != 0 {
		num1, num2 = num2%num1, num1
	}
	return num2
}
