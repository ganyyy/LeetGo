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
	var idx = sort.Search(len(vals), func(i int) bool {
		return vals[i].T > timestamp
	})
	if idx > 0 {
		return vals[idx-1].V
	} else {
		return ""
	}
}
