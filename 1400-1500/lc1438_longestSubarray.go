//go:build ignore

package main

import (
	"container/heap"
)

type Queue []int

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x.(int))
}

func (q *Queue) Pop() (ret interface{}) {
	ret, *q = (*q)[q.Len()-1], (*q)[:q.Len()-1]
	return
}

type MinQueue struct {
	Queue
}

func (m MinQueue) Less(i, j int) bool {
	return m.Queue[i] < m.Queue[j]
}

type MaxQueue struct {
	Queue
}

func (m MaxQueue) Less(i, j int) bool {
	return m.Queue[i] > m.Queue[j]
}

func longestSubarray(nums []int, limit int) int {
	if len(nums) == 0 {
		return 0
	}

	var l, r, res int
	var maxQueue = MaxQueue{Queue{nums[0]}}
	var minQueue = MinQueue{Queue{nums[0]}}

	for r = 1; r < len(nums); r++ {
		heap.Push(&maxQueue, nums[r])
		heap.Push(&minQueue, nums[r])
		if abs(maxQueue.Queue[0]-minQueue.Queue[0]) <= limit {
			// 当前也算
			res = max(res, r-l+1)
			continue
		}
		// 删除一次就好了? 反正下次也会继续删除
		// 从堆里找到最左边的那个值

		// golang这个就离谱
		for i, v := range maxQueue.Queue {
			if v == nums[l] {
				heap.Remove(&maxQueue, i)
				break
			}
		}
		for i, v := range minQueue.Queue {
			if v == nums[l] {
				heap.Remove(&minQueue, i)
				break
			}
		}
		l++
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func longestSubarray2(nums []int, limit int) int {

	// 保存的是按照值顺序排行对应的索引
	minQ := make([]int, 0, len(nums)>>1)
	maxQ := make([]int, 0, len(nums)>>1)
	res := 0
	l := 0
	r := 0

	// 我算是服了, 手动维护的队列性能好于官方的堆
	minQ = append(minQ, 0)
	maxQ = append(maxQ, 0)
	for r < len(nums) {
		maxId := maxQ[0]
		minId := minQ[0]
		if nums[maxId]-nums[minId] <= limit {
			if r-l+1 > res {
				res = r - l + 1
			}
			r++
			for r < len(nums) && len(minQ) > 0 && nums[minQ[len(minQ)-1]] > nums[r] {
				minQ = minQ[:len(minQ)-1]
			}
			minQ = append(minQ, r)
			for r < len(nums) && len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] < nums[r] {
				maxQ = maxQ[:len(maxQ)-1]
			}
			maxQ = append(maxQ, r)
		} else {
			l++
			for maxQ[0] < l {
				maxQ = maxQ[1:]
			}
			for minQ[0] < l {
				minQ = minQ[1:]
			}
		}
	}
	return res
}

func longestSubarray3(nums []int, limit int) int {
	var ln = len(nums)
	if ln <= 1 {
		return ln
	}

	// 这个性能差了好多
	// 应该是copy的问题

	var maxQ = make([]int, 0, ln>>1)
	var minQ = make([]int, 0, ln>>1)

	var l, r, res int
	var mi, ma int
	maxQ = append(maxQ, 0)
	minQ = append(minQ, 0)
	for r < ln {
		mi, ma = minQ[0], maxQ[0]

		// 先看一下是否符合标准
		if nums[ma]-nums[mi] <= limit {
			res = max(res, r-l+1)

			r++
			if r >= ln {
				break
			}

			// 更新一下 r 的位置信息
			var idx = len(maxQ) - 1
			for idx >= 0 && nums[maxQ[idx]] < nums[r] {
				idx--
			}
			maxQ = maxQ[:idx+1]
			maxQ = append(maxQ, r)

			idx = len(minQ) - 1
			for idx >= 0 && nums[minQ[idx]] > nums[r] {
				idx--
			}
			minQ = minQ[:idx+1]
			minQ = append(minQ, r)

		} else {
			var idx int
			for maxQ[idx] <= l {
				idx++
			}
			copy(maxQ, maxQ[idx:])
			maxQ = maxQ[:len(maxQ)-idx]

			idx = 0
			for minQ[idx] <= l {
				idx++
			}
			copy(minQ, minQ[idx:])
			minQ = minQ[:len(minQ)-idx]

			l++
		}
	}

	return res
}

func main() {
	/*
		[8,2,4,7]
		4
	*/
	println(longestSubarray3([]int{8, 2, 4, 7}, 4))
}
