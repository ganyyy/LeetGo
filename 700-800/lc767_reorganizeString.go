package _00_800

import (
	"sort"
	"unsafe"
)

func reorganizeString(S string) string {
	if len(S) <= 1 {
		return S
	}

	// 统计每个字符的数量
	// 0 表示字符, 1 表示数量
	var cnt = make([][2]int, 26)
	for i := 0; i < 26; i++ {
		cnt[i][0] = i
	}

	var max = 0
	// 一共存在多少字符
	// 如果数量最多的字符超过了一半? 就是不可行的
	for i := range S {
		var idx = S[i] - 'a'
		cnt[idx][1]++
		// 统计量最大的字符
		if cnt[idx][1] > max {
			max = cnt[idx][1]
		}
	}
	// 数量最多的字符超过了一半, 没法实现目标
	if max > (len(S)+1)/2 {
		return ""
	}

	// 组合成一个新的结果
	var ret = make([]byte, len(S))

	// 还是需要排序, 就离谱
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i][1] > cnt[j][1]
	})

	//隔着填, 越界了就从1开始
	var pre = 0
	for _, v := range cnt {
		if v[1] == 0 {
			continue
		}
		for j := 0; j < v[1]; j++ {
			if pre >= len(S) {
				pre = 1
			}
			ret[pre] = byte(v[0]) + 'a'
			pre += 2
		}
	}
	return *(*string)(unsafe.Pointer(&ret))
}
