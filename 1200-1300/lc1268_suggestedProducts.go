//go:build ignore

package main

import "sort"

func suggestedProducts(products []string, searchWord string) [][]string {
	// 提前排序，题意要求按字典序选取
	sort.Strings(products)
	var ans [][]string
	for i := range searchWord {
		var tmp []string
		// 有点类似字典树的思路
		// 就是计算前缀长度为k相同的，可以从上一轮的产品列表里面找
		// 因为上一轮的产品列表都是前缀长度为k - 1相同的
		// 因此只要比较第k位的字符即可
		// 当然，对于第1轮，我们就老实遍历整个产品列表就好了
		if i == 0 {
			for _, p := range products {
				// 注意i不能越界
				if i < len(p) && p[i] == searchWord[i] {
					tmp = append(tmp, p)
				}
			}
		} else {
			if len(ans) == 0 {
				continue
			}
			for _, w := range ans[len(ans)-1] {
				if i < len(w) && w[i] == searchWord[i] {
					tmp = append(tmp, w)
				}
			}
		}
		ans = append(ans, tmp)
	}
	// 最后截取每个[]string，限制最多为3个
	// 不可以在上面生成ans的过程中就截取了
	// 这样会截掉很多有用的字符串，这些字符串在下一轮生成的过程中可能就是满足前缀要求的
	for i, _ := range ans {
		ans[i] = ans[i][:min(3, len(ans[i]))]
	}
	return ans
}

type Trie struct {
	Node
}

type Node struct {
	Next  [26]*Node
	V     string
	IsEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) search(word string, add bool) *Node {
	cur := &t.Node
	for i := range word {
		b := word[i]
		next := cur.Next[b-'a']
		if next == nil {
			if add {
				next = &Node{}
				cur.Next[b-'a'] = next
			} else {
				return nil
			}
		}
		cur = next
	}
	if add {
		cur.V = word
	}
	return cur
}

func (t *Trie) Insert(word string) {
	t.search(word, true).IsEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.search(word, false)
	return node != nil && node.IsEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.search(prefix, false) != nil
}

func suggestedProductsTRIE(products []string, searchWord string) [][]string {
	// trie ?
	var root Trie
	for _, product := range products {
		root.Insert(product)
	}

	var dfs func(node *Node, buf []string) []string
	dfs = func(node *Node, buf []string) []string {
		if node == nil || len(buf) >= 3 {
			return buf
		}
		if node.IsEnd {
			buf = append(buf, node.V)
		}
		if len(buf) >= 3 {
			return buf
		}
		for _, nxt := range node.Next {
			buf = dfs(nxt, buf)
			if len(buf) >= 3 {
				break
			}
		}
		return buf
	}

	var ret = make([][]string, 0, len(searchWord))
	for i := range searchWord {
		ret = append(ret, dfs(root.search(searchWord[:i+1], false), nil))
	}
	return ret
}
