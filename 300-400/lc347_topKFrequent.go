package main

import (
	"container/heap"
	"fmt"
)

type topK [][2]int

func (t *topK) Len() int {
	return len(*t)
}

func (t *topK) Less(i, j int) bool {
	return (*t)[i][1] < (*t)[j][1]
}

func (t *topK) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *topK) Push(x interface{}) {
	*t = append(*t, x.([2]int))
}

func (t *topK) Pop() interface{} {
	var x interface{}
	x, *t = (*t)[len(*t)-1], (*t)[:len(*t)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) <= k {
		return nums
	}
	// 堆
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	res := make([]int, 0, k)
	// 如果总数小的话, 直接返回
	if len(m) <= k {
		for i := range m {
			res = append(res, i)
		}
		return res
	}

	var t topK = make([][2]int, 0, k)

	var cnt int
	for i, v := range m {
		if cnt >= k {
			break
		}
		cnt++
		heap.Push(&t, [2]int{i, v})
		delete(m, i)
	}

	// 走一个小顶堆
	for i, v := range m {
		top := heap.Pop(&t).([2]int)
		if top[1] < v {
			top[0] = i
			top[1] = v
		}
		heap.Push(&t, top)
	}

	for _, v := range t {
		res = append(res, v[0])
	}
	return res
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	}
}
