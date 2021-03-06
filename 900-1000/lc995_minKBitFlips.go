package main

func minKBitFlips(A []int, K int) int {
	// 0 要反转奇数次
	// 1 要反转偶数次

	// 本质上, 如果数组开头的值1, 就不应该进行反转
	// 每一个从i开始大小为K的数组, 要么反转0次, 要么反转1次. 这样才可以保证反转的次数最少

	// 临时数组, 初始状态下都是0
	var hint = make([]int, len(A))

	// res: 需要反转的次数
	// flip: 对比值, 反转后更新( ^= 1). 0表示反转了偶数次, 1表示反转了奇数次
	var res, flip int

	for i := 0; i < len(A); i++ {
		// 判断当前是否需要进行反转

		// hint[i]只会在i是某个窗口的结束事件时才会成为1, 表示进行了一次反转
		// 默认情况下是0, 表示没有进行反转

		// 如果没有进行过反转, 那么flip ^ 0 就不会发生改变. 维持上一次反转的结果
		flip ^= hint[i]

		// 如果A[i]是需要进行反转
		// 两种情况: 都需要更新下flip和结束事件的值
		// 奇数次反转的1, 偶数次反转的0
		if A[i] == flip {
			// 反转次数+1
			res++
			if i+K > len(A) {
				// i+K-1 = 窗口K
				// 如果超出了这个范围, 说明当前窗口不足K个数字, 无法完成反转
				return -1
			}
			// i就是这个窗口的开始事件, i+k是这个窗口的结束事件

			// 从当前开始, flip表示进行了一次反转, 直到 i+K 位置时会恢复成旧值
			// flip ^ 1 ^ 1 = flip
			flip ^= 1
			if i+K < len(A) {
				hint[i+K] ^= 1
			}
		}
	}
	return res
}

func minKBitFlips2(A []int, K int) int {
	// 优化一下空间复杂度
	// 基于1进行优化

	var res, flip int

	for i := 0; i < len(A); i++ {
		// 这里判断一下i-K >= 0
		// 如果满足条件再看一下 A[i-K] > 1
		// 如果都满足, 说明需要将flip值回复过去
		if i-K >= 0 && A[i-K] > 1 {
			flip ^= 1
			A[i-K] -= 2
		}
		// 说明需要进行反转
		if A[i] != flip {
			continue
		}

		// 判断一下数组是否发生了越界
		if i+K > len(A) {
			return -1
		}

		// 当前需要反转
		res++
		flip ^= 1
		A[i] += 2
	}

	return res
}
