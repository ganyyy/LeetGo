package main

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	var sum int

	for i, v := range nums {
		// 反转负数
		if v < 0 && k > 0 {
			sum += -v
			k--
			continue
		}
		// 没有剩余的反转次数
		if k <= 0 {
			sum += v
			continue
		}
		// 如果首位就大于0, 那么直接反转k次最小值
		if i == 0 {
			if k&1 == 0 {
				sum += v
			} else {
				sum -= v
			}
			k = 0 // 清空即可
			continue
		}
		// 非首位的情况,
		// 对比大于0的最小值和小于0的最大值, 取二者绝对值中的较小值进行反转
		// 此时最大的负数已经 **被反转了**
		// 所以如果选择负数, 那么就需要在总和中减去两次/零次;
		// 如果选择整数, 就会 加一次/减一次

		var add = -nums[i-1]

		if add >= v {
			// 选择 v 进行反转
			if k&1 == 0 {
				sum += v
			} else {
				sum -= v
			}
		} else {
			sum += v
			if k&1 == 1 {
				sum -= add * 2
			}
		}
		// 本次计算完成后, 直接清空计数即可
		k = 0
	}

	// 末尾还要看看有没有剩余的
	if k > 0 {
		if k&1 == 1 {
			sum += nums[len(nums)-1] * 2
		}
	}

	return sum
}

func largestSumAfterKNegationsGood(nums []int, k int) int {
	sort.Ints(nums)

	var sum int
	var minVal = math.MaxInt32
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for _, v := range nums {
		if v < 0 && k > 0 {
			k--
			v = -v
		}
		sum += v
		minVal = min(minVal, v)
	}

	// 这种情况只有负数的数量小于k时才会出现.
	// 此时相当于选取了最大的负数(最小的绝对值)进行多次反转
	// 如果剩余奇数次反转, 说明最终表现为负数.
	// 因为之前将其当成正数加入到了结果中, 所以需要减去两次
	if k > 0 && k&1 == 1 {
		sum -= 2 * minVal
	}
	return sum
}
