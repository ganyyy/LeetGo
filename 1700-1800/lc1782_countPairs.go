package main

import (
	"fmt"
	"sort"
)

func countPairs1782(n int, edges [][]int, queries []int) []int {
	degree := make([]int, n)
	cnt := map[int]int{}
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		if x > y {
			x, y = y, x
		}
		degree[x]++
		degree[y]++
		// 类似于转换成唯一id. 除了去重外, 也是可以查询到相连边的一种方式
		// 按照题目要求, 相关边 = degree[x]+degree[y]-cnt[x*n+y]
		cnt[x*n+y]++
	}

	arr := make([]int, n)
	copy(arr, degree)
	// 按照度数排序
	sort.Ints(arr)
	var ans = make([]int, 0, len(queries))
	for _, bound := range queries {
		total := 0
		// 这个算法啊就很有意思
		// 现在找到相当于是 arr[i] + XXX > bound
		// XXX的范围可以通过双指针进行缩小
		// 假设 找到了位置j, 使得 arr[i]+arr[j] <= bound,
		// 那么从 [j, n-1]范围内的所有数字都是满足需求的
		// 为什么要取max(i,j)呢?
		// 可以这样理解:
		//  1. i增加就意味着j需要减少, 当j == i时, 此时[i:]部分的和都是 > bound的, 使用i是为了避免重复计算
		//  2. 假设j==n-1时, 都不满足需求, 此时针对i而言, totel 应该是0
		for i, j := 0, n-1; i < n; i++ {
			for j > i && arr[i]+arr[j] > bound {
				j--
			}
			total += n - 1 - max(i, j)
		}
		fmt.Println(total)
		for val, freq := range cnt {
			x, y := val/n, val%n
			if degree[x]+degree[y] > bound && degree[x]+degree[y]-freq <= bound {
				total--
			}
		}
		ans = append(ans, total)
	}
	return ans
}
