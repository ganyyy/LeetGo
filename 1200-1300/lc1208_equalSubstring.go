package main

import "fmt"

func equalSubstring(s string, t string, maxCost int) int {
	//if maxCost == 0 {
	//	// 双指针查一下最大的相同子串
	//	var l, r int
	//	var res int
	//	for r < len(s) {
	//		if s[r] != t[r] {
	//			res = max(res, r-l)
	//			l = r + 1
	//		}
	//		r++
	//	}
	//	res = max(res, r-l)
	//	return res
	//}

	// 然后就是最小代价
	// 过程和上边相同, 但是需要维护一个窗口

	// 当前窗口中, 每个字符的消耗

	// 窗口不是必须的, 可以优化掉
	var window []int

	var l, r, res, sub int
	for r < len(s) {
		if s[r] != t[r] {
			sub = abs(int(s[r]) - int(t[r]))
			if maxCost < sub && len(window) > 0 {
				res = max(res, r-l)
				fmt.Println(s[l:r], t[l:r])
				// 从前边的窗口中获取值, 并将左指针进行位移
				// 这里会尽量保存更多的指针
				for len(window) > 0 && maxCost < sub {
					maxCost += window[0]
					window = window[1:]
					l++
				}
			}
			if maxCost >= sub {
				// 当前位置的元素可以加入到窗口中
				window = append(window, sub)
				maxCost -= sub
			} else {
				// 当前元素不能加入到窗口中, 计算当前可拥有的最大值
				res = max(res, r-l)
				fmt.Println(s[l:r], t[l:r])
				l = r + 1
			}
		} else {
			// 相等的情况下, 添加一个0
			window = append(window, 0)
		}
		r++
	}

	return max(res, r-l)
}

func equalSubstring2(s, t string, maxCost int) int {
	var l, r, res, sub int
	for r < len(s) {
		if s[r] != t[r] {
			res = max(res, r-l)
			sub += abs(int(s[r]) - int(t[r]))
			for maxCost < sub {
				// 当l和r相等时, 此时 maxCost == sub, 会直接跳出循环
				sub -= abs(int(s[l]) - int(t[l]))
				l++
			}
		}
		r++
	}
	return max(res, r-l)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {

	fmt.Println(equalSubstring2("ujteygggjwxnfl", "nstsenrzttikoy", 43))

	var a, b = "gjwxnf", "zttiko"
	var sum int
	for i := range a {
		sum += abs(int(a[i]) - int(b[i]))
	}
	fmt.Println(sum)
}
