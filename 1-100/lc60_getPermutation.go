package main

import (
	"fmt"
	"sort"
)

func getPermutation(n int, k int) string {
	// 全排
	res := make([]string, 0)
	remain := make([]byte, n)
	for i := 1; i <= n; i++ {
		remain[i-1] = byte(i) + '0'
	}
	dfs(make([]byte, 0, n), remain, &res)
	sort.Strings(res)
	// 返回结果
	return res[k-1]
}

func dfs(now, remain []byte, res *[]string) {
	if len(remain) == 1 {
		// 注意一下, 切片的形式必须存下来
		*res = append(*res, string(append(now, remain[0])))
		return
	}
	for i := 0; i < len(remain); i++ {
		remain[0], remain[i] = remain[i], remain[0]
		dfs(append(now, remain[0]), remain[1:], res)
		remain[0], remain[i] = remain[i], remain[0]
	}
}

// 超时了, 直接pass

// 康托展开
func getPermutation2(n int, k int) string {
	// 最多9个数字
	bytes := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	// 对应的(0-9)!的值
	factors := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	res := make([]byte, 0, n)
	// 前边一共有k-1个排列
	k--
	for n > 0 {
		// 找出第n个数(从右向左查).
		// 对于第n个数而言, 后面的排列共有 (n-1)!种,
		// 又因为是按照字典序的全排列, 所以可以确定该数是剩下的可选的数据中的第几个
		// 如 n=5, k=96, 则k-1=95表示前边一共存在95个排列
		// 95/4! = 3表示前边已经拍完了三个数, 所以当前该位置上的数应该是4
		n--
		val := k / factors[n]
		res = append(res, bytes[val])
		// 筛选完成后删除该数
		bytes = append(bytes[:val], bytes[val+1:]...)
		// 在确定剩下的n-1个数, k %= factors[n] 表示剩下的n-1个数中要选取的第几个排列
		k %= factors[n]
	}
	return string(res)
}

// 想一下如何逆向结果?
func reversePermutation(n int, res string) int {
	fmt.Println(res)
	// 记录每个出现的数字的位置
	// 最多9个数字
	bytes := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	// 对应的(0-9)!的值
	factors := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	ret := 0
	for i := 0; i < len(res); i++ {
		// 找到该位置对应字典序的第几个
		j := 0
		for ; j < len(bytes); j++ {
			if res[i] == bytes[j] {
				// 移除已选择的数
				bytes = append(bytes[:j], bytes[j+1:]...)
				break
			}
		}
		n--
		ret += factors[n] * j
	}
	return ret + 1
}

func main() {
	n, v := 9, 1
	fmt.Println(reversePermutation(n, getPermutation2(n, v)))
}
