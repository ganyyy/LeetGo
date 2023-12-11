package main

const (
	InvalidByte byte = 0xFF
	InitialByte byte = 0
	LastByte    byte = 'k'

	InvalidFrog = -1
)

func PreByte(b byte) byte {
	switch b {
	case 'c':
		return InitialByte
	case 'r':
		return 'c'
	case 'o':
		return 'r'
	case 'a':
		return 'o'
	case LastByte:
		return 'a'
	default:
		return InvalidByte
	}
}

func minNumberOfFrogs(croakOfFrogs string) int {
	var count [128]int
	var remain int // 某一时刻未匹配成功的字节数
	var total int  // 所需要的最多的
	for i := range croakOfFrogs {
		b := croakOfFrogs[i]
		pre := PreByte(b)
		if pre == InvalidByte {
			return InvalidFrog
		}
		if pre != InitialByte {
			// 存在前置的情况下, 需要保证前置数量可以满足条件
			if count[pre] <= 0 {
				return InvalidFrog
			}
			count[pre]--
			remain--
		}
		if b != LastByte {
			count[b]++
			remain++
		} else {
			// 剩余的字符, 表示的是未完成匹配的数量, 也就是所需要的最大的数量
			// 当前的k已经闭合了一组单词, 所以需要额外加一下
			total = max(total, remain+1)
		}
	}
	if remain != 0 {
		// 存在剩余的未匹配的字符
		return InvalidFrog
	}
	return total
}
