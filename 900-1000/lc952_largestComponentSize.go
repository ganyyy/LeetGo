package main

// https://leetcode.cn/endlesscheng
const mx int = 1e5

var pf [mx + 1][]int32

func init() {
	// 预处理所有的根节点
	for i := 2; i <= mx; i++ {
		if pf[i] == nil {
			// 2->4->6->8->10 ==> 2
			// 3->6->9 ==> 3
			for j := i; j <= mx; j += i {
				pf[j] = append(pf[j], int32(i))
			}
		}
	}
}

var fa [mx + 1]int32

func largestComponentSize(nums []int) int {
	for i := range fa {
		fa[i] = int32(i)
	}
	var find func(int32) int32
	find = func(x int32) int32 {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	// 2.4.6.8.10 => 2
	// 4,8 => 2,4
	// ...etc
	for _, v := range nums {
		for _, p := range pf[v] {
			fa[find(p)] = find(int32(v))
		}
	}
	cnt := [mx + 1]int16{}
	ans := int16(0)
	for _, v := range nums {
		v := find(int32(v))
		cnt[v]++
		if cnt[v] > ans {
			ans = cnt[v]
		}
	}
	return int(ans)
}

func main() {

}
