package main

import (
	"fmt"
	"math/bits"
)

func findNumOfValidWords(words []string, puzzles []string) []int {
	const puzzleLength = 7
	cnt := map[int]int{}
	for _, s := range words {
		// mash当成是一种hash,
		// 因为 aaaa -> aaa -> aa -> a
		mask := 0
		for _, ch := range s {
			mask |= 1 << (ch - 'a')
		}
		// 最多存在7种不同的数字, 所以如果大于7就没必要查找了
		if bits.OnesCount(uint(mask)) <= puzzleLength {
			cnt[mask]++
		}
	}

	ans := make([]int, len(puzzles))
	for i, s := range puzzles {
		// 首先必须要保证存在首字母
		first := 1 << (s[0] - 'a')

		// 枚举子集方法一
		//for choose := 0; choose < 1<<(puzzleLength-1); choose++ {
		//    mask := 0
		//    for j := 0; j < puzzleLength-1; j++ {
		//        if choose>>j&1 == 1 {
		//            mask |= 1 << (s[j+1] - 'a')
		//        }
		//    }
		//    ans[i] += cnt[mask|first]
		//}

		// 枚举子集方法二
		mask := 0
		for _, ch := range s[1:] {
			mask |= 1 << (ch - 'a')
		}
		subset := mask
		for {
			ans[i] += cnt[subset|first]
			if subset == 0 {
				break
			}
			subset = (subset - 1) & mask
		}
	}
	return ans
}

func main() {
	var mask = 0b1111
	var subset = 0b111

	for {
		fmt.Printf("%b\n", subset)
		if subset == 0 {
			break
		}
		subset = (subset - 1) & mask
	}
}
