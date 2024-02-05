package main

import (
	"container/heap"
	"sort"
)

type PriorityQueue struct {
	sort.IntSlice
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.IntSlice[i] < pq.IntSlice[j]
}

func (pq *PriorityQueue) Push(v interface{}) {
	pq.IntSlice = append(pq.IntSlice, v.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	arr := pq.IntSlice
	v := arr[len(arr)-1]
	pq.IntSlice = arr[:len(arr)-1]
	return v
}

func magicTower(nums []int) int {
	q := &PriorityQueue{}
	ans, hp, delay := 0, int64(1), int64(0)
	for _, num := range nums {
		if num < 0 {
			// 如果是负数, 加入到小根堆
			heap.Push(q, num)
		}
		hp += int64(num)
		if hp <= 0 {
			// 这就意味着需要一次调整.
			// 将堆顶元素移动到末尾
			ans++
			// 所有被移动到后边的元素的总和
			delay += int64(q.IntSlice[0])
			// 当前hp额外增加的血量
			hp -= int64(heap.Pop(q).(int))
		}
	}
	if hp+delay <= 0 {
		// 当前血量补不了延迟计算的血量
		return -1
	}
	return ans
}
