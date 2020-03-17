package main

import "fmt"

func lengthLongestPath(input string) int {
	var depth [30]int // 每一级的最大长度
	var res int
	input += "\n" // 强制遍历到最后一位
	// 标记是不是文件
	// 当前等级
	// 当前文件(夹)长度
	isFile, level, cur := false, 0, 0
	for i := 0; i < len(input); i++ {
		c := input[i]
		// 如果是换行, 就将level归零
		if c == '\n' {
			if level > 0 {
				depth[level] = depth[level-1] + cur + 1
			} else {
				depth[level] = cur
			}
			if isFile {
				res = max(res, depth[level])
				isFile = false
			}
			cur = 0
			level = 0
			continue
		}
		if c == '\t' {
			level++
			continue
		}
		if c == '.' {
			isFile = true
		}
		cur++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
}
