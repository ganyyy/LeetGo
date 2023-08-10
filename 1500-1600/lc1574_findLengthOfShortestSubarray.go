package main

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	j := n - 1
	// 从后向前的最长递减子数组
	for j > 0 && arr[j-1] <= arr[j] {
		j--
	}
	if j == 0 {
		return 0
	}
	// 默认前边全部删除
	res := j
	for i := 0; i < n; i++ {
		// [j] > [i]
		for j < n && arr[j] < arr[i] {
			j++
		}
		// 更新窗口的最小值
		res = min(res, j-i-1)
		// 找到第一个非递增的位置
		if i+1 < n && arr[i] > arr[i+1] {
			break
		}
	}
	return res
}
