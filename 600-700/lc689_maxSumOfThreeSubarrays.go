package main

func maxSumOfThreeSubarrays(nums []int, k int) (ans []int) {
	// sum1, sum2, sum3 分别指的是当前三个窗口的累加和
	// maxSum1, maxSum12, maxTotal 分别指的是历史窗口1的最大值, 窗口1+2的最大值, 窗口1+2+3的最大值
	// maxSum1Idx是当窗口1最大时的起始位置
	// maxSum12Idx1表示的是窗口1+2最大时, 窗口1的起始位置
	// maxSum12Idx2表示的是窗口1+2最大时, 窗口2的起始位置

	// maxSum1动态计算, 计算maxSum12时, 取历史maxSum1和sum2计算, 保证不会有重叠
	var sum1, maxSum1, maxSum1Idx int
	var sum2, maxSum12, maxSum12Idx1, maxSum12Idx2 int
	var sum3, maxTotal int
	// 分别从0, k, 2k 开始迭代整个数组, 求出三个窗口的最大值
	var a [3]int
	for i := k * 2; i < len(nums); i++ {
		// 加上右端点的值
		sum1 += nums[i-k*2]
		sum2 += nums[i-k]
		sum3 += nums[i]
		if i >= k*3-1 {
			// 数量满足k个, 计算最大值
			if sum1 > maxSum1 {
				maxSum1 = sum1
				maxSum1Idx = i - k*3 + 1
			}
			if maxSum1+sum2 > maxSum12 {
				maxSum12 = maxSum1 + sum2
				maxSum12Idx1, maxSum12Idx2 = maxSum1Idx, i-k*2+1
			}
			if maxSum12+sum3 > maxTotal {
				maxTotal = maxSum12 + sum3
				a[0], a[1], a[2] = maxSum12Idx1, maxSum12Idx2, i-k+1
			}
			// 去掉左端点的值
			sum1 -= nums[i-k*3+1]
			sum2 -= nums[i-k*2+1]
			sum3 -= nums[i-k+1]
		}
	}
	return a[:]
}
