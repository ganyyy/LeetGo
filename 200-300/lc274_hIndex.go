package main

import "sort"

func hIndex2(citations []int) int {
	sort.Ints(citations)

	// 什么狗屁计数
	// 核心就是: 数组中存在一个值, 这个值不会超过数组的数量.
	// 这个值可以将数组分为两部分
	// 前半部分的值都比这个数组小, 后半部分值都大于等于这个值
	// 后半部分的长度小于等于这个值, 结果要求这个值的最大可能

	// 引用的次数从小到大
	// 长度-当前索引表示剩余的论文数量, 引用次数大于等于当前值的论文的数量
	for i, v := range citations {
		// 所以, 首个出现的引用次数大于等于论文数量的值就是数量最多的?
		if v >= len(citations)-i {
			return len(citations) - i
		}
	}

	return 0
}

func hIndex(citations []int) (h int) {
	n := len(citations)
	counter := make([]int, n+1)
	for _, citation := range citations {
		// H计数肯定不可能超过总的论文数量
		if citation >= n {
			counter[n]++
		} else {
			counter[citation]++
		}
	}
	// 倒序计算, tot是累计的论文数量, i 是当前的引用计数
	// 如果 tot >= i, 说明当前满足条件
	for i, tot := n, 0; i >= 0; i-- {
		tot += counter[i]
		if tot >= i {
			return i
		}
	}
	return 0
}
