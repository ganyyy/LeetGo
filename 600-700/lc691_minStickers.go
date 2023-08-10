package main

func minStickers(stickers []string, target string) int {
	m := len(target)
	f := make([]int, 1<<m)
	for i := range f {
		f[i] = -1
	}
	f[0] = 0
	var dp func(int) int
	dp = func(mask int) int {
		if f[mask] != -1 {
			return f[mask]
		}
		f[mask] = m + 1
		for _, sticker := range stickers {
			left := mask
			cnt := [26]int{}
			for _, c := range sticker {
				cnt[c-'a']++
			}
			for i, c := range target {
				if mask>>i&1 == 1 && cnt[c-'a'] > 0 {
					cnt[c-'a']--
					left ^= 1 << i
				}
			}
			if left < mask {
				f[mask] = min(f[mask], dp(left)+1)
			}
		}
		return f[mask]
	}
	ans := dp(1<<m - 1)
	if ans <= m {
		return ans
	}
	return -1
}
