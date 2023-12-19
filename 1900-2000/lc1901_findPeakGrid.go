package main

func findPeakGrid(mat [][]int) []int {
	maxIndex := func(row []int) (int, int) {
		if len(row) == 0 {
			return 0, -1
		}
		var mIdx int
		for i, v := range row[1:] {
			if v > row[mIdx] {
				mIdx = i + 1
			}
		}
		return row[mIdx], mIdx
	}

	// 原理是啥呢?
	// 其实找到一个封顶就行了, 不代表是最大值, 但是一定是一个封顶
	// 假设通过二分确定的行是第i行, 那么对比第i行的最大值和下一行的最大值,
	// 如果小于, 则说明从[i, n-1]这个区间内一定存在一个封顶,

	var left, right = 0, len(mat) - 1
	for left < right {
		mid := left + (right-left)/2
		v, idx := maxIndex(mat[mid])
		if v > mat[mid+1][idx] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	_, col := maxIndex(mat[left])
	return []int{left, col}
}
