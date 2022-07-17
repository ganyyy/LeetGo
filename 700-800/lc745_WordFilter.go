//go:build ignore

package main

type Node struct {
	Child [26]*Node
	Idx   []int
}

func positive(str string, cb func(c int) bool) {
	for _, v := range str {
		if !cb(int(v) - 'a') {
			return
		}
	}
}

func reverse(str string, cb func(c int) bool) {
	for i := len(str) - 1; i >= 0; i-- {
		if !cb(int(str[i]) - 'a') {
			return
		}
	}
}

func (n *Node) Add(idx int, str string, r bool) {
	var cur = n

	var loop = func(c int) bool {
		var next = cur.Child[c]
		if next == nil {
			next = &Node{}
			cur.Child[c] = next
		}
		next.Idx = append(next.Idx, idx)
		cur = next
		return true
	}

	if !r {
		positive(str, loop)
	} else {
		reverse(str, loop)
	}
}

func (n *Node) Search(str string, r bool) []int {
	var cur = n
	var find = true
	var loop = func(c int) bool {
		var next = cur.Child[c]
		if next == nil {
			find = false
			return false
		}
		cur = next
		return true
	}
	if !r {
		positive(str, loop)
	} else {
		reverse(str, loop)
	}
	if !find {
		return nil
	}
	return cur.Idx
}

type WordFilter struct {
	Pre  *Node
	Suff *Node
}

func Constructor(words []string) WordFilter {
	// 我的评价是: 两个TRIE又不是不能用
	var filter = WordFilter{
		Pre:  &Node{},
		Suff: &Node{},
	}

	for i, str := range words {
		filter.Pre.Add(i, str, false)
		filter.Suff.Add(i, str, true)
	}

	return filter
}

func (this *WordFilter) F(pref string, suff string) int {
	var p = this.Pre.Search(pref, false)
	// fmt.Println(pref, p)
	if p == nil {
		return -1
	}
	var s = this.Suff.Search(suff, true)
	// fmt.Println(suff, s)
	if s == nil {
		return -1
	}
	var l, r int
	r = len(s) - 1
	for l = len(p) - 1; l >= 0 && r >= 0; {
		if p[l] == s[r] {
			return p[l]
		} else if p[l] > s[r] {
			l--
		} else {
			r--
		}
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
