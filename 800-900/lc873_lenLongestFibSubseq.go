package main

func lenLongestFibSubseq(arr []int) int {
	var m = make(map[int]int)
	for i, v := range arr {
		m[v] = i
	}
	var ret int
	for i := 1; i < len(arr); i++ {
		var v, hv = arr[i], arr[i] / 2
		for j := i - 1; j >= 0 && arr[j] > hv; j-- {
			var cur = v
			var pre = arr[j]
			// var total = []int{cur, pre}
			var round int
			pIdx, ok := m[cur-pre]
			for ok {
				cur, pre = pre, arr[pIdx]
				if pre >= cur {
					break
				}
				// total = append(total, pre)
				round++
				pIdx, ok = m[cur-pre]
			}
			if round != 0 {
				// fmt.Println(total)
				ret = max(ret, round+2)
			}
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	lenLongestFibSubseq([]int{2, 4, 5, 6, 7, 8, 11, 13, 14, 15, 21, 22, 34})
}
