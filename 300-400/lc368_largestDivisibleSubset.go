package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	var res []int
	for j := 0; j <= len(nums)-1; j++ {
		var cur = nums[j]
		var tmp = []int{cur}
		for k := j + 1; k < len(nums); k++ {
			if nums[k]%cur == 0 {
				// 只要和最大的取余为0, 那就一定能和中间的值取余为0
				// 无法过滤中间值. 错误的解法
				tmp = append(tmp, nums[k])
				cur = nums[k]
			}
		}
		if len(tmp) > len(res) {
			res = tmp
		}
	}
	return res
}

func largestDivisibleSubset2(nums []int) []int {
	sort.Ints(nums)

	var n = len(nums)
	var dp = make([]int, n)
	var mi, ml int
	//
	for i := 0; i < n; i++ {
		dp[i] = 1
		// 遍历小于 i 的每个数, 获取余数最大的那个值
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// 同步更新一下最大的结果长度和最大到达的索引位置
		if dp[i] > ml {
			ml = dp[i]
			mi = i
		}
	}

	var res = make([]int, ml)
	for i := mi; i >= 0; i-- {
		// 通过回退的模式, 依次获取这条链路上的值
		if nums[mi]%nums[i] == 0 && dp[i] == ml {
			// 这里采用了一种链路追踪的模式进行处理. 必须要实时的更新mi和ml的值
			// 否则可能会连不上
			res[ml-1] = nums[i]
			ml--
			mi = i
		}
	}
	return res
}

func largestDivisibleSubset3(nums []int) []int {
	sort.Ints(nums)

	var n = len(nums)
	var dp = make([]int, n)
	// 这里通过chain 记录了每次的位置, 所以可以直接找到前一步的值
	var chain = make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		chain[i] = -1
	}
	var mi, ml int
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if dp[i] <= dp[j] {
					dp[i] = dp[j] + 1
					chain[i] = j
				}
			}
		}
	}

	// 将比较放在外边性能更好, 因为至多比较n次. 放在里面会比较n^2次
	for k, v := range dp {
		if v > ml {
			mi = k
			ml = v
		}
	}

	var res = make([]int, 0, ml)
	for i := mi; i >= 0; i = chain[i] {
		res = append(res, nums[i])
	}

	return res
}

func main() {
	fmt.Println(largestDivisibleSubset3([]int{3, 4, 16, 8}))
}
