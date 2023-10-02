package main

import (
	"container/list"
	"fmt"
	"sort"
	"strings"
)

// 远比想象的要复杂, 很明显,
// 感觉这个LFU算法有问题啊. 新加入的元素被淘汰的概率很大

type LFUNode struct {
	Key       int
	Val       int
	Frequency int
}

func (l LFUNode) String() string {
	return fmt.Sprintf("[K:%v, V:%v, F:%v]", l.Key, l.Val, l.Frequency)
}

type LFUCache struct {
	capacity       int
	minFrequency   int
	frequencyCache map[int]*list.List    // 所有频率对应的列表
	cache          map[int]*list.Element // 所有节点的数据
}

func Constructor460(capacity int) LFUCache {
	return LFUCache{
		capacity:       capacity,
		minFrequency:   1,
		frequencyCache: map[int]*list.List{},
		cache:          make(map[int]*list.Element, capacity),
	}
}

func (lfu *LFUCache) Get(key int) int {
	if ele, ok := lfu.cache[key]; ok {
		var lfuNode = ele.Value.(*LFUNode)
		lfu.updateNode(lfuNode, ele)
		// lfu.Show(fmt.Sprintf("Get:%v", key))
		return lfuNode.Val
	} else {
		return -1
	}
}

func (lfu *LFUCache) getFrequencyList(frequency int) *list.List {
	var frequencyList, ok = lfu.frequencyCache[frequency]
	if !ok {
		frequencyList = list.New()
		lfu.frequencyCache[frequency] = frequencyList
	}
	return frequencyList
}

func (lfu *LFUCache) updateNode(lfuNode *LFUNode, ele *list.Element) {
	// 从旧的队列中删除
	var oldList = lfu.getFrequencyList(lfuNode.Frequency)
	oldList.Remove(ele)
	if oldList.Len() == 0 {
		delete(lfu.frequencyCache, lfuNode.Frequency)
		// 更新最低的频率
		if lfuNode.Frequency == lfu.minFrequency {
			lfu.minFrequency++
		}
	}
	// 增加节点的频率
	lfuNode.Frequency++
	// 插入到新的列表中
	lfu.cache[lfuNode.Key] = lfu.getFrequencyList(lfuNode.Frequency).PushFront(lfuNode)
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity <= 0 {
		return
	}
	if ele, ok := lfu.cache[key]; ok {
		var lfuNode = ele.Value.(*LFUNode)
		lfuNode.Val = value
		lfu.updateNode(lfuNode, ele)
		// lfu.Show(fmt.Sprintf("Put:%v,%v", key, value))
	} else {
		if len(lfu.cache) >= lfu.capacity {
			// 移除最小频率列表对应的最后一个值
			var minList = lfu.getFrequencyList(lfu.minFrequency)
			var lfuNode = minList.Remove(minList.Back()).(*LFUNode)
			delete(lfu.cache, lfuNode.Key)
			if minList.Len() == 0 {
				delete(lfu.frequencyCache, lfu.minFrequency)
				lfu.minFrequency++
			}
		}
		var lfuNode = &LFUNode{
			Key:       key,
			Val:       value,
			Frequency: 1,
		}
		lfu.minFrequency = 1
		lfu.cache[key] = lfu.getFrequencyList(lfu.minFrequency).PushFront(lfuNode)
		// lfu.Show(fmt.Sprintf("Put Remove:%v,%v", key, value))
	}
}

func (lfu *LFUCache) Show(reason string) {
	type Show struct {
		F    int
		list *list.List
	}
	var tmp []Show
	for f, l := range lfu.frequencyCache {
		tmp = append(tmp, Show{
			F:    f,
			list: l,
		})
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].F < tmp[j].F
	})

	var sb strings.Builder
	sb.WriteString(reason)
	sb.WriteString(fmt.Sprintf("\nminFrequency:%v\n", lfu.minFrequency))
	for _, show := range tmp {
		sb.WriteString(fmt.Sprintf("%v:{", show.F))
		for ele := show.list.Front(); ele != nil; ele = ele.Next() {
			sb.WriteString(fmt.Sprintf("%s,", ele.Value.(LFUNode).String()))
		}
		sb.WriteString("}\n")
	}
	fmt.Println(sb.String())
}
