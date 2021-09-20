package main

import "fmt"

// 再优化一下

type trie struct {
	next [26]*trie // 子节点
	word string    // 以该节点为结尾的字符串
	cnt  int32     // 该节点被使用了多少次
	end  bool      // 是不是尾节点
}

func (t *trie) insert(s string) {
	if len(s) == 0 {
		return
	}
	var cur = t
	for i := 0; i < len(s); i++ {
		if cur.get(s[i]) == nil {
			cur.next[s[i]-'a'] = &trie{}
		}
		cur = cur.get(s[i])
		cur.cnt++ // 每经过一次, 都会增加一下计数
	}
	cur.word = s
	cur.end = true
}

func (t *trie) delete(s string) {
	var cur = t
	for i := range s {
		cur = cur.get(s[i])
		cur.cnt--
	}
	cur.end = false
}

func (t *trie) get(s byte) *trie {
	return t.next[s-'a']
}

type root struct {
	trie
}

var dir = [4][2]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

const (
	POW  = 32
	MASK = POW - 1
)

type bitset []uint32

func (b bitset) check(v int) bool {
	return b[v/POW]&(1<<(v&MASK)) != 0
}

func (b bitset) set(v int) {
	b[v/POW] |= 1 << (v & MASK)
}

func (b bitset) del(v int) {
	b[v/POW] &= ^(1 << (v & MASK))
}

func findWords(board [][]byte, words []string) []string {
	// 需要构建一个字典树吗...
	var r = &root{}
	for _, w := range words {
		r.insert(w)
	}

	var ret []string

	var row, col = len(board), len(board[0])

	var use bitset = make([]uint32, (row*col+MASK)/POW)

	var check func(i, j int, t *trie)

	check = func(i, j int, t *trie) {
		if t == nil || t.cnt == 0 { // 提前枝减一下, 该路径已经走过了, 不能重复利用第二次
			return
		}
		var pos = i*col + j
		if use.check(pos) {
			return
		}
		use.set(pos)
		if t.end {
			ret = append(ret, t.word)
			r.delete(t.word)
		}
		for _, d := range dir {
			var x, y = i + d[0], j + d[1]
			if x >= 0 && x < row && y >= 0 && y < col {
				check(x, y, t.get(board[x][y]))
			}
		}
		use.del(pos)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			var t = r.get(board[i][j])
			if t == nil {
				continue
			}
			check(i, j, t)
		}
	}

	//fmt.Println(use)

	return ret
}

/*
[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]
["oath","pea","eat","rain"]
*/

func main() {
	fmt.Println(findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
		[]string{"oath", "pea", "eat", "rain"}))
}
