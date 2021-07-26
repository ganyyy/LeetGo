package main

import "sort"

func minOperations(target, arr []int) int {
	n := len(target)
	// 记录每个数字出现的位置
	pos := make(map[int]int, n)
	for i, val := range target {
		pos[val] = i
	}
	// 这里保存的是arr中存在于 target原始书顺序的数字, 即不添加的情况下, 最长的公共子数组部分有多长
	var d []int
	for _, val := range arr {
		if idx, has := pos[val]; has {
			// sort.SearchInts这个接口, 返回的是元素所在的位置或者要插入的位置
			// 通过这种方式可以保证留在 数组 d 中的元素是有序的
			if p := sort.SearchInts(d, idx); p < len(d) {
				// 如果搜索到的位置小于当前数组的长度
				// 说明要么存在这个idx, 要么这个idx应该放在返回的位置p上
				// dp[p]的值 >= idx. 这样做了替换之后, 整体有序, 且可以保证数组中的元素整体从大到小存在
				d[p] = idx
			} else {
				// 如果超过了数组大小, 说明需要后续补一个
				d = append(d, idx)
			}
		}
	}
	// 去掉正常的, 就是需要填补的
	return n - len(d)
}

func main() {
	minOperations([]int{6, 4, 8, 1, 3, 2}, []int{4, 7, 6, 2, 3, 8, 6, 1})
}
