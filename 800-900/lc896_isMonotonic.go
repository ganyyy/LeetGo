package main

func isMonotonic(A []int) bool {
	if len(A) <= 2 {
		return true
	}

	// 找到第一个不相等的数
	var l = 1
	for l < len(A) {
		if A[l] != A[l-1] {
			break
		}
		l++
	}

	if l == len(A) {
		return true
	}

	var small = A[l-1] < A[l]

	for l < len(A) {
		if A[l] != A[l-1] {
			if A[l-1] < A[l] != small {
				return false
			}
		}
		l++
	}

	return true
}

// 瞻仰一下大佬的操作
func isMonotonic2(A []int) bool {
	// Ok2 表示是递增序列
	// Ok1 表示是递减序列
	// 二者不能同时存在
	length, OK1, OK2 := len(A), true, true
	for i := 1; i < length; i++ {
		if OK2 && A[i-1] > A[i] {
			OK2 = false
		} else if OK1 && A[i-1] < A[i] {
			OK1 = false
		}
	}
	if OK1 || OK2 {
		return true
	}
	return false
}
