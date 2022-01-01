package main

import "sort"

func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	var ln = len(ages)
	var ret int
	for i := 0; i < ln; i++ {
		var a = ages[i]
		var endIdx = sort.Search(ln, func(idx int) bool {
			return ages[idx] > ages[i]
		})
		var start = a*5 + 70
		var startIdx = sort.Search(ln, func(idx int) bool {
			return ages[idx]*10 > start
		})
		if endIdx < startIdx {
			continue
		}
		if i < endIdx && i >= startIdx {
			ret--
		}
		ret += endIdx - startIdx
	}
	return ret
}

func numFriendRequestsGood(ages []int) int {
	group := make([]int, 121)
	for i := 0; i < len(ages); i++ {
		group[ages[i]]++
	}

	// 桶排序, 因为区间固定且较小. 完全可以统计个区间的人数
	res := 0
	for i := 1; i < 121; i++ {
		if group[i] == 0 {
			continue
		}
		for k := 1; k < 121; k++ {
			if i < k || (i < 100 && k > 100) || (k <= i/2+7) {
				continue
			}
			tmp := group[i] * group[k]
			// 去掉自己
			// 每个自己都要去掉自己, 所以直接减去全部即可
			if i == k {
				tmp -= group[i]
			}
			res += tmp
		}
	}
	return res
}

func main() {
	println(numFriendRequests([]int{20, 30, 100, 110, 120}))
}