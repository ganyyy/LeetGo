package main

import "fmt"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var res [][]string
	// 对字符串数组进行去重, 题目中声明是是不会存在相同的字符串, 这里先不检测
	// 每一个单词都分配一个唯一ID(就是当前索引即可), 构建一个单词到ID间的映射关系
	wordToID := make(map[string]int, len(wordList))
	for index, v := range wordList {
		wordToID[v] = index
	}
	// 判断一下结尾字符是否存在于 映射中
	if _, ok := wordToID[endWord]; !ok {
		return res
	}

	// 看看开启字符串是否在集合中
	if _, ok := wordToID[beginWord]; !ok {
		wordToID[beginWord] = len(wordList)
		wordList = append(wordList, beginWord)
	}

	// 理论上, 这个和修改后的 wordList长度应该是一样的
	lm := len(wordToID)

	// 通过一个二维数组来保存双向边
	edges := make([][]int, lm)
	for i := 0; i < lm; i++ {
		edges[i] = make([]int, 0)
	}

	// 构建图
	// 两个单词之间的编辑距离为1, 那就可以认为他们之间存在一条双向边
	for i := 0; i < lm; i++ {
		for j := i + 1; j < lm; j++ {
			if diffOne(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}

	// 终点的位置
	end := wordToID[endWord]
	// 到每一个点的代价, 初始化为无穷大
	cost := make([]int, lm)
	INF := 1 << 31
	for i := 0; i < lm; i++ {
		cost[i] = INF
	}

	// 创建队列, 这个是保存当前节点的所有边的队列
	queue := make([][]int, 0)
	id := wordToID[beginWord]
	tmpBegin := []int{
		id,
	}
	queue = append(queue, tmpBegin)
	// 起点的开销肯定为0
	cost[id] = 0

	for len(queue) != 0 {
		// 第一层出队
		now := queue[0]
		queue = queue[1:]
		// 获取当前访问的最后一个节点
		last := now[len(now)-1]

		// 到终点了
		if last == end {
			tmp := make([]string, len(now))
			for i, v := range now {
				tmp[i] = wordList[v]
			}
			res = append(res, tmp)
		} else {
			// 遍历 最后位置的所有邻接边
			edge := edges[last]
			for _, to := range edge {
				// 保留所有相同代价的不同路径
				if cost[last]+1 <= cost[to] {
					// 更新下一个点的消耗
					cost[to] = cost[last] + 1
					// 将下一个点添加到当前路径中
					tmp := make([]int, len(now)+1)
					copy(tmp, now)
					tmp[len(now)] = to
					queue = append(queue, tmp)
				}
			}
		}
	}
	return res
}

// 检测是否差距一个编辑距离
func diffOne(a, b string) bool {
	var diff int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 1
}

func main() {
	old := []int{
		1, 2, 3, 4, 5,
	}

	tmp := make([]int, 0, 5)
	fmt.Println(copy(tmp, old))
	fmt.Println(tmp)

	tmp = make([]int, 5, 5)
	fmt.Println(copy(tmp, old))
	fmt.Println(tmp)
}
