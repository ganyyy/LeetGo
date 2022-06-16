package main

func duplicateZeros(arr []int) {
	n := len(arr)
	// top是复写0后的数组长度
	top := 0
	i := -1
	// 统计满足要求的终点位置i
	for top < n {
		i++
		if arr[i] != 0 {
			top++
		} else {
			top += 2
		}
	}
	// 针对末尾0特殊处理
	j := n - 1
	// 如果末尾不是0, 那么top == n
	// 否则top == n+1
	if top == n+1 {
		arr[j] = 0
		j--
		i--
	}
	// 倒叙复写
	for j >= 0 {
		arr[j] = arr[i]
		j--
		if arr[i] == 0 {
			arr[j] = arr[i]
			j--
		}
		i--
	}
}
