package main

import "container/heap"

// CountN 前32位是count, 后32位是N
type CountN uint64

// GetCount 获取count
func (cn *CountN) GetCount() int { return int(*cn >> 32) }

// GetN 获取N
func (cn *CountN) GetN() int { return int(*cn & 0xffffffff) }

// SetCount 设置count
func (cn *CountN) SetCount(count int) {
	*cn = CountN((uint64(count) << 32) | (uint64(*cn) & 0xffffffff))
}

// SetN 设置N
func (cn *CountN) SetN(n int) { *cn = CountN((uint64(*cn) & 0xffffffff00000000) | uint64(n)) }

// Less 比较大小
func (cn *CountN) Less(other CountN) bool {
	if cn.GetCount() == other.GetCount() {
		return cn.GetN() < other.GetN()
	}
	return cn.GetCount() > other.GetCount()
}

// CountNHeap 堆
type CountNHeap []CountN

func (h CountNHeap) Len() int           { return len(h) }
func (h CountNHeap) Less(i, j int) bool { return h[i].Less(h[j]) }
func (h CountNHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 添加元素
func (h *CountNHeap) Push(x interface{}) {
	*h = append(*h, x.(CountN))
}

// Pop 弹出元素
func (h *CountNHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func rearrangeBarcodes(barcodes []int) []int {
	// 统计每个数字出现的次数
	counts := make(map[int]int)
	for _, barcode := range barcodes {
		counts[barcode]++
	}
	// 用堆来存储每个数字出现的次数
	h := &CountNHeap{}
	for barcode, count := range counts {
		heap.Push(h, CountN((uint64(count)<<32)|uint64(barcode)))
	}
	// 从堆中取出数字, 每次取两个, 保证不会有相邻的数字
	res := make([]int, len(barcodes))
	for i := 0; i < len(res); i += 2 {
		cn1 := heap.Pop(h).(CountN)
		res[i] = cn1.GetN()
		if i+1 < len(res) {
			cn2 := heap.Pop(h).(CountN)
			res[i+1] = cn2.GetN()
			cn2.SetCount(cn2.GetCount() - 1)
			if cn2.GetCount() > 0 {
				heap.Push(h, cn2)
			}
		}
		cn1.SetCount(cn1.GetCount() - 1)
		if cn1.GetCount() > 0 {
			heap.Push(h, cn1)
		}

	}
	return res
}

func rearrangeBarcodes2(barcodes []int) []int {
	if len(barcodes) < 2 {
		return barcodes
	}

	counts := make(map[int]int)
	for _, b := range barcodes {
		counts[b] = counts[b] + 1
	}

	// 偶数idx
	evenIndex := 0
	// 奇数idx
	oddIndex := 1
	// 如果数组中, 某个元素超过了长度的一半:
	//  数组长度为奇数
	//  有且只有一个元素可以超过一半
	halfLength := len(barcodes) / 2
	res := make([]int, len(barcodes))
	for x, count := range counts {
		// 针对 小于数组一半的数字, 优先从奇数位开始放置, 直到奇数下标超过了数组的长度
		for count > 0 && count <= halfLength && oddIndex < len(barcodes) {
			res[oddIndex] = x
			count--
			oddIndex += 2
		}
		// 如果某个元素数量超过了数组的一半, 或者奇数下标跑慢了数组, 就从偶数位开始递增
		for count > 0 {
			res[evenIndex] = x
			count--
			evenIndex += 2
		}
	}
	return res
}

func main() {
	barcodes := []int{1, 1, 1, 2, 2, 2}
	rearrangeBarcodes(barcodes)
}
