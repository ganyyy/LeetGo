package main

import "fmt"

// 标记一下, 又臭又长的代码. 比写业务逻辑还烦人

func fullJustify(words []string, maxWidth int) []string {
	// 最小配置为单词数量n + 单词数量n-1 个空格 = maxWidth
	// 如果第n+1个单词 + 前边的长度 > maxWidth, 则剩余的空格数量
	// 为 maxWidth - sum(len(words[0:n+1]))
	// 剩余的空格分配问题:
	//  1. 优先平分
	//     white % n-1 == 0
	//  2. 左边数量>=右边数量
	// 最后一行要求左对齐, 即每个单词之间留一个空格即可

	res := make([]string, 0)

	genSeq := func(wordSeq []string, numSeq []int) {
		seq := make([]byte, maxWidth)
		index, wordIndex, numIndex := 0, 0, 0
		for {
			if wordIndex < len(wordSeq) {
				w := wordSeq[wordIndex]
				for i := 0; i < len(w); i++ {
					seq[index] = w[i]
					index++
				}
				wordIndex++
			}
			if numIndex < len(numSeq) {
				for i := 0; i < numSeq[numIndex]; i++ {
					seq[index] = ' '
					index++
				}
				numIndex++
			} else {
				break
			}
		}
		// 结尾的
		for i := index; i < maxWidth; i++ {
			seq[i] = ' '
		}
		res = append(res, string(seq))
	}
	// start表示每句话开头的单词的索引
	// i 表示当前是第几个单词
	// count 当前句子最小长度(n个单词长度)
	start, i, count := 0, 0, 0
	for i < len(words) {
		// 单词的数量+最小空格数量
		w := words[i]
		if count+len(w)+i-start <= maxWidth {
			count += len(w)
			i++
		} else {
			wordCount := i - start
			seq := make([]byte, maxWidth)
			if wordCount == 1 {
				// 只有一个单词
				w, k := words[start], 0
				for ; k < len(w); k++ {
					seq[k] = w[k]
				}
				for k < maxWidth {
					seq[k] = ' '
					k++
				}
				res = append(res, string(seq))
			} else {
				// 实际的空格数量
				white := maxWidth - count
				whiteCount := wordCount - 1
				whiteSeq := make([]int, whiteCount)
				if white%whiteCount == 0 {
					// 正好分配完
					num := white / whiteCount
					for k := 0; k < whiteCount; k++ {
						whiteSeq[k] = num
					}
					genSeq(words[start:i], whiteSeq)
				} else {
					// 有的多有的少
					num := white / whiteCount
					remain := white % whiteCount
					for k := 0; k < whiteCount; k++ {
						if k < remain {
							whiteSeq[k] = num + 1
						} else {
							whiteSeq[k] = num
						}
					}
					// 填充字符串
					genSeq(words[start:i], whiteSeq)
				}
			}
			count = 0
			start = i
		}
	}
	// 结尾的处理
	// 到最后了还有剩下的, 就把剩下的添加进去
	// 最后的长度一定是 <= maxWidth 的
	seq := make([]byte, maxWidth)
	index := 0
	for k := start; k < len(words); k++ {
		word := words[k]
		for j := 0; j < len(word); j++ {
			seq[index] = word[j]
			index++
		}
		if index < maxWidth {
			seq[index] = ' '
			index++
		}
	}
	for index < maxWidth {
		seq[index] = ' '
		index++
	}
	res = append(res, string(seq))
	return res
}

func main() {
	res := fullJustify([]string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}, 16)
	for _, str := range res {
		fmt.Println(len(str))
		fmt.Println(str)
	}
}
