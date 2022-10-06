package main

func threeEqualParts(arr []int) []int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%3 != 0 {
		return []int{-1, -1}
	}
	if sum == 0 {
		return []int{0, 2}
	}

	// 三等分
	// 从每一个部分首个不为0的位置开始
	partial := sum / 3
	first, second, third, cur := 0, 0, 0, 0
	for i, x := range arr {
		if x == 1 {
			if cur == 0 {
				first = i
			} else if cur == partial {
				second = i
			} else if cur == 2*partial {
				third = i
			}
			cur++
		}
	}

	// 末尾0的个数一定是相同的, 所以third-end这段区间决定有效的位数
	n := len(arr)
	length := n - third
	if first+length <= second && second+length <= third {
		i := 0
		for third+i < n {
			// 对比三个部分的值是否相同
			if arr[first+i] != arr[second+i] || arr[first+i] != arr[third+i] {
				return []int{-1, -1}
			}
			i++
		}
		return []int{first + length - 1, second + length}
	}
	return []int{-1, -1}
}
