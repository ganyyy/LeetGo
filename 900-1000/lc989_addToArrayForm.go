package main

func addToArrayForm(A []int, K int) []int {
	var k = K

	var add int
	for i := len(A) - 1; i >= 0; i-- {
		A[i] += k%10 + add
		if A[i] >= 10 {
			A[i] -= 10
			add = 1
		} else {
			add = 0
		}
		k /= 10

		// k 位数小于 A长度的情况
		if k == 0 && add == 0 {
			break
		}
	}

	var res []int
	for add != 0 || k != 0 {
		var t = k%10 + add
		if t >= 10 {
			add = 1
			t -= 10
		} else {
			add = 0
		}
		k /= 10
		res = append(res, t)
	}

	// 反转
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}

	// 返回结果
	return append(res, A...)
}

func addToArrayForm2(A []int, K int) []int {
	// 每次都去掉K的末尾
	var k = K
	var res = make([]int, 0, len(A))

	var i = len(A) - 1

	for i >= 0 || k > 0 {
		if i >= 0 {
			k += A[i]
		}
		// 每次插入k的末尾
		res = append(res, k%10)
		k /= 10
		i--
	}

	// 反转
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
