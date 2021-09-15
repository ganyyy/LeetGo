package main

type trie struct {
	v    byte
	end  bool
	next [26]*trie
}

func (t *trie) insert(s string) {
	if len(s) == 0 {
		return
	}
	t.v = s[0]
	var cur = t
	for i := 1; i < len(s); i++ {
		if cur.get(s[i]) == nil {
			cur.next[s[i]-'a'] = &trie{
				v: s[i],
			}
		}
		cur = cur.get(s[i])
	}
	cur.end = true
}

func (t *trie) get(s byte) *trie {
	if s == '*' {
		return nil
	}
	return t.next[s-'a']
}

type root struct {
	trie
}

func (r *root) insert(s string) {
	var nb = r.next[s[0]-'a']
	if nb == nil {
		nb = &trie{}
		r.next[s[0]-'a'] = nb
	}
	nb.insert(s)
}

var dir = [4][2]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func findWords(board [][]byte, words []string) []string {
	// 需要构建一个字典树吗...
	var r = &root{}
	for _, w := range words {
		r.insert(w)
	}

	var ret []string

	var cur []byte

	var row, col = len(board), len(board[0])

	var check func(i, j int, t *trie)

	check = func(i, j int, t *trie) {
		var b = board[i][j]
		if b == '*' {
			return
		}
		board[i][j] = '*'
		cur = append(cur, b)
		if t.end {
			ret = append(ret, string(cur))
			t.end = false
		}
		for _, d := range dir {
			var x, y = i + d[0], j + d[1]
			if x >= 0 && x < row && y >= 0 && y < col {
				var nt = t.get(board[x][y])
				if nt == nil {
					continue
				}
				check(x, y, nt)
			}
		}
		cur = cur[:len(cur)-1]
		board[i][j] = b
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

	return ret
}

func main() {
	findWords(nil, []string{"abcd", "abce", "abcf"})
}
