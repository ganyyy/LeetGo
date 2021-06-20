package main

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	// 二分?

	// 中值怎么取

	// 末尾至少要预留一位

	// [3, base-'1']

	// 大小界定

	// 关键是怎么转换? 如何运算?

	// 想不明白啊

	//

	// 要实现一个字符串的加减乘算法?

	// 丫的, 能转成数字啊!
	var i, _ = strconv.ParseInt(n, 10, 64)

	// 转换成进制数表示其等比序列求和的值为 n

	// k 为最优进制
	// m 为该进制所需要的位数
	// n = k^(m-1) + k^(m-2) + ... k^1 + k^0
	// n = (k^m - 1)/(k-1)
	// 要求k尽量小, 那么m就要尽量的大
	// 当k为2的时候, m最大. 此时 m的值为 log2(n+1)
	// m最小为2位数, 即 1,1

	// 外围确定M的值
	for m := int(math.Log2(float64(i + 1))); m >= 2; m-- {
		// 内部确认K的值
		// 这里的上限是通过
		var l, r int64 = 2, int64(math.Pow(float64(i), 1/float64(m-1))) + 1

		for l < r {
			var mid = l + (r-l)>>1

			// 对当前m和k进行求和
			var sum int64
			for j := 0; j < m; j++ {
				sum = sum*mid + 1
			}

			if sum == i {
				return strconv.Itoa(int(mid))
			} else if sum > i {
				r = mid
			} else {
				l = mid + 1
			}
		}
	}

	return strconv.Itoa(int(i - 1))
}
func main() {
	println(smallestGoodBase("13"))
}
