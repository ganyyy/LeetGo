package main

import (
	"container/heap"
	"sort"
)

type PriorityQueue [][2]int64

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	(*pq) = append(*pq, x.([2]int64))
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	x := (*pq)[n-1]
	(*pq) = (*pq)[:n-1]
	return x
}

func kSum(nums []int, k int) int64 {
	n, total := len(nums), int64(0)
	for i := range nums {
		if nums[i] >= 0 {
			total += int64(nums[i])
		} else {
			nums[i] = -nums[i]
		}
	}
	sort.Ints(nums)

	// 这个转换思路很有意思..
	// total是所有的正数和, 想要求出数组的第k个最大和, 等同于 total - 第k个最小和
	// 换个理解方式: 经过改造后的nums中存储的都是非负数,
	// 此时, 任意子序列的和都可以表示为:
	//      total - 不在子序列中的正数 + 不在子序列中的负数
	//      = total - (不在子序列中的正数+不在子序列中的负数的绝对值)
	// 通过堆, 获取到了最小的子序列和tk, 那么使用total - tk 就是第k大的子序列和
	//      最大和最小子序列二者是互补的!
	ret := int64(0)
	pq := PriorityQueue{
		[2]int64{int64(nums[0]), 0},
	}
	// 每次heap.Pop都会弹出当前堆顶的最小值, 也就是说: 我们通过弹出K次获取到第k个最小值
	for j := 2; j <= k; j++ {
		t, i := pq[0][0], pq[0][1]
		heap.Pop(&pq)
		ret = t
		if i == int64(n-1) {
			continue
		}
		// 添加nums[i+1]
		heap.Push(&pq, [2]int64{t + int64(nums[i+1]), i + 1})
		// 替换nums[i]为nums[i+1]
		heap.Push(&pq, [2]int64{t - int64(nums[i]-nums[i+1]), i + 1})
	}
	return total - ret
}
