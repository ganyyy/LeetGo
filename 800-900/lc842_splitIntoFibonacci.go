package main

import (
	"fmt"
	"math"
	"strconv"
)

func splitIntoFibonacci(S string) []int {
	var res []int

	// 标准的dfs, 一个一个查找
	var dfs func(idx int) bool

	dfs = func(idx int) bool {
		if idx == len(S) {
			// 如果到了结尾, 需要保证当前数组至少存在3个数
			return len(res) > 2
		}

		for i := idx + 1; i <= len(S); i++ {
			// 除了0本身外, 不允许以0作为数字的开头
			if S[idx] == '0' && i > idx+1 {
				break
			}
			// 超过10位肯定超过了 math.MaxInt32
			if i-idx > 10 {
				break
			}
			// 超过了 MaxInt32也要跳出
			var val, _ = strconv.Atoi(S[idx:i])
			if val > math.MaxInt32 {
				break
			}

			var ln = len(res)
			// 进去
			res = append(res, val)
			if ln < 2 {
				if dfs(i) {
					return true
				}
			} else {
				// 计算是否相等
				var add = res[ln-1] + res[ln-2]
				// 如果当前的val已经比add要大了, 没必要继续跑下去了
				// 因为越往后越大
				if add < val {
					break
				}
				if add == val && dfs(i) {
					return true
				}
			}
			// 不合适再出来
			res = res[:ln]
		}
		return false
	}

	dfs(0)

	return res
}

func main() {
	fmt.Println(splitIntoFibonacci("123456579"))
}
