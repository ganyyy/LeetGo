package main

func prevPermOpt1(arr []int) []int {
	for i := len(arr) - 2; i >= 0; i-- {
		// 找到第一个不满足 递增的位置 i
		// 1, (9), 4, 6, 7
		if arr[i] > arr[i+1] {
			// 从 [i:] 中, 选取出 最接近 i 且 小于 [i] 的 j
			// 只有 [j] < [i] 才能保证字典序小
			// 只有 最接近 i 的 j, 才可以保证 字典序 最接近
			j := len(arr) - 1
			for arr[j] >= arr[i] || arr[j] == arr[j-1] {
				j--
			}
			arr[i], arr[j] = arr[j], arr[i]
			break
		}
	}
	return arr
}
