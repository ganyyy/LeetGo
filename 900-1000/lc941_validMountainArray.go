package main

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	// 先找到拐点
	var i int
	for i = 0; i < len(A)-1; i++ {
		if A[i] >= A[i+1] {
			break
		}
	}

	// 单调递增或者单调递减 都不符合要求
	if i == len(A)-1 || i == 0 {
		return false
	}
	for ; i < len(A)-1; i++ {
		if A[i] <= A[i+1] {
			return false
		}
	}
	return true
}
