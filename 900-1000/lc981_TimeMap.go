package main

import "sort"

type Val struct {
	T int
	V string
}

type SortVal []Val

type TimeMap struct {
	m map[string]SortVal
}

func Constructor() TimeMap {
	return TimeMap{m: map[string]SortVal{}}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	t.m[key] = append(t.m[key], Val{T: timestamp, V: value})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	var vals = t.m[key]
	if len(vals) == 0 {
		return ""
	}
	// 这个search接口, 如果函数返回true表示答案在左边; 否则在右边
	// 如果找不到相等的, 返回的就是应该在的位置(距离最近的 函数返回值为true的索引)
	var idx = sort.Search(len(vals), func(i int) bool {
		return vals[i].T > timestamp
	})
	if idx > 0 {
		return vals[idx-1].V
	} else {
		return ""
	}
}
