package main

import "math"

func findShortestSubArray(nums []int) int {
	var m = make(map[int]int)
	// 第一步: 统计每个数字的个数
	var cnt int
	for _, v := range nums {
		m[v]++
		cnt = max(cnt, m[v])
	}

	// 第二步: 保留次数最多的那个值.
	for k, v := range m {
		if v == cnt {
			continue
		}
		delete(m, k)
	}
	// 时间复杂度 O(n)
	// 空间复杂度 O(n)

	// 标记起始位置
	var res = math.MaxInt32
	var sm = make(map[int]int, len(m))
	// 第三步: 找出相应的最小子数组
	for i, num := range nums {
		var v, ok = m[num]
		if !ok {
			continue
		}
		start, ok := sm[num]
		if !ok {
			sm[num] = i
		}

		v--
		if v > 0 {
			m[num] = v
			continue
		}
		delete(m, num)
		res = min(res, i-start+1)
	}

	return res
}

// [0]代表 起始位置
// [1]代表 最后发现的位置
// [2]代表 出现的次数
type numCnt [3]int

const (
	StartPos = iota
	EndPos
	Count
)

func findShortestSubArray2(nums []int) int {
	var m = make(map[int]numCnt, 100)
	// 第一步: 统计每个数字的个数
	var nm int
	for i, v := range nums {
		var cnt, ok = m[v]
		if ok {
			cnt[EndPos] = i
			cnt[Count]++
		} else {
			cnt[StartPos] = i
			cnt[EndPos] = i
			cnt[Count] = 1
		}
		nm = max(nm, cnt[2])
		m[v] = cnt
	}

	var res = math.MaxInt32

	// 第一步: 计算最大值
	for _, v := range nums {
		var cnt, ok = m[v]
		if !ok {
			continue
		}
		if cnt[Count] == nm {
			res = min(res, cnt[EndPos]-cnt[StartPos]+1)
		}
		delete(m, v)
		// 提前退出
		if len(m) == 0 {
			break
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
