//go:build ignore

package main

import "fmt"

func singleNumber(nums []int) int {
	// mark
	// 核心是 - 逢3变0

	// 如果是两个清零, 直接异或就可
	// x^0 = x
	// x^x = 0
	var a, b int
	for _, v := range nums {
		// 对于三次变零的处理
		// x[i] 分别放到 a[i]和b[i]上进行处理
		// 第一次遇到: a[i]=0, b[i]=1
		// 第二次遇到: a[i]=1, b[i]=0
		// 第三次遇到: a[i]=0, b[i]=0

		// 这里使用2来举例
		// 0010
		// b     a
		// 0010  0000 第一次
		// 0000  0010 第二次
		// 0000  0000 第三次

		// 所以如果一个数只出现一次, 那么就会存在于b中
		b = (b ^ v) &^ a
		a = (a ^ v) &^ b
	}
	return b
}

// 解法2: 判断每一位上为1的个数, 如果个数不是3的倍数, 那么不同的数这一位一定是1
// 这个可以扩充到任意个相同数字, 只需要修改count取余的值即可
func singleNumber2(nums []int) int {
	var res int
	for i := 0; i < 32; i++ {
		var mask, count = 1 << i, 0
		for _, v := range nums {
			if v&mask != 0 {
				count++
			}
		}
		if count%3 != 0 {
			res |= mask
		}
	}
	// 需要注意的是: golang的int可能是32位的, 也可能是64位的
	return int(int32(res))
}

func singleNumber3(nums []int) int {
	// 核心是 - 逢3变0

	// 如果是两个清零, 直接异或就可
	// x^0 = x
	// x^x = 0
	var a, b int
	for _, v := range nums {
		// 对于三次变零的处理
		// x[i] 分别放到 a[i]和b[i]上进行处理
		// 第一次遇到: a[i]=0, b[i]=1
		// 第二次遇到: a[i]=1, b[i]=0
		// 第三次遇到: a[i]=0, b[i]=0

		// 这里使用2来举例
		// 0010
		// b     a
		// 0010  0000 第一次
		// 0000  0010 第二次
		// 0000  0000 第三次

		// 所以如果一个数只出现一次, 那么就会存在于b中
		// 第一次放在b中, 第二次转移到a中并从b中清除, 第三次从a中清除

		// 如果b是0, a是0, 那么就是第一次遇到, 放到b
		// 如果b是1, a是0, 那么就是第二次遇到, 清空b, 放到a
		// 如果b是0, a是1, 那么就是第三次遇到, 清空a
		b = (b ^ v) &^ a
		a = (a ^ v) &^ b
	}
	return b
}

func main() {
	// 123
	fmt.Println(singleNumber([]int{-1, 3, 3, 3}))
}
