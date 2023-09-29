package main

import (
	"sort"
	"unsafe"
)

func earliestFullBloom(plantTime []int, growTime []int) int {
	if len(plantTime) == 0 {
		return 0
	}
	// 高地址存储id, 低地址保持原始值
	for id := range growTime {
		sp := (*[2]int32)(unsafe.Pointer(&growTime[id]))
		sp[1] = int32(id)
	}

	sort.Slice(growTime, func(i, j int) bool {
		return int32(growTime[i]) > int32(growTime[j])
	})

	// 按照生长时间来排逆序, 不管播种时间几何, 核心就是让生长慢的尽可能的先播种
	// 播种无需叉开.
	var plant, finish int
	for _, id := range growTime {
		plant += plantTime[id>>32]
		finish = max(finish, plant+int(int32(id)))
	}
	return finish
}
