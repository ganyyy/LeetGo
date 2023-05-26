package main

func sampleStats(count []int) []float64 {
	n := len(count)
	total := 0
	for i := 0; i < n; i++ {
		total += count[i]
	}
	mean := 0.0
	median := 0.0
	minimum := 256
	maxnum := 0
	mode := 0

	left := (total + 1) / 2
	right := (total + 2) / 2
	cnt := 0
	maxfreq := 0
	sum := 0
	for i := 0; i < n; i++ {
		sum += int(count[i]) * int(i)
		if count[i] > maxfreq {
			// 最大频率
			maxfreq = count[i]
			mode = i
		}
		if count[i] > 0 {
			// 最小值
			if minimum == 256 {
				minimum = i
			}
			// 最大值
			maxnum = i
		}
		if cnt < right && cnt+count[i] >= right {
			// 如果 i 横跨了 right 和 left, 那么他就可能是中位数的一部分
			median += float64(i)
		}
		if cnt < left && cnt+count[i] >= left {
			// 如果 i 横跨了 right 和 left, 那么他就可能是中位数的一部分
			median += float64(i)
		}
		cnt += count[i]
	}
	mean = float64(sum) / float64(total)
	// 一定会添加两次, 所以这里需要除一下
	median = median / 2.0
	return []float64{float64(minimum), float64(maxnum), mean, median, float64(mode)}
}
