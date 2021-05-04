package main

import (
	"fmt"
	"sort"
)

func deleteAndEarnOvertime(nums []int) int {
	// 先排个序呗..

	// 不出意外的超时了

	// 想想怎么搞记忆化/dp

	sort.Ints(nums)

	// 转换思路, 先统计每个数字的个数, 方便进行排除处理

	var count [][2]int

	count = append(count, [2]int{nums[0], 1})
	var cur int
	for i := 1; i < len(nums); i++ {
		if nums[cur] != nums[i] {
			count = append(count, [2]int{nums[i], 1})
			cur = i
		} else {
			count[len(count)-1][1]++
		}
	}

	fmt.Println(count)

	// dfs 算一个最大值
	var res int
	var dfs func([][2]int, int)
	dfs = func(count [][2]int, pre int) {
		if len(count) == 0 {
			if pre > res {
				res = pre
			}
			return
		}
		if len(count) == 1 {
			pre += count[0][0] * count[0][1]
			if pre > res {
				res = pre
			}
			return
		}

		for i := 0; i < len(count); i++ {
			var tmp = make([][2]int, len(count))
			copy(tmp, count)
			var cur = i
			if cur < len(tmp)-1 {
				if count[cur+1][0] == count[cur][0]+1 {
					tmp = append(tmp[:cur+1], tmp[cur+2:]...)
				}
			}
			if cur > 0 {
				if count[cur-1][0] == count[cur][0]-1 {
					tmp = append(tmp[:cur-1], tmp[cur:]...)
					// 去掉前边的数据后, 需要重新计算一下i所处的位置
					cur--
				}
			}

			// 全加上呗, 没必要再来一次了
			var del = tmp[cur]
			var val = del[0] * del[1]
			tmp = append(tmp[:cur], tmp[cur+1:]...)
			fmt.Printf("src:%v, delete:%v, after:%v, pre:%v, cur:%v\n", count, del, tmp, pre, pre+val)
			dfs(tmp, pre+val)
		}
	}

	dfs(count, 0)

	return res
}

func deleteAndEarn(nums []int) int {

	sort.Ints(nums)

	// 转换思路, 先统计每个数字的个数, 方便进行排除处理

	var count [][2]int

	count = append(count, [2]int{nums[0], 1})
	var cur int
	for i := 1; i < len(nums); i++ {
		if nums[cur] != nums[i] {
			count = append(count, [2]int{nums[i], 1})
			cur = i
		} else {
			count[len(count)-1][1]++
		}
	}

	// 还可以进一步压缩, 因为只用到了两个值!

	fmt.Println(count)

	// dp 需要记录两个值吗? 一个删除的, 一个不删除的..?
	var dp = make([]int, len(count)+1)

	// dp[i] 表示到当前位置的最大计算结果
	// 每一个位置可以选择删除或者不删除
	// dp[i] = max()

	dp[1] = getVal(count[0])

	for i := 1; i < len(count); i++ {
		if count[i-1][0] == count[i][0]-1 {
			// dp[i-1] 表示删除前一个数, dp[i]表示删除当前数
			dp[i+1] = max(dp[i-1]+getVal(count[i]), dp[i])
		} else {
			dp[i+1] = dp[i] + getVal(count[i])
		}
	}

	fmt.Println(dp)

	return dp[len(count)]
}

func deleteAndEarnZip(nums []int) int {

	sort.Ints(nums)

	// 转换思路, 先统计每个数字的个数, 方便进行排除处理

	var count [][2]int

	count = append(count, [2]int{nums[0], 1})
	var cur int
	for i := 1; i < len(nums); i++ {
		if nums[cur] != nums[i] {
			count = append(count, [2]int{nums[i], 1})
			cur = i
		} else {
			count[len(count)-1][1]++
		}
	}

	// 还可以进一步压缩, 因为只用到了两个值!

	fmt.Println(count)

	// dp[i] 表示到当前位置的最大计算结果
	// 每一个位置可以选择删除或者不删除
	// dp[i] = max()

	var a, b = 0, getVal(count[0])

	for i := 1; i < len(count); i++ {
		if count[i-1][0] == count[i][0]-1 {
			// dp[i-1] 表示删除前一个数, dp[i]表示删除当前数
			a = max(a+getVal(count[i]), b)
		} else {
			a = b + getVal(count[i])
		}
		a, b = b, a
	}

	fmt.Println(a, b)

	return b
}

func getVal(pair [2]int) int {
	return pair[0] * pair[1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(deleteAndEarnZip([]int{12, 32, 93, 17, 100, 72, 40, 71, 37, 92, 58, 34, 29, 78, 11, 84, 77, 90, 92, 35, 12, 5, 27, 92, 91, 23, 65, 91, 85, 14, 42, 28, 80, 85, 38, 71, 62, 82, 66, 3, 33, 33, 55, 60, 48, 78, 63, 11, 20, 51, 78, 42, 37, 21, 100, 13, 60, 57, 91, 53, 49, 15, 45, 19, 51, 2, 96, 22, 32, 2, 46, 62, 58, 11, 29, 6, 74, 38, 70, 97, 4, 22, 76, 19, 1, 90, 63, 55, 64, 44, 90, 51, 36, 16, 65, 95, 64, 59, 53, 93}))
}
