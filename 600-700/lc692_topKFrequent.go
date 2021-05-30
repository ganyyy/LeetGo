package main

import (
	"fmt"
	"sort"
)

type node struct {
	v   string
	cnt int
}

type nodes []node

func (n nodes) Len() int {
	return len(n)
}

func (n nodes) Less(i, j int) bool {
	var a, b = n[i], n[j]
	if a.cnt != b.cnt {
		return a.cnt > b.cnt
	}
	return a.v < b.v
}

func (n nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func topKFrequent(words []string, k int) []string {
	var m = make(map[string]int, len(words))

	for _, v := range words {
		m[v]++
	}

	if len(m) <= k {
		var ret = make([]string, len(m))
		var n nodes = make([]node, 0, len(m))
		for key, val := range m {
			n = append(n, node{v: key, cnt: val})
		}
		sort.Sort(n)
		for i := range n {
			ret[i] = n[i].v
		}
		return ret
	}

	var n nodes = make([]node, 0, len(m))
	for key, val := range m {
		n = append(n, node{key, val})
	}

	//var n nodes = []node{
	//	{"leetcode", 1}, {"coding", 1}, {"i", 2}, {"love", 2},
	//}
	// TopK快排

	var quickSort func(st, ed int)

	quickSort = func(st, ed int) {
		fmt.Println(n)
		if st >= ed {
			return
		}
		var l, r = st, ed
		for l < r {
			for l < r && n.Less(st, r) {
				r--
			}
			for l < r && !n.Less(st, l) {
				l++
			}
			if l < r {
				n.Swap(l, r)
			}
		}
		n.Swap(l, st)
		quickSort(st, l-1)
		if l < k {
			quickSort(l+1, ed)
		}
	}

	quickSort(0, len(n)-1)

	var ret = make([]string, len(n))
	for i := 0; i < k; i++ {
		ret[i] = n[i].v
	}
	return ret
}

func quickSort(nums []int) {
	var sub func(st, ed int)

	sub = func(st, ed int) {
		if st >= ed {
			return
		}
		var l, r = st, ed
		for l < r {
			for l < r && nums[r] > nums[st] {
				r--
			}
			for l < r && nums[l] < nums[st] {
				l++
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
		}
		//if l < r {
		nums[l], nums[st] = nums[st], nums[l]
		//}
		sub(st, l-1)
		sub(l+1, ed)
	}
	sub(0, len(nums)-1)
}

func quickSort2(nums []int) {
	var sub func(i, j int)

	sub = func(st, ed int) {
		if st >= ed {
			return
		}
		var index = st + 1
		for i := index; i <= ed; i++ {
			if nums[i] <= nums[st] {
				// index 左边的小于当前值, index右边的大于当前值
				nums[index], nums[i] = nums[i], nums[index]
				index++
			}
		}
		nums[st], nums[index-1] = nums[index-1], nums[st]
		sub(st, index-2)
		sub(index, ed)
	}

	sub(0, len(nums)-1)
}

func main() {
	//fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))

	//var nums = make([]int, 100)
	//for i := range nums {
	//	nums[i] = i
	//}
	//rand.Shuffle(len(nums), func(i, j int) {
	//	nums[i], nums[j] = nums[j], nums[i]
	//})
	//quickSort(nums)
	//
	//fmt.Println(nums)
	//
	//nums = []int{1, 2}
	//quickSort(nums)
	//
	//fmt.Println(nums)

	//sort.Sort()
	var nums = []int{7, 3, 2, 1}
	quickSort2(nums)
	fmt.Println(nums)
}
