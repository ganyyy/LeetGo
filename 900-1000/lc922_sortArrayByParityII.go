package main

func sortArrayByParityII(A []int) []int {
	l, r := 0, 1
	for {
		for l < len(A) && A[l]%2 == 0 {
			l += 2
		}

		for r < len(A) && A[r]%2 == 1 {
			r += 2
		}

		if l > len(A) || r > len(A) {
			break
		}

		A[l], A[r] = A[r], A[l]

	}
	return A
}
