package main

import "fmt"

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	// 先pick, 再merge

	var res []int
	for i := 0; i <= k; i++ {
		if i > len(nums1) || k-i > len(nums2) {
			continue
		}
		var t = merge(pick(nums1, i), pick(nums2, k-i))
		if !compare(res, t) {
			res = t
		}
	}
	return res
}

func pick(num []int, n int) []int {
	if n == len(num) {
		return num
	}
	if n == 0 {
		return nil
	}
	var t = make([]int, 0, n)

	// 这里的drop是要删除的数字的个数
	var drop = len(num) - n
	for _, v := range num {
		// 如果栈不为空并且栈顶小于当前值并且要删除的元素不为空, 栈顶出栈
		for len(t) != 0 && t[len(t)-1] < v && drop > 0 {
			t = t[:len(t)-1]
			drop--
		}
		t = append(t, v)
	}
	return t[:n]
}

func merge(num1, num2 []int) []int {
	var res = make([]int, len(num1)+len(num2))
	var c int
	for len(num1) != 0 && len(num2) != 0 {
		if compare(num1, num2) {
			res[c] = num1[0]
			num1 = num1[1:]
		} else {
			res[c] = num2[0]
			num2 = num2[1:]
		}
		c++
	}
	if len(num1) != 0 {
		copy(res[c:], num1)
	}
	if len(num2) != 0 {
		copy(res[c:], num2)
	}
	return res
}

func compare(a, b []int) bool {
	var m = min(len(a), len(b))
	for i := 0; i < m; i++ {
		if a[i] == b[i] {
			continue
		}
		return a[i] > b[i]
	}
	if m != len(a) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(maxNumber([]int{6, 7, 5}, []int{4, 8, 1}, 3))
}
