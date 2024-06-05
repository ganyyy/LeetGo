package main

func minimumSteps(s string) (ans int64) {
	// 不停的累加左边的1的个数.
	// 可以看成是出现一个0, 就把这个0移动到最左边的1的位置需要交换的次数
	cnt1 := 0
	for _, c := range s {
		if c == '1' {
			cnt1++
		} else {
			ans += int64(cnt1)
		}
	}
	return
}
