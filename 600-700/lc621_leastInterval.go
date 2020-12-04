package main

import "sort"

func leastInterval(tasks []byte, n int) int {

	if n == 0 {
		return len(tasks)
	}

	var cnt = make([]int, 26)

	var max int
	for _, t := range tasks {
		cnt[int(t)-'A']++
	}

	// 倒序排一下
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})

	// max就是最大的哪个
	max = cnt[0]

	// 假设A为出现频率最高的任务
	// 那么最长的情况是
	// A...A...A...
	// 最短的情况是
	// A...A...A...A
	// 需要根据和A频率相同的任务的数量来决定后边要补几个任务
	var i int
	for i = 1; i < 26; i++ {
		if cnt[i] != max {
			break
		}
	}
	// 假设出现了正好可以分配的一组活动
	// [A,A,B,B,C,D] n = 1
	// 按照原来的计算方式可得
	// (2-1)*(1+1)+2 = 4 < len(tasks)
	// 可以通过适当的内嵌找到一个合理的方法进行分配, 此时恰好需要 len(tasks) 个操作时间
	if t := (max-1)*(n+1) + i; t > len(tasks) {
		return t
	} else {
		return len(tasks)
	}

}
