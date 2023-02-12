package main

import "sort"

func fillCups(amount []int) int {
	sort.Ints(amount)
	// 尽可能的将多的和少的均摊
	// 设排完序后的大小关系为 a, b, c
	// 1. 如果 c >= a + b, (1,2,3) -> (1+1) + (2+2) -> 3
	// 2. 如果 c <  a + b, (2,3,4) -> (2+2) + (2+2)+1 -> 5
	//                     (3,4,5) -> (1+1)+2 + (4+4) -> 7
	//                     (3,4,5) -> (3+3) + (2+2)+2 -> 7
	//                     (3,4,5) 先合并1,2  --> (2,3,5) -> (2+2) + (3+3) -> 5+1 = 6
	//    t = (a+b-z), t <= a <= b <= c
	//    如果t是偶数, a/b均分剩余的 t/2 -> c+t/2           -> (a+b+c)/2
	//    如果t是奇数, a/b均分剩余的 (t-1)/2 -> c+((t-1)/2)+1 -> (a+b+c+1)/2
	if amount[2] > amount[1]+amount[0] {
		return amount[2]
	}
	return (amount[0] + amount[1] + amount[2] + 1) / 2
}
