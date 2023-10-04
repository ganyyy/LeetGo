package main

import (
	"fmt"
	"math"
	"unsafe"
)

// 换bfs试试, 首次BFS, 内存溢出了
var empty127 = struct{}{}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	if len(beginWord) != len(endWord) {
		return 0
	}

	// 换双端搜索, 每次都从备选词库较少的那一方进行筛选
	// begin -> ... tmp ... -> end
	// 初始情况下, begin只包含 beginWord
	// end只包含endWord. 经过一次迭代后, 会将当前beginSet所有的合法变化放入到tmpSet中
	// 因为变化是双向的, 所以每次遍历要从 数量较短的一方进行切入. 可以有效地降低时间复杂度
	// 同样的,
	// 双指针的变形记

	// 所有单词的集合
	var dict = make(map[string]struct{}, len(wordList))
	for _, w := range wordList {
		if len(w) != len(beginWord) {
			continue
		}
		dict[w] = empty127
	}

	// 首先判断一下结束单词是否在备选词库中
	if _, ok := dict[endWord]; !ok {
		return 0
	}

	// 启用双端搜索

	// 开始的集合
	var beginSet = map[string]struct{}{
		beginWord: empty127,
	}
	// 结束的集合
	var endSet = map[string]struct{}{
		endWord: empty127,
	}

	var buffer = make([]byte, len(beginWord))

	// 总的步数
	var level = 1

	// 由beginSet转移到下一层的集合
	var nextSet = map[string]struct{}{}
	for len(beginSet) != 0 {
		clear(nextSet)
		// 层次+1
		level++
		// 从备选词库中去掉开始集合中存在的元素, 防止重复遍历
		for s := range beginSet {
			delete(dict, s)
		}
		// 从剩余备选词中获取可转移的对象放入到 临时集合中
		for bs := range beginSet {
			// 挨个位置, 挨个字母的进行替换. 进行过滤和对比, 将时间复杂度降低到可控范围(len(bs)*26)
			for i := 0; i < len(bs); i++ {
				copy(buffer, bs)
				for c := 'a'; c <= 'z'; c++ {
					buffer[i] = byte(c)
					// 如果已经过滤掉了
					var s = toString127(buffer)
					if _, ok := dict[s]; !ok {
						continue
					}
					// 如果在endSet中, 返回变化的步骤
					if _, ok := endSet[s]; ok {
						return level
					}
					// 不符合上述条件的直接放入到临时集合中, 等待下一次的查找
					// 这里需要进行一次copy, 不然就会出现重复的情况
					nextSet[string(buffer)] = empty127
				}
			}
		}
		// 比较nextSet和endSet的大小,
		// 如果nextSet较小, 从nextSet向endSet进行查找
		// 否则从 endSet 向 nextSet 进行查找
		if len(nextSet) < len(endSet) {
			beginSet, nextSet = nextSet, beginSet
		} else {
			beginSet, endSet, nextSet = endSet, nextSet, beginSet
		}
	}

	// 找不到的情况
	return 0
}

func toString127(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func ladderLengthErrMemory(beginWord string, endWord string, wordList []string) int {
	var dict = make(map[string]map[string]bool, len(wordList))

	var mk func(s string)

	mk = func(s string) {
		var m = make(map[string]bool, len(wordList))
		// 结尾u不需要处理
		if s == endWord {
			dict[s] = m
			return
		}
		for _, t := range wordList {
			if t == s {
				continue
			}
			if checkDiff(s, t) {
				m[t] = true
			}
		}
		dict[s] = m
	}

	// 构建一个转移映射关系
	mk(beginWord)
	for _, s := range wordList {
		mk(s)
	}

	// 如果找不到结束单词..
	if _, ok := dict[endWord]; !ok {
		return 0
	}

	// 换bfs了...

	// 当前的层级
	var level = 1
	// 队列
	var queue = make([]string, 0, len(wordList))
	queue = append(queue, beginWord)
	// 标记是否查找过
	var check = make(map[string]struct{}, len(wordList))

	// 当前队列的长度
	var ln int
	// 队头字符串
	var cur string
	// 对比字符串
	var s string
	// 是否已经查看过的标记
	var ok bool

	for len(queue) != 0 {
		ln = len(queue)
		// 遍历当前队列的所有元素
		for i := 0; i < ln; i++ {
			cur = queue[0]
			if cur == endWord {
				return level
			}
			// 出队
			queue = queue[1:]
			// 标记
			check[cur] = empty127
			// 获取当前元素的所有可以转移的目标
			for s = range dict[cur] {
				if _, ok = check[s]; ok {
					continue
				}
				// 标记
				check[s] = empty127
				// 入队
				queue = append(queue, s)
			}
		}
		// 表示完成了一次转移, 层次++
		level++
	}
	// 找不到
	return 0
}

// 超时了, dfs
func ladderLengthErr(beginWord string, endWord string, wordList []string) int {
	var dict = make(map[string]map[string]bool, len(wordList))

	var mk func(s string)

	mk = func(s string) {
		var m = make(map[string]bool, len(wordList))
		// 结尾u不需要处理
		if s == endWord {
			dict[s] = m
			return
		}
		for _, t := range wordList {
			if t == s {
				continue
			}
			if checkDiff(s, t) {
				m[t] = true
			}
		}
		dict[s] = m
	}

	// 构建一个转移映射关系
	mk(beginWord)
	for _, s := range wordList {
		mk(s)
	}

	// 如果找不到结束单词..
	if _, ok := dict[endWord]; !ok {
		return 0
	}

	var res = math.MaxInt32
	var dfs func(s string, cur int)

	dfs = func(s string, cur int) {
		if s == endWord {
			// 找到终点就更新一下最小值
			if cur < res {
				res = cur
			}
			return
		}
		var m = dict[s]
		// 从可到达路径中查找所有可能到达的地方
		for v, valid := range m {
			if !valid {
				continue
			}
			// 先清理, 防止无限的dfs
			m[v] = false
			dfs(v, cur+1)
			m[v] = true
		}
	}

	dfs(beginWord, 1)

	if res == math.MaxInt32 {
		return 0
	}

	return res
}

func checkDiff(a, b string) bool {
	var diff int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if diff == 1 {
				return false
			}
			diff++
		}
	}
	return diff == 1
}

func main() {
	fmt.Println(ladderLength(
		"hit",
		"cog",
		[]string{
			"hot", "dot", "dog", "lot", "log", "cog",
		}))
}
