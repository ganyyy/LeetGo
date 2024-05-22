package main

func longestEqualSubarray(nums []int, k int) int {

	// 记录每个数出现的位置
	var indexes = make(map[int][]int)
	for idx, num := range nums {
		// 注意: 这里存储的不是位置, 而是 pos[i] - i, 方便后续的计算
		indexes[num] = append(indexes[num], idx-len(indexes[num]))
	}

	var ret int
	for _, pos := range indexes {
		if len(pos) <= ret {
			continue
		}
		ret = max(ret, 1)
		// 滑动窗口, 统计让空洞数目小于等于k的最大差距是多少.
		var left int
		for right := 1; right < len(pos); right++ {
			// cnt := pos(right)-pos(left)+1 // 区间内有多少个数
			// total := nums(right)-nums(left)+1 // 总共多少个数
			// total - cnt = nums(right)-nums(left) - pos(right)+pos(left)
			p := pos[right]
			// p = nums(right) - pos(right)
			// pos[left] = nums(left) - pos(left)
			// p - pos[left] = nums(right)-nums(left) - pos(right)+pos(left)
			for p-pos[left] > k {
				left++
			}
			ret = max(ret, right-left+1)
		}
	}

	return ret
}
