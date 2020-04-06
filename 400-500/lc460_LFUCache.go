package main

import "time"

// 远比想象的要复杂, 很明显,

type LFUCache struct {
	queen    []*Node
	m        map[int]*Node
	capacity int
}

type Node struct {
	value, count int
	t            int64
}

func Constructor(capacity int) LFUCache {
	// 初始化一个map
	return LFUCache{
		m:        map[int]*Node{},
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	// count++
	if node, ok := this.m[key]; ok {
		node.count++
		return node.value
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	// 判断是否越界
	if len(this.m) >= this.capacity {
		var mKey int
		var mCount = 1<<31 - 1
		var mTime int64 = 1<<63 - 1
		for key, v := range this.m {
			if v.count <= mCount && v.t <= mTime {
				mKey = key
			}
		}
		delete(this.m, mKey)
	}
	this.m[key] = &Node{
		value: value,
		count: 0,
		t:     time.Now().UnixNano(),
	}
}

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)    // 返回 1
	cache.Put(3, 3) // 去除 key 2
	cache.Get(2)    // 返回 -1 (未找到key 2)
	cache.Get(3)    // 返回 3
	cache.Put(4, 4) // 去除 key 1
	cache.Get(1)    // 返回 -1 (未找到 key 1)
	cache.Get(3)    // 返回 3
	cache.Get(4)    // 返回 4

}
