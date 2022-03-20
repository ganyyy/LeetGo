package main

import (
	"container/list"
	"math"
)

type StrObj struct {
	V   string
	Cnt int
}

// 当成LFU来搞了

type AllOne struct {
	keys           map[string]*list.Element
	all            map[int]*list.List
	minIdx, maxIdx int
}

func Constructor() AllOne {
	return AllOne{
		keys:   map[string]*list.Element{},
		all:    map[int]*list.List{},
		maxIdx: math.MinInt32,
		minIdx: math.MaxInt32,
	}
}

func (ao *AllOne) getList(idx int) *list.List {
	if idx < ao.minIdx {
		ao.minIdx = idx
	}
	if idx > ao.maxIdx {
		ao.maxIdx = idx
	}
	var allList = ao.all[idx]
	if allList == nil {
		allList = list.New()
		ao.all[idx] = allList
	}
	return allList
}

func (ao *AllOne) updateStrObj(ele *list.Element, add bool) {
	var obj = ele.Value.(StrObj)
	var oldLevel = obj.Cnt
	if add {
		obj.Cnt++
	} else {
		obj.Cnt--
	}

	if oldLevel != 0 {
		var preList = ao.getList(oldLevel)
		preList.Remove(ele)
		if preList.Len() == 0 {
			if ao.minIdx == oldLevel {
				ao.minIdx++
			}
			if ao.maxIdx == oldLevel {
				ao.maxIdx--
			}
		}
	}
	if obj.Cnt > 0 {
		var newList = ao.getList(obj.Cnt)
		ao.keys[obj.V] = newList.PushFront(obj)
	} else {
		delete(ao.keys, obj.V)
	}
}

func (ao *AllOne) Inc(key string) {
	if ele, ok := ao.keys[key]; ok {
		ao.updateStrObj(ele, true)
	} else {
		ao.updateStrObj(&list.Element{Value: StrObj{V: key}}, true)
	}
}

func (ao *AllOne) Dec(key string) {
	if ele, ok := ao.keys[key]; ok {
		ao.updateStrObj(ele, false)
	}
}

func (ao *AllOne) GetMaxKey() string {
	if len(ao.keys) == 0 {
		return ""
	}
	for lst := ao.getList(ao.maxIdx); ao.maxIdx >= ao.minIdx; {
		if lst.Len() == 0 {
			ao.maxIdx--
			lst = ao.getList(ao.maxIdx)
			continue
		}
		return lst.Front().Value.(StrObj).V
	}
	return ""
}

func (ao *AllOne) GetMinKey() string {
	if len(ao.keys) == 0 {
		return ""
	}
	for lst := ao.getList(ao.minIdx); ao.maxIdx >= ao.minIdx; {
		if lst.Len() == 0 {
			ao.minIdx++
			lst = ao.getList(ao.minIdx)
			continue
		}
		return lst.Front().Value.(StrObj).V
	}
	return ""
}
