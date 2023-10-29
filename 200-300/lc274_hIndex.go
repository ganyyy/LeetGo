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
	// v: 引用次数. 因为是从小到大, 所以往后的v一定是大于等于当前的v
	// len(citations)-i: 剩余的论文数量
	// v >= len(citations)-i: 剩余的论文的引用次数都是大于等于当前的v

	// 好他妈绕啊. 存在一个值, 假设这个值为h, 那么需要有超过h篇的论文的引用次数大于等于h
	// 首先可以肯定的是: h不一定是数组中的值, 因为如果数组中的值都大于n, 就无解了
	// 所以这也解释了为啥返回的是 len(citations)-i, 而不是v
	// 从最终结果上来考量, h指数一定是论文的数量, 因为h指数是指有h篇论文的引用次数大于等于h

	total := len(citations)
	for i, v := range citations {
		// 所以, 首个出现的引用次数大于等于论文数量的值就是数量最多的?
		if total-i <= v {
			return len(citations) - i
		}
	}

	return 0
}

func hIndex5(citations []int) int {
	sort.Ints(citations)

	total := len(citations)
	for n, v := range citations {
		// remain 代表 发表的引用次数 >= v 的论文数,
		// 这些论文数至少被引用了v次
		// remain 和 v 不是对等关系.
		remain := total - n
		if v >= remain {
			// 卧槽, 一眼贪心啊
			// 第一个一定是remain最大的!
			return remain
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

func hIndex4(citations []int) (ret int) {
	n := len(citations)
	counter := make([]int, n+1)
	for _, citation := range citations {
		if citation >= n {
			counter[n]++
		} else {
			counter[citation]++
		}
	}
	// h指数是论文的数量!
	// h: h指数, 最大也就是n. 因为总共发表了n篇论文
	// count: 截止到h指数时, 总的论文数量
	for h, count := n, 0; h >= 0; h-- {
		// 有count篇论文的引用次数大于等于h
		// 根据贪心的思想, 首次出现的i就是最大的h指数
		count += counter[h]
		if count >= h {
			return h
		}
	}
	return 0
}
