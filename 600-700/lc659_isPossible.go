package main

import "fmt"

func isPossible(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	// 贪心算法
	// 一个hash存储数组中每个元素的剩余个数
	// 一个hash存储数组中每个数字作为结尾的子序列数量

	var cntMap = make(map[int]int, len(nums))
	var endMap = make(map[int]int, len(nums))

	// 统计每个数字的数量
	for _, v := range nums {
		cntMap[v]++
	}

	// 对于数组元素 x, 如果存在 x-1 组成的序列, 那么就将 x组合到x-1的序列中
	// 因为短序列越少越容易满足条件

	for _, v := range nums {
		// 如果某个数值计数已经为0, 说明以作为其他序列的一部分, 直接跳过
		if cntMap[v] == 0 {
			continue
		}
		// 首先先看一下是否存在 v-1 组成的序列
		if endMap[v-1] > 0 {
			// 如果存在,

			// 消耗一个v
			cntMap[v]--
			// 以v-1的为结尾的序列数量-1
			endMap[v-1]--
			// 就将v并入到v-1的序列中, 以v结尾的序列数量+1
			endMap[v]++
		} else {
			// 如果不存在, 就要将v作为序列的第一个元素

			// 此时需要判断一下 v+1, v+2 是否存在
			if cntMap[v+1] > 0 && cntMap[v+2] > 0 {
				// 如果存在, [v, v+1, v+2] 组成一个最短序列
				cntMap[v]--
				cntMap[v+1]--
				cntMap[v+2]--
				// v+2 作为结尾的子序列+1
				endMap[v+2]++
			} else {
				// 否则无法组成一个数量最少为3的子序列
				return false
			}
		}
	}

	// 遍历完依旧可以分割, 返回true
	return true
}

func isPossible2(nums []int) bool {
	var n = len(nums)
	// 当前值为x, 前一个值为 prev
	// dp1: 以prev为结尾长度为1的子序列
	// dp2: 以prev为结尾长度为2的子序列
	// dp3: 以prev为结尾长度>=3的子序列
	var dp1, dp2, dp3 int

	// start: x起点
	// val: x值
	// cnt: x的数量
	var start, val, cnt int
	// keep: 一次转换剩余的val的数量		= min(dp3, cnt-dp1-dp2)
	// left: 去掉给dp1和dp2之后的剩余的数量 = cnt-dp1-dp2
	var keep, left int
	for i := 0; i < n; {
		start = i
		val = nums[i]
		// 统计x的数量
		for i < n && nums[i] == val {
			i++
		}
		cnt = i - start

		if start > 0 && val != nums[start-1]+1 {
			// 如果 val和前边的数不连续
			if dp1+dp2 > 0 {
				// 如果还存在 dp1和dp2, 说明无法分割
				return false
			} else {
				// 如果dp1和dp2都不存在, 说明需要开一个新序列
				dp1 = cnt
				dp2, dp3 = 0, 0
			}
		} else {
			// 如果 val == prev+1, 说明 x可以加入到 prev组成的序列中
			if dp1+dp2 > cnt {
				// 如果对val的需求大于其数量, 说明无法分割
				return false
			}
			// 根据贪心策略, 优先添加到 数量 <= 2的prev序列中
			left = cnt - dp1 - dp2
			if dp3 > left {
				keep = left
			} else {
				keep = dp3
			}
			// 进行转置. 这代表的是以val为结尾的不同长度的序列数量

			dp3 = keep + dp2  // val的dp3 为 prev的dp2 + prev的dp3能加入的值
			dp2 = dp1         // val的dp2 为 prev的dp1
			dp1 = left - keep // val的dp1 为 剩下的val
		}
	}

	// 最终正确的标志是没有任何剩余的 dp1和dp2
	return dp1 == 0 && dp2 == 0
}

func main() {
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
}
