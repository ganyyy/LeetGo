package main

import "container/heap"

//type StockPrice struct {
//	price      map[int]int // 时间戳->价格
//	latestTime int         // 最后一次出价的时间点
//	maxTime    int
//	minTime    int
//	maxValid   bool // 延时刷新用的
//	minValid   bool // 延时刷新用的
//}
//
//const DEFAULT = math.MinInt32
//const DEFAULT2 = math.MaxInt32
//
//func Constructor() StockPrice {
//	return StockPrice{
//		price: map[int]int{
//			DEFAULT:  DEFAULT,
//			DEFAULT2: DEFAULT2,
//		},
//		latestTime: DEFAULT,
//		maxTime:    DEFAULT,
//		minTime:    DEFAULT2,
//	}
//}
//
//func (sp *StockPrice) Update(timestamp int, price int) {
//
//	if price >= sp.price[sp.maxTime] {
//		sp.maxTime = timestamp
//	} else {
//		// 比之前的小, 是否合法取决于是不是当前缓存的最大时间的值变小了
//		sp.maxValid = sp.maxTime != timestamp
//	}
//
//	if price <= sp.price[sp.minTime] {
//		sp.minTime = timestamp
//	} else {
//		// 比之前的大, 是否合法取决于是不是当前缓存的最小时间的值变大了
//		sp.minValid = sp.minTime != timestamp
//	}
//
//	if timestamp >= sp.latestTime {
//		sp.latestTime = timestamp
//	}
//
//	sp.price[timestamp] = price
//}
//
//func (sp *StockPrice) Current() int {
//	return sp.price[sp.latestTime]
//}
//
//func (sp *StockPrice) Maximum() int {
//	if sp.maxValid {
//		return sp.price[sp.maxTime]
//	}
//	sp.maxValid = true
//	var maxV = DEFAULT
//	for t, v := range sp.price {
//		if v >= maxV {
//			sp.maxTime = t
//			maxV = v
//		}
//	}
//
//	return maxV
//}
//
//func (sp *StockPrice) Minimum() int {
//	if sp.minValid {
//		return sp.price[sp.minTime]
//	}
//	sp.minValid = true
//	var minV = DEFAULT2
//	for t, v := range sp.price {
//		if v >= minV {
//			sp.minTime = t
//			minV = v
//		}
//	}
//
//	return minV
//}

type StockPrice struct {
	maxPrice, minPrice hp
	timePriceMap       map[int]int
	maxTimestamp       int
}

func Constructor() StockPrice {
	return StockPrice{timePriceMap: map[int]int{}}
}

func (sp *StockPrice) Update(timestamp, price int) {
	// 通过符号实现大顶堆/小顶堆, 还能这么玩?
	heap.Push(&sp.maxPrice, pair{-price, timestamp})
	heap.Push(&sp.minPrice, pair{price, timestamp})
	sp.timePriceMap[timestamp] = price
	if timestamp > sp.maxTimestamp {
		sp.maxTimestamp = timestamp
	}
}

func (sp *StockPrice) Current() int {
	return sp.timePriceMap[sp.maxTimestamp]
}

func (sp *StockPrice) Maximum() int {
	for {
		// 不停的查找到当前最新的那个价格
		if p := sp.maxPrice[0]; -p.price == sp.timePriceMap[p.timestamp] {
			return -p.price
		}
		heap.Pop(&sp.maxPrice)
	}
}

func (sp *StockPrice) Minimum() int {
	for {
		// 不停的查找到当前最新的那个价格
		// 因为同一个时间戳可能会有多个价格, 需要通过迭代的形式获取正确的/符合要求的价格
		if p := sp.minPrice[0]; p.price == sp.timePriceMap[p.timestamp] {
			return p.price
		}
		heap.Pop(&sp.minPrice)
	}
}

type pair struct{ price, timestamp int }
type hp []pair

// 每次都要手写一个, 就很烦

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
