package main

import (
	"container/heap"
	"sort"
)

func busiestServers(k int, arrival, load []int) (ans []int) {
	available := hi{make([]int, k)} // 所有可以工作的服务器的集合, 以服务器id构成的小顶堆
	for i := 0; i < k; i++ {
		available.IntSlice[i] = i
	}
	var busy serverHeap // 当前正在工作的服务器集合, 以及对应的结束时间, 按照结束时间构成的小顶堆
	var requests = make([]int, k)
	var maxRequest = 0
	for i, t := range arrival {
		for len(busy) > 0 && busy[0].end <= t {
			heap.Push(&available, i+((busy[0].id-i)%k+k)%k) // 保证得到的是一个不小于 i 的且与 id 同余的数
			heap.Pop(&busy)
		}
		if available.Len() == 0 { // 当前没有可用的服务器, 直接放弃当前任务
			continue
		}
		id := heap.Pop(&available).(int) % k // 获取一个可用的服务器,
		requests[id]++                       // 增加统计计数
		if requests[id] > maxRequest {
			maxRequest = requests[id]
			ans = []int{id}
		} else if requests[id] == maxRequest {
			ans = append(ans, id)
		}
		heap.Push(&busy, serverPair{t + load[i], id}) // 更新当前工作中的服务器
	}
	return
}

type hi struct{ sort.IntSlice }

func (h *hi) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hi) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type serverPair struct{ end, id int }
type serverHeap []serverPair

func (h serverHeap) Len() int            { return len(h) }
func (h serverHeap) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h serverHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *serverHeap) Push(v interface{}) { *h = append(*h, v.(serverPair)) }
func (h *serverHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
